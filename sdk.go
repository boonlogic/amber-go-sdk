package amber_client

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	amberClient "github.com/boonlogic/amber-go-sdk/client"
	amberOps "github.com/boonlogic/amber-go-sdk/client/operations"
	amberModels "github.com/boonlogic/amber-go-sdk/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type LicenseProfile struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Server      string `json:"server"`
	OauthServer string `json:"oauth-server"`
}

type licenseProfiles map[string]LicenseProfile

type AmberClient struct {
	amberServer    *amberClient.AmberAPIServer
	oauthServer    *amberClient.AmberAPIServer
	oauthParams    *amberOps.PostOauth2Params
	reauthTime     time.Time
	verify         bool
	cert           string
	username       string
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

	if client.licenseProfile.Username == "" {
		return nil, errors.New("missing username in profile")
	}
	if client.licenseProfile.Password == "" {
		return nil, errors.New("missing password in profile")
	}
	if client.licenseProfile.Server == "" {
		return nil, errors.New("missing server in profile")
	}

	client.timeout = 360 * time.Second

	// set reauthTime to the past
	client.reauthTime = time.Now().Add(time.Second * -1)

	oauth2Request := amberModels.PostAuth2Request{
		Username: &client.licenseProfile.Username,
		Password: &client.licenseProfile.Password,
	}

	client.oauthParams = amberOps.NewPostOauth2Params()
	client.oauthParams.SetPostAuth2Request(&oauth2Request)
	client.oauthParams.WithTimeout(client.timeout)

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

func (a *AmberClient) ListSensors() (*amberModels.GetSensorsResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.GetSensorsParams{}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetSensors(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetSensorsUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetSensorsNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetSensorsBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetSensorsInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return &aok.Payload, nil
}

func (a *AmberClient) GetSensor(sensorId string) (*amberModels.GetSensorResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.GetSensorParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetSensor(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetSensorUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetSensorNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetSensorBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetSensorInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) CreateSensor(label string) (*amberModels.PostSensorResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PostSensorParams{
		PostSensorRequest: &amberModels.PostSensorRequest{
			Label: label,
		},
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostSensor(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostSensorUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostSensorNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostSensorBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostSensorInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) UpdateLabel(sensorId string, label string) (*amberModels.PutSensorResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PutSensorParams{
		PutSensorRequest: &amberModels.PutSensorRequest{
			Label: &label,
		},
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PutSensor(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PutSensorUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PutSensorNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PutSensorBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PutSensorInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) ConfigureSensor(sensorId string, payload amberModels.PostConfigRequest) (*amberModels.PostConfigResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PostConfigParams{
		PostConfigRequest: &payload,
		SensorID:          sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostConfig(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostConfigUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostConfigNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostConfigBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostConfigInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PostOutage(sensorId string) (*amberModels.PostOutageResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PostOutageParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostOutage(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostOutageUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostOutageNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostOutageBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostOutageInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) ConfigureFusion(sensorId string, payload amberModels.PutConfigRequest) (*amberModels.PutConfigResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PutConfigParams{
		PutConfigRequest: &payload,
		SensorID:         sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PutConfig(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PutConfigUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PutConfigNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PutConfigBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PutConfigInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) EnableLearning(sensorId string, payload amberModels.PutConfigRequest) (*amberModels.PutConfigResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PutConfigParams{
		PutConfigRequest: &payload,
		SensorID:         sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PutConfig(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PutConfigUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PutConfigNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PutConfigBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PutConfigInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) GetConfig(sensorId string) (*amberModels.GetConfigResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.GetConfigParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetConfig(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetConfigUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetConfigNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetConfigBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetConfigInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) DeleteSensor(sensorId string) *amberModels.Error {
	if result, aErr := a.authenticate(); result != true {
		return aErr
	}
	params := &amberOps.DeleteSensorParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	_, err := a.amberServer.Operations.DeleteSensor(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return err.(*amberOps.DeleteSensorUnauthorized).Payload
		case notFound:
			return err.(*amberOps.DeleteSensorNotFound).Payload
		case badRequest:
			return err.(*amberOps.DeleteSensorBadRequest).Payload
		case internalServerError:
			return err.(*amberOps.DeleteSensorInternalServerError).Payload
		default:
			return &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return nil
}

func (a *AmberClient) StreamSensor(sensorId string, payload amberModels.PostStreamRequest) (*amberModels.PostStreamResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PostStreamParams{
		PostStreamRequest: &payload,
		SensorID:          sensorId,
	}
	if params.PostStreamRequest.SaveImage == nil {
		saveImage := true
		params.PostStreamRequest.SaveImage = &saveImage
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostStream(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostStreamUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostStreamNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostStreamBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostStreamInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) StreamFusion(sensorId string, payload amberModels.PutStreamRequest) (*amberModels.PutStreamResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PutStreamParams{
		PutStreamRequest: &payload,
		SensorID:         sensorId,
	}
	params.WithTimeout(a.timeout)
	aok200, aok202, err := a.amberServer.Operations.PutStream(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PutStreamUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PutStreamNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PutStreamBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PutStreamInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}

	var putStreamResponse *amberModels.PutStreamResponse
	if aok200 != nil {
		putStreamResponse = aok200.Payload
	} else {
		putStreamResponse = aok202.Payload
	}
	return putStreamResponse, nil
}

func (a *AmberClient) GetStatus(sensorId string) (*amberModels.GetStatusResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.GetStatusParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetStatus(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetStatusUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetStatusNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetStatusBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetStatusInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) PretrainSensor(sensorId string, payload amberModels.PostPretrainRequest) (*amberModels.PostPretrainResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.PostPretrainParams{
		PostPretrainRequest: &payload,
		SensorID:            sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, accepted, err := a.amberServer.Operations.PostPretrain(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.PostPretrainUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.PostPretrainNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.PostPretrainBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.PostPretrainInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	if accepted != nil {
		acceptedResponse := amberModels.PostPretrainResponse{
			State:   accepted.Payload.State,
			Message: accepted.Payload.Message,
		}
		return &acceptedResponse, nil
	}
	return aok.Payload, nil
}

func (a *AmberClient) PretrainSensorXL(sensorId string, request amberModels.PostPretrainRequest) (*amberModels.PostPretrainResponse, *amberModels.Error) {

	var err error

	if result, aErr := a.authenticate(); result != true {
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

	params := &amberOps.PostPretrainParams{
		PostPretrainRequest: &request,
		SensorID:            sensorId,
	}

	var aok *amberOps.PostPretrainOK
	var accepted *amberOps.PostPretrainAccepted
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
		amberChunk := fmt.Sprintf("%v:%v", chunkIdx+1, chunkMax)
		params.AmberChunk = &amberChunk

		params.WithTimeout(a.timeout)
		aok, accepted, err = a.amberServer.Operations.PostPretrain(params, a.authWriter)
		if err != nil {
			switch errToken(err) {
			case unauthorized:
				return nil, err.(*amberOps.PostPretrainUnauthorized).Payload
			case notFound:
				return nil, err.(*amberOps.PostPretrainNotFound).Payload
			case badRequest:
				return nil, err.(*amberOps.PostPretrainBadRequest).Payload
			case internalServerError:
				return nil, err.(*amberOps.PostPretrainInternalServerError).Payload
			default:
				return nil, &amberModels.Error{Code: 500, Message: err.Error()}
			}
		}
		if accepted != nil {
			// accepted responses will have ambertransaction in the header
			params.AmberTransaction = &accepted.AmberTransaction
		}
	}

	if aok != nil {
		// pretraining complete
		aokResponse := amberModels.PostPretrainResponse{
			State:   aok.Payload.State,
			Message: aok.Payload.Message,
		}
		return &aokResponse, nil
	}

	if accepted != nil {
		// pretraining still running
		acceptedResponse := amberModels.PostPretrainResponse{
			State:   accepted.Payload.State,
			Message: accepted.Payload.Message,
		}
		return &acceptedResponse, nil
	}

	// if this code is reached, all of the chunks were sent but pretraining did not complete
	return nil, err.(*amberOps.PostPretrainInternalServerError).Payload
}

func (a *AmberClient) GetPretrainState(sensorId string) (*amberModels.GetPretrainResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
		return nil, aErr
	}
	params := &amberOps.GetPretrainParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, accepted, err := a.amberServer.Operations.GetPretrain(params, a.authWriter)
	if accepted != nil {
		acceptedResponse := amberModels.GetPretrainResponse{
			Message: accepted.Payload.Message,
			State:   accepted.Payload.State,
		}
		return &acceptedResponse, nil
	}
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetPretrainUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetPretrainNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetPretrainBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetPretrainInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return aok.Payload, nil
}

func (a *AmberClient) GetRootCause(sensorId string, clusterIds []int32, patterns [][]float32) (*amberModels.GetRootCauseResponse, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
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

	params := &amberOps.GetRootCauseParams{
		ClusterID: clusterp,
		Pattern:   patternp,
		SensorID:  sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetRootCause(params, a.authWriter)
	if err != nil {
		switch errToken(err) {
		case unauthorized:
			return nil, err.(*amberOps.GetRootCauseUnauthorized).Payload
		case notFound:
			return nil, err.(*amberOps.GetRootCauseNotFound).Payload
		case badRequest:
			return nil, err.(*amberOps.GetRootCauseBadRequest).Payload
		case internalServerError:
			return nil, err.(*amberOps.GetRootCauseInternalServerError).Payload
		default:
			return nil, &amberModels.Error{Code: 500, Message: err.Error()}
		}
	}
	return &aok.Payload, nil
}

func (a *AmberClient) GetVersion() (*amberModels.Version, *amberModels.Error) {
	if result, aErr := a.authenticate(); result != true {
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
		response, err := a.oauthServer.Operations.PostOauth2(a.oauthParams)
		if err != nil {
			switch errToken(err) {
			case unauthorized:
				return false, err.(*amberOps.PostOauth2Unauthorized).Payload
			case internalServerError:
				return false, err.(*amberOps.PostOauth2InternalServerError).Payload
			default:
				return false, &amberModels.Error{Code: 500, Message: err.Error()}
			}
		}

		// save the token as an authWriter
		a.authWriter = httptransport.BearerToken(*response.Payload.IDToken)

		// save the expiration time (-60 seconds)
		expiresIn, err := strconv.ParseUint(*response.Payload.ExpiresIn, 10, 64)
		if err != nil {
			return false, &amberModels.Error{Code: 500, Message: err.Error()}
		}
		a.reauthTime = tIn.Add(time.Second * time.Duration(expiresIn-60))
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
	if username := os.Getenv("AMBER_USERNAME"); username != "" {
		a.licenseProfile.Username = username
	}
	if password := os.Getenv("AMBER_PASSWORD"); password != "" {
		a.licenseProfile.Password = password
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
		for _, value := range line[:] {
			flattened = append(flattened, value)
		}
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
	case strings.Index(strType, unauthorized) != -1:
		return unauthorized
	case strings.Index(strType, internalServerError) != -1:
		return internalServerError
	case strings.Index(strType, badRequest) != -1:
		return badRequest
	case strings.Index(strType, notFound) != -1:
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
