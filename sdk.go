package amber_client

import (
	aclient "amber-go-sdk/ambergen/client"
	aops "amber-go-sdk/ambergen/client/operations"
	amodels "amber-go-sdk/ambergen/models"
	"encoding/json"
	"errors"
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
	amberServer    *aclient.AmberAPIServer
	oauthServer    *aclient.AmberAPIServer
	oauthParams    *aops.PostOauth2Params
	reauthTime     time.Time
	licenseFile    string
	licenseId      string
	verify         bool
	cert           string
	username       string
	licenseProfile LicenseProfile
	timeout        uint
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
		ac.timeout = 360
	} else {
		ac.timeout = *timeout
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

	ac.oauthParams = &aops.PostOauth2Params{
		PostAuth2Request: &amodels.PostAuth2Request{
			Username: &ac.licenseProfile.Username,
			Password: &ac.licenseProfile.Password,
		},
	}

	// set up default transport
	ac.amberServer = aclient.New(httptransport.New(ac.licenseProfile.Server, "", nil), strfmt.Default)

	// set up oauth2 transport
	ac.oauthServer = aclient.New(httptransport.New(ac.licenseProfile.OauthServer, "", nil), strfmt.Default)

	return &ac, nil
}

func (a AmberClient) authenticate() bool {
	tIn := time.Now()
	if a.reauthTime.Before(tIn) {
		response, err := a.oauthServer.Operations.PostOauth2(a.oauthParams, nil)
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

func (a AmberClient) ListSensors() (*amodels.GetSensorsResponse, error) {

	if a.authenticate() == false {
		return nil, errors.New("authentication failed")
	}

	params := &aops.GetSensorsParams{}
	aok, err := a.amberServer.Operations.GetSensors(params, a.authWriter, nil)
	if err != nil {
		return nil, err
	}

	return &aok.Payload, nil
}

func (a AmberClient) GetSensor(sensorId string) (*amodels.GetSensorResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.GetSensorResponse{}
	return response, nil
}

func (a AmberClient) CreateSensor(label string) (*amodels.PostSensorResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.PostSensorResponse{}
	return response, nil
}

func (a AmberClient) UpdateLabel(sensorId string, label string) (*amodels.PutSensorResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.PutSensorResponse{}
	return response, nil
}

func (a AmberClient) ConfigureSensor(sensorId string, featureCount int, streamingWindowSize int,
	samplesToBuffer int, learningRateNumerator int,
	learningRateDenominator int, learningMaxClusters int,
	learningMaxSamples int, anomalyHistoryWindow int,
	features amodels.PostFeatureConfig) (*amodels.PostConfigResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.PostConfigResponse{}
	return response, nil
}

func (a AmberClient) GetConfig(sensorId string) (*amodels.GetConfigResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.GetConfigResponse{}
	return response, nil
}

func (a AmberClient) DeleteSensor(sensorId string) error {

	if (a.authenticate()) == false {
		return errors.New("authentication failed")
	}
	return nil
}

func (a AmberClient) StreamSensor(sensorId string, csv string, saveImage bool) (*amodels.PostStreamResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.PostStreamResponse{}
	return response, nil
}

func (a AmberClient) GetStatus(sensorId string) (*amodels.GetStatusResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.GetStatusResponse{}
	return response, nil
}

func (a AmberClient) PretrainSensor(sensorId string, csv string, autotuneConfig bool) (*amodels.PostPretrainResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.PostPretrainResponse{}
	return response, nil
}

func (a AmberClient) GetPretrainState(sensorId string) (*amodels.GetPretrainResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.GetPretrainResponse{}
	return response, nil
}

func (a AmberClient) GetRootCause(sensorId string, clusterId string, pattern string) (*amodels.GetRootCauseResponse, error) {

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	response := &amodels.GetRootCauseResponse{}
	return response, nil
}

func (a AmberClient) GetVersion() (string, error) {

	if (a.authenticate()) == false {
		return "", errors.New("authentication failed")
	}

	return "", nil
}
