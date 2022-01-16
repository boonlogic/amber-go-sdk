package amber_client

import (
	"encoding/json"
	"errors"
	amberClient "github.com/boonlogic/amber-go-sdk/client"
	amberOps "github.com/boonlogic/amber-go-sdk/client/operations"
	amberModels "github.com/boonlogic/amber-go-sdk/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

type LicenseProfiles map[string]LicenseProfile

type AmberClient struct {
	amberServer    *amberClient.AmberAPIServer
	oauthServer    *amberClient.AmberAPIServer
	oauthParams    *amberOps.PostOauth2Params
	reauthTime     time.Time
	licenseFile    string
	licenseId      string
	verify         bool
	cert           string
	username       string
	licenseProfile LicenseProfile
	timeout        time.Duration
	authWriter     runtime.ClientAuthInfoWriter
	proxy          string
}

func envFallback(key string, defVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defVal
	}
	return val
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

func NewAmberClient(licenseId *string, licenseFile *string, verify *bool, cert *string, timeout *uint) (*AmberClient, error) {
	var ac AmberClient
	var err error

	// set default values
	if licenseId == nil {
		ac.licenseId = "default"
	} else {
		ac.licenseId = *licenseId
	}
	if licenseFile == nil {
		ac.licenseFile = "~/.Amber.license"
	} else {
		ac.licenseFile = *licenseFile
	}
	if verify == nil {
		ac.verify = true
	} else {
		ac.verify = *verify
	}
	if cert == nil {
		ac.cert = ""
	} else {
		ac.cert = *cert
	}
	if timeout == nil {
		ac.timeout = 360 * time.Second
	} else {
		ac.timeout = time.Duration(*timeout) * time.Second
	}

	// initialize reauth timer to expired time
	ac.reauthTime = time.Now().Add(time.Second * -1)

	// set default license file
	licenseEnv := os.Getenv("AMBER_LICENSE_FILE")
	if licenseEnv != "" {
		ac.licenseFile = licenseEnv
	}

	var lp LicenseProfiles
	if ac.licenseFile != "" {
		// expand home directory if necessary
		if strings.HasPrefix(ac.licenseFile, "~/") {
			dirname, _ := os.UserHomeDir()
			ac.licenseFile = filepath.Join(dirname, ac.licenseFile[2:])
		}
		if _, err = os.Stat(ac.licenseFile); err == nil {
			//
			blob, err := ioutil.ReadFile(ac.licenseFile)
			if err != nil {
				return nil, err
			}

			if err := json.Unmarshal(blob, &lp); err != nil {
				return nil, err
			}
		}
		ac.licenseProfile = lp[ac.licenseId]
	}

	// override from environment
	ac.licenseProfile.Username = envFallback("AMBER_USERNAME", ac.licenseProfile.Username)
	ac.licenseProfile.Password = envFallback("AMBER_PASSWORD", ac.licenseProfile.Password)
	ac.licenseProfile.Server = envFallback("AMBER_SERVER", ac.licenseProfile.Server)
	ac.licenseProfile.OauthServer = envFallback("AMBER_OAUTH_SERVER", ac.licenseProfile.OauthServer)
	if ac.licenseProfile.OauthServer == "" {
		ac.licenseProfile.OauthServer = ac.licenseProfile.Server
	}
	ac.proxy = envFallback("AMBER_PROXY", "")
	ac.cert = envFallback("AMBER_SSL_CERT", ac.cert)
	ac.verify = strings.ToLower(envFallback("AMBER_SSL_VERIFY", "false")) == "true"

	// verify required profile elements have been created
	if ac.licenseProfile.Username == "" {
		return nil, errors.New("missing username in profile")
	}
	if ac.licenseProfile.Password == "" {
		return nil, errors.New("missing username in profile")
	}
	if ac.licenseProfile.Server == "" {
		return nil, errors.New("missing username in profile")
	}

	oauth2Request := amberModels.PostAuth2Request{
		Username: &ac.licenseProfile.Username,
		Password: &ac.licenseProfile.Password,
	}

	ac.oauthParams = amberOps.NewPostOauth2Params()
	ac.oauthParams.SetPostAuth2Request(&oauth2Request)
	ac.oauthParams.WithTimeout(ac.timeout)

	// set up default server
	_, host, basePath, _ := parseServer(ac.licenseProfile.Server)
	ac.amberServer = amberClient.New(httptransport.New(host, basePath, nil), strfmt.Default)

	// set up oauth server
	_, host, basePath, _ = parseServer(ac.licenseProfile.OauthServer)
	ac.oauthServer = amberClient.New(httptransport.New(host, basePath, nil), strfmt.Default)

	return &ac, nil
}

func (a *AmberClient) authenticate() bool {
	tIn := time.Now()
	if a.reauthTime.Before(tIn) {
		response, err := a.oauthServer.Operations.PostOauth2(a.oauthParams)
		if err != nil {
			return false
		}

		// save the token as an authWriter
		a.authWriter = httptransport.BearerToken(*response.Payload.IDToken)

		// save the expiration time (-60 seconds)
		expiresIn, err := strconv.ParseUint(*response.Payload.ExpiresIn, 10, 64)
		if err != nil {
			return false
		}
		a.reauthTime = tIn.Add(time.Second * time.Duration(expiresIn-60))
	}
	return true
}

func (a *AmberClient) ListSensors() (*amberModels.GetSensorsResponse, error) {
	if a.authenticate() == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.GetSensorsParams{}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetSensors(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return &aok.Payload, nil
}

func (a AmberClient) GetSensor(sensorId string) (*amberModels.GetSensorResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.GetSensorParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetSensor(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) CreateSensor(label string) (*amberModels.PostSensorResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.PostSensorParams{
		PostSensorRequest: &amberModels.PostSensorRequest{
			Label: label,
		},
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostSensor(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) UpdateLabel(sensorId string, label string) (*amberModels.PutSensorResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
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
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) ConfigureSensor(sensorId string, payload amberModels.PostConfigRequest) (*amberModels.PostConfigResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.PostConfigParams{
		PostConfigRequest: &payload,
		SensorID:          sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostConfig(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) GetConfig(sensorId string) (*amberModels.GetConfigResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.GetConfigParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetConfig(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) DeleteSensor(sensorId string) error {
	if (a.authenticate()) == false {
		return errors.New("authentication failed")
	}
	params := &amberOps.DeleteSensorParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	_, err := a.amberServer.Operations.DeleteSensor(params, a.authWriter)
	if err != nil {
		return err
	}
	return nil
}

func (a AmberClient) StreamSensor(sensorId string, payload amberModels.PostStreamRequest) (*amberModels.PostStreamResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.PostStreamParams{
		PostStreamRequest: &payload,
		SensorID:          sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.PostStream(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) GetStatus(sensorId string) (*amberModels.GetStatusResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.GetStatusParams{
		SensorID: sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetStatus(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) PretrainSensor(sensorId string, payload amberModels.PostPretrainRequest) (*amberModels.PostPretrainResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.PostPretrainParams{
		PostPretrainRequest: &payload,
		SensorID:            sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, accepted, err := a.amberServer.Operations.PostPretrain(params, a.authWriter)
	if err != nil {
		return nil, err
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

func (a AmberClient) GetPretrainState(sensorId string) (*amberModels.GetPretrainResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
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
		return nil, err
	}
	return aok.Payload, nil
}

func (a AmberClient) GetRootCause(sensorId string, clusterId *string, pattern *string) (*amberModels.GetRootCauseResponse, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.GetRootCauseParams{
		ClusterID: clusterId,
		Pattern:   pattern,
		SensorID:  sensorId,
	}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetRootCause(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return &aok.Payload, nil
}

func (a AmberClient) GetVersion() (*amberModels.Version, error) {
	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}
	params := &amberOps.GetVersionParams{}
	params.WithTimeout(a.timeout)
	aok, err := a.amberServer.Operations.GetVersion(params, a.authWriter)
	if err != nil {
		return nil, err
	}
	return aok.Payload, nil
}
