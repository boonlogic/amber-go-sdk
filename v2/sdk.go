package amber_client_v2

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	amberClient "github.com/boonlogic/amber-go-sdk/v2/client"
	amberOps "github.com/boonlogic/amber-go-sdk/v2/client/operations"
	amberModels "github.com/boonlogic/amber-go-sdk/v2/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type LicenseProfile struct {
	LicenseKey  string `json:"license-key"`
	SecretKey   string `json:"secret-key"`
	Server      string `json:"server"`
	OauthServer string `json:"oauth-server"`
}

type licenseProfiles map[string]LicenseProfile

type AmberClient struct {
	amberServer    *amberClient.AmberAPIServer
	oauthServer    *amberClient.AmberAPIServer
	accessParams   *amberOps.PostOauth2AccessParams
	refreshParams  *amberOps.PostOauth2RefreshParams
	reauthTime     time.Time
	verify         bool
	cert           string
	licenseProfile LicenseProfile
	timeout        time.Duration
	authWriter     runtime.ClientAuthInfoWriter
	proxy          string
	tlsOptions     httptransport.TLSClientOptions
}

// Create new AmberClient given LicenseProfile structure
func NewAmberClientFromProfile(profile LicenseProfile) (*AmberClient, error) {

	// create client when given LicenseProfile
	var client AmberClient
	client.verify = true
	client.licenseProfile = profile

	// override from environment
	client.loadFromEnv()

	if client.licenseProfile.OauthServer == "" {
		client.licenseProfile.OauthServer = client.licenseProfile.Server
	}

	if client.licenseProfile.LicenseKey == "" {
		return nil, errors.New("missing license-key in profile")
	}
	if client.licenseProfile.SecretKey == "" {
		return nil, errors.New("missing secret-key in profile")
	}
	if client.licenseProfile.Server == "" {
		return nil, errors.New("missing server in profile")
	}

	client.timeout = 360 * time.Second

	// set reauthTime to the past
	client.reauthTime = time.Now().Add(time.Second * -1)

	oauth2Request := amberModels.PostOauth2AccessRequest{
		LicenseKey: &client.licenseProfile.LicenseKey,
		SecretKey:  &client.licenseProfile.SecretKey,
	}

	client.accessParams = amberOps.NewPostOauth2AccessParams()
	client.accessParams.SetPostOauth2AccessRequest(&oauth2Request)
	client.accessParams.WithTimeout(client.timeout)

	client.refreshParams = nil //This will be filled in once we authenticate

	client.updateHttpClients()

	return &client, nil
}

// Create new AmberClient using Amber license file
func NewAmberClientFromFile(licenseId *string, licenseFile *string) (*AmberClient, error) {

	var file = "~/.Amber.license"
	var id = "default"
	var profile LicenseProfile

	// construct a license profile based on configuration file
	if licenseFile != nil {
		// use default
		file = *licenseFile
	}
	if licenseId != nil {
		// use default
		id = *licenseId
	}

	// override licenseFile and licenseId from environment
	var envValue string
	envValue = os.Getenv("AMBER_LICENSE_FILE")
	if envValue != "" {
		file = envValue
	}
	envValue = os.Getenv("AMBER_LICENSE_ID")
	if envValue != "" {
		id = envValue
	}

	// expand home directory if necessary
	if strings.HasPrefix(file, "~/") {
		dirname, _ := os.UserHomeDir()
		file = filepath.Join(dirname, file[2:])
	}

	var lp licenseProfiles
	if _, err := os.Stat(file); err == nil {
		blob, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(blob, &lp); err != nil {
			return nil, err
		}
		if _, ok := lp[id]; !ok {
			return nil, fmt.Errorf("license id '%v' not found", id)
		}
		profile = lp[id]
	} else {
		return nil, fmt.Errorf("license file '%v' not found", file)
	}
	return NewAmberClientFromProfile(profile)
}

func (a *AmberClient) SetVerify(value bool) error {
	a.verify = value
	a.updateHttpClients()
	return nil
}

func (a *AmberClient) SetCert(cert string) error {
	a.cert = cert
	a.updateHttpClients()
	return nil
}

func (a *AmberClient) SetTimeout(timeout int) error {
	a.timeout = time.Duration(timeout) * time.Second
	return nil
}

func (a *AmberClient) SetProxy(proxy string) error {
	a.proxy = proxy
	_ = os.Setenv("HTTP_PROXY", proxy)
	return nil
}

func (a *AmberClient) GetModels() (*amberModels.GetModelsResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.GetModelsParams{}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetModels(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetModelsUnauthorized).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetModelsInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) GetModel(modelId string) (*amberModels.PostModelResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.GetModelParams{
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetModel(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetModelUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetModelNotFound).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetModelInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PostModel(label string) (*amberModels.PostModelResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.PostModelParams{
		PostModelRequest: &amberModels.PostModelRequest{
			Label: label,
		},
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostModel(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostModelUnauthorized).Payload
		case badRequest:
			return nil, err.(*amberOps.PostModelBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostModelInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PutModel(modelId string, label string) (*amberModels.PostModelResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.PutModelParams{
		PutModelRequest: &amberModels.PutModelRequest{
			Label: &label,
		},
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PutModel(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PutModelUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PutModelNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PutModelBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PutModelInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PostModelConfig(modelId string, payload amberModels.PostConfigRequest) (*amberModels.PostConfigResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.PostModelConfigParams{
		PostConfigRequest: &payload,
		ModelID:           modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostModelConfig(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostModelConfigUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostModelConfigNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostModelConfigBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostModelConfigInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PostModelOutage(modelId string) *amberModels.Error {
	if result, aErr := a.authenticate(); !result {
		return aErr
	}
	params := &amberOps.PostModelOutageParams{
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	_, err := a.amberServer.Operations.PostModelOutage(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return err.(*amberOps.PostModelOutageUnauthorized).Payload
		case notFound:
			return err.(*amberOps.PostModelOutageNotFound).Payload
		case internalServerError:
			return err.(*amberOps.PostModelOutageInternalServerError).Payload
		default:
			return &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return nil
}

func (a *AmberClient) PostModelLearning(modelId string, payload amberModels.PostLearningRequest) (*amberModels.PostLearningResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.PostModelLearningParams{
		PostLearningRequest: &payload,
		ModelID:             modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostModelLearning(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostModelLearningUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostModelLearningNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostModelLearningBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostModelLearningInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) GetModelConfig(modelId string) (*amberModels.PostConfigResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.GetModelConfigParams{
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetModelConfig(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetModelConfigUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetModelConfigNotFound).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetModelConfigInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) DeleteModel(modelId string) *amberModels.Error {
	if result, aErr := a.authenticate(); !result {
		return aErr
	}
	params := &amberOps.DeleteModelParams{
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	_, err := a.amberServer.Operations.DeleteModel(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return err.(*amberOps.DeleteModelUnauthorized).Payload
		case notFound:
			return err.(*amberOps.DeleteModelNotFound).Payload
		case internalServerError:
			return err.(*amberOps.DeleteModelInternalServerError).Payload
		default:
			return &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return nil
}

func (a *AmberClient) PostModelData(modelId string, payload amberModels.PostDataRequest) (*amberModels.PostDataResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.PostModelDataParams{
		PostDataRequest: &payload,
		ModelID:         modelId,
	}
	if params.PostDataRequest.SaveImage == nil {
		saveImage := true
		params.PostDataRequest.SaveImage = &saveImage
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostModelData(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostModelDataUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostModelDataNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostModelDataBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostModelDataInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PutModelData(modelId string, payload amberModels.PutDataRequest) (*amberModels.PutDataResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.PutModelDataParams{
		PutDataRequest: &payload,
		ModelID:        modelId,
	}
	params.WithTimeout(a.timeout)
	aok200, aok202, err := a.amberServer.Operations.PutModelData(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PutModelDataUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PutModelDataNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PutModelDataBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PutModelDataInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}

	var putStreamResponse *amberModels.PutDataResponse
	if aok200 != nil {
		putStreamResponse = aok200.Payload
	} else {
		putStreamResponse = aok202.Payload
	}
	return putStreamResponse, nil
}

func (a *AmberClient) GetModelStatus(modelId string) (*amberModels.GetStatusResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.GetModelStatusParams{
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetModelStatus(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetModelStatusUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetModelStatusNotFound).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetModelStatusInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PretrainModel(modelId string, request amberModels.PostPretrainRequest) (*amberModels.PostPretrainResponse, *amberModels.Error) {

	var err error

	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}

	// send payload in chunks by default
	floatFormat := "packed-float"
	request.Format = &floatFormat

	// convert to packed floats
	packedFloats, err := PackCsvAsByteSlice(*request.Data)
	if err != nil {
		return nil, &amberModels.Error{Code: 400, Message: err.Error()}
	}

	// post chunks of 4MB
	chunkSize := 4000000
	chunkMax := len(packedFloats) / chunkSize
	if len(packedFloats)%chunkSize != 0 {
		chunkMax += 1
	}

	params := &amberOps.PostModelPretrainParams{
		PostPretrainRequest: &request,
		ModelID:             modelId,
	}

	var aok *amberOps.PostModelPretrainOK
	var accepted *amberOps.PostModelPretrainAccepted
	for chunkIdx := 0; chunkIdx < chunkMax; chunkIdx++ {

		start := chunkIdx * chunkSize
		end := (chunkIdx + 1) * chunkSize
		if end > len(packedFloats) {
			end = len(packedFloats)
		}

		// prepare next chunk
		next := packedFloats[start:end]
		chunk := base64.StdEncoding.EncodeToString(next)
		params.PostPretrainRequest.Data = &chunk

		// apply amberchunk http header
		chunkSpec := fmt.Sprintf("%v:%v", chunkIdx+1, chunkMax)
		params.Chunkspec = &chunkSpec

		params.WithTimeout(a.timeout)
		aok, accepted, err = a.amberServer.Operations.PostModelPretrain(params, a.authWriter)
		if err != nil {
			switch errToken(err) {
			case unauthorized:
				return nil, err.(*amberOps.PostModelPretrainUnauthorized).Payload
			case notFound:
				return nil, err.(*amberOps.PostModelPretrainNotFound).Payload
			case badRequest:
				return nil, err.(*amberOps.PostModelPretrainBadRequest).Payload
			case internalServerError:
				return nil, err.(*amberOps.PostModelPretrainInternalServerError).Payload
			default:
				return nil, &amberModels.Error{Code: 500, Message: err.Error()}
			}
		}
		if accepted != nil {
			// accepted responses will have ambertransaction in the header
			params.TxnID = &accepted.TxnID
		}
	}

	if aok != nil {
		// pretraining complete
		aokResponse := amberModels.PostPretrainResponse{
			PretrainStatus: amberModels.PretrainStatus{
				Status:  aok.Payload.Status,
				Message: aok.Payload.Message,
			},
		}
		return &aokResponse, nil
	}

	if accepted != nil {
		// pretraining still running
		acceptedResponse := amberModels.PostPretrainResponse{
			PretrainStatus: amberModels.PretrainStatus{
				Status:  accepted.Payload.Status,
				Message: accepted.Payload.Message,
			},
		}
		return &acceptedResponse, nil
	}

	// if this code is reached, all of the chunks were sent but pretraining did not complete
	return nil, err.(*amberOps.PostModelPretrainInternalServerError).Payload
}

func (a *AmberClient) GetPretrainState(modelId string) (*amberModels.GetPretrainResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.GetModelPretrainParams{
		ModelID: modelId,
	}
	params.WithTimeout(a.timeout)
	aok, accepted, err := a.amberServer.Operations.GetModelPretrain(params, a.authWriter)
	if accepted != nil {
		acceptedResponse := amberModels.GetPretrainResponse{
			PretrainStatus: amberModels.PretrainStatus{
				Status:  accepted.Payload.Status,
				Message: accepted.Payload.Message,
			},
		}
		return &acceptedResponse, nil
	}
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetModelPretrainUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetModelPretrainNotFound).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetModelPretrainInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) GetRootCause(modelId string, clusterIds []int32, patterns [][]float32) (*amberModels.GetRootCauseResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}

	lenc := len(clusterIds)
	lenp := len(patterns)

	if lenc == 0 && lenp == 0 { // neither specified
		err := &amberModels.Error{Code: 400, Message: "Must specify either patterns or cluster IDs for analysis"}
		return nil, err
	}
	if lenc > 0 && lenp > 0 { // both specified
		err := &amberModels.Error{Code: 400, Message: "Cannot specify both patterns and cluster IDs for analysis"}
		return nil, err
	}

	var patternp *string
	var clusterp *string

	if lenc > 0 {
		clusterb, e := json.Marshal(clusterIds)
		if e == nil {
			c := string(clusterb)
			clusterp = &c
		}
	} else {
		patternb, e := json.Marshal(patterns)
		if e == nil {
			p := string(patternb)
			patternp = &p
		}
	}

	params := &amberOps.GetModelRootCauseParams{
		Clusters: clusterp,
		Vectors:  patternp,
		ModelID:  modelId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetModelRootCause(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetModelRootCauseUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetModelRootCauseNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetModelRootCauseBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetModelRootCauseInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) GetVersion() (*amberModels.GetVersionResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); !result {
		return nil, aErr
	}
	params := &amberOps.GetVersionParams{}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetVersion(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case internalServerError:
			return nil, err.(*amberOps.GetVersionInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func parseServer(server string) (string, string, string, error) {

	var scheme, host, basepath string

	// parse the scheme
	if strings.HasPrefix(server, "https://") {
		scheme = "https"
		server = strings.TrimPrefix(server, "https://")
	} else if strings.HasPrefix(server, "http://") {
		scheme = "http"
		server = strings.TrimPrefix(server, "http://")
	} else {
		scheme = ""
	}

	indexOf := strings.Index(server, "/")
	if indexOf == -1 {
		host = server
		basepath = ""
	} else {
		host = server[0:indexOf]
		basepath = server[indexOf:]
	}

	return scheme, host, basepath, nil
}

func (a *AmberClient) authenticate() (bool, *amberModels.Error) {
	tIn := time.Now()
	if a.reauthTime.Before(tIn) {
		if a.refreshParams != nil {
			//If we have a refresh token, lets use that
			response, err := a.oauthServer.Operations.PostOauth2Refresh(a.refreshParams)
			if err != nil {
				//we must do an access again
				a.refreshParams = nil
				return a.authenticate()
			}

			// save the token as an authWriter
			a.authWriter = httptransport.BearerToken(response.Payload.IDToken)

			//save refresh token
			oauth2Refresh := amberModels.PostOauth2RefreshRequest{
				RefreshToken: &response.Payload.RefreshToken,
			}
			a.refreshParams.SetPostOauth2RefreshRequest(&oauth2Refresh)

			// save the expiration time (-60 seconds)
			expiresIn, err := strconv.ParseUint(response.Payload.ExpiresIn, 10, 64)
			if err != nil {
				return false, &amberModels.Error{Code: 500, Message: err.Error()}
			}
			a.reauthTime = tIn.Add(time.Second * time.Duration(expiresIn-60))
		} else {
			//If we dont have a refresh token, post an access request
			response, err := a.oauthServer.Operations.PostOauth2Access(a.accessParams)
			if err != nil {
				switch errToken(err) {
				case unauthorized:
					return false, err.(*amberOps.PostOauth2AccessUnauthorized).Payload
				case internalServerError:
					return false, err.(*amberOps.PostOauth2AccessInternalServerError).Payload
				default:
					return false, &amberModels.Error{Code: 500, Message: err.Error()}
				}
			}

			// save the token as an authWriter
			a.authWriter = httptransport.BearerToken(response.Payload.IDToken)

			//save refresh token
			oauth2Refresh := amberModels.PostOauth2RefreshRequest{
				RefreshToken: &response.Payload.RefreshToken,
			}
			a.refreshParams = amberOps.NewPostOauth2RefreshParams()
			a.refreshParams.SetPostOauth2RefreshRequest(&oauth2Refresh)
			a.refreshParams.WithTimeout(a.timeout)

			// save the expiration time (-60 seconds)
			expiresIn, err := strconv.ParseUint(response.Payload.ExpiresIn, 10, 64)
			if err != nil {
				return false, &amberModels.Error{Code: 500, Message: err.Error()}
			}
			a.reauthTime = tIn.Add(time.Second * time.Duration(expiresIn-60))
		}
	}
	return true, nil
}

type CustomRoundTripper struct {
	Proxied http.RoundTripper
}

func (crt CustomRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {

	// compress the payload if content-length is getting big
	if req.ContentLength > 10000 {
		var err error

		// squeeze out all white space
		buf := new(bytes.Buffer)
		if _, err = buf.ReadFrom(req.Body); err != nil {
			return nil, err
		}
		newStr := buf.String()
		newStr = strings.ReplaceAll(newStr, " ", "")

		// gzip the buffer
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		if _, err = gz.Write([]byte(newStr)); err != nil {
			return nil, err
		}
		if err = gz.Close(); err != nil {
			return nil, err
		}
		reader := bytes.NewReader(b.Bytes())
		req.Body = io.NopCloser(reader)

		// set the content-encoding
		req.Header.Set("content-encoding", "gzip")
		req.Header.Set("content-type", "application/json")

		// Set ContentLength to reflect the new size
		req.ContentLength = int64(b.Len())
	}

	req.Header.Set("User-Agent", "Boon Logic / amber-go-sdk / http")

	// Send the request, get the response
	res, e = crt.Proxied.RoundTrip(req)
	return res, e
}

func (a *AmberClient) updateHttpClients() {

	// set default verify and cert
	a.tlsOptions.InsecureSkipVerify = !a.verify
	a.tlsOptions.Certificate = a.cert

	// set server http client
	scheme, host, basePath, _ := parseServer(a.licenseProfile.Server)
	if scheme == "https" {
		httpClient, _ := httptransport.TLSClient(a.tlsOptions)
		httpClient.Transport = CustomRoundTripper{httpClient.Transport}
		clientRuntime := httptransport.NewWithClient(host, basePath, []string{scheme}, httpClient)
		a.amberServer = amberClient.New(clientRuntime, strfmt.Default)
	} else {
		transport := httptransport.New(host, basePath, []string{scheme})
		transport.Transport = CustomRoundTripper{transport.Transport}
		a.amberServer = amberClient.New(transport, strfmt.Default)
	}

	// set up oauth http client
	scheme, host, basePath, _ = parseServer(a.licenseProfile.OauthServer)
	if scheme == "https" {
		httpClient, _ := httptransport.TLSClient(a.tlsOptions)
		httpClient.Transport = CustomRoundTripper{httpClient.Transport}
		clientRuntime := httptransport.NewWithClient(host, basePath, []string{scheme}, httpClient)
		a.oauthServer = amberClient.New(clientRuntime, strfmt.Default)
	} else {
		a.oauthServer = amberClient.New(httptransport.New(host, basePath, []string{scheme}), strfmt.Default)
	}
}

func (a *AmberClient) loadFromEnv() {

	// construct a license profile based on environment
	if license_key := os.Getenv("AMBER_LICENSE_KEY"); license_key != "" {
		a.licenseProfile.LicenseKey = license_key
	}
	if secret_key := os.Getenv("AMBER_SECRET_KEY"); secret_key != "" {
		a.licenseProfile.SecretKey = secret_key
	}
	if server := os.Getenv("AMBER_SERVER"); server != "" {
		a.licenseProfile.Server = server
	}
	if oauthServer := os.Getenv("AMBER_OAUTH_SERVER"); oauthServer != "" {
		a.licenseProfile.OauthServer = oauthServer
	}
	verifyEnv := os.Getenv("AMBER_SSL_VERIFY")
	a.verify = true
	if strings.ToLower(verifyEnv) == "false" {
		a.verify = false
	}
	if proxy := os.Getenv("AMBER_PROXY"); proxy != "" {
		a.proxy = strings.ToLower(proxy)
	}
	if cert := os.Getenv("AMBER_SSL_CERT"); cert != "" {
		a.cert = strings.ToLower(cert)
	}
}

// utility function to create csv payloads given a csv data file.
// output will be one continuous string of comma separated values
func LoadCsvRecords(csvFile string) ([]string, error) {

	// open file
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	// flatten
	flattened := []string{}
	for _, line := range data[:] {
		flattened = append(flattened, line[:]...)
	}
	return flattened, nil
}

const unauthorized = "Unauthorized"
const notFound = "NotFound"
const internalServerError = "InternalServerError"
const badRequest = "BadRequest"
const unknown = "Unknown"

func errToken(err error) string {
	strType := reflect.TypeOf(err).String()
	switch {
	case strings.Contains(strType, unauthorized):
		return unauthorized
	case strings.Contains(strType, internalServerError):
		return internalServerError
	case strings.Contains(strType, badRequest):
		return badRequest
	case strings.Contains(strType, notFound):
		return notFound
	}

	return unknown
}

func PackCsvAsByteSlice(csv string) ([]byte, error) {
	csv = strings.TrimSpace(csv)
	csvSlice := strings.Split(csv, ",")
	sampleCnt := len(csvSlice)
	floatBuf := make([]byte, sampleCnt*4)
	for i, arg := range csvSlice {
		if n, err := strconv.ParseFloat(arg, 32); err == nil {
			binary.LittleEndian.PutUint32(floatBuf[i*4:], math.Float32bits(float32(n)))
		} else {
			// data not properly formatted
			return nil, fmt.Errorf("csv string not formatted properly")
		}
	}
	return floatBuf, nil
}
