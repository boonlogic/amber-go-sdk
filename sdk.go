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

	// "github.com/go-openapi/runtime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type LicenseProfile struct {
	username    string `json:"username"`
	password    string `json:"password"`
	server      string `json:"server"`
	oauthServer string `json:"oauth-server"`
}

type LicenseProfiles map[string]LicenseProfile

type AmberClient struct {
	amberServer    *aclient.AmberAPIServer
	oauthServer    *aclient.AmberAPIServer
	oauthParams    *aops.PostOauth2Params
	reauthTime     time.Time
	licenseFile    string
	licenseId      string
	username       string
	licenseProfile LicenseProfile
	timeout        uint
	authWriter     runtime.ClientAuthInfoWriter
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

	// initialize reauth timer to expired time
	ac.reauthTime = time.Now().Add(time.Second * -1)

	// load from license file if necessary, override from environment if specified
	var exists bool
	ac.licenseFile, exists = os.LookupEnv("AMBER_LICENSE_FILE")
	if exists {
		// expand home directory if necessary
		if strings.HasPrefix(ac.licenseFile, "~/") {
			dirname, _ := os.UserHomeDir()
			ac.licenseFile = filepath.Join(dirname, ac.licenseFile[2:])
		}
		if _, err = os.Stat(ac.licenseFile); err == nil {
			//
			blob, err := os.ReadFile(ac.licenseFile)
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(blob, &ac.licenseProfile); err != nil {
				return nil, err
			}
		} else {
			if licenseFile != nil {
				// if license file is something other than default, throw exception
				ac.licenseFile = *licenseFile
			}
		}
	}

	// override from environment
	ac.licenseProfile.username = envFallback("AMBER_USERNAME", ac.licenseProfile.username)
	ac.licenseProfile.password = envFallback("AMBER_PASSWORD", ac.licenseProfile.password)
	ac.licenseProfile.server = envFallback("AMBER_SERVER", ac.licenseProfile.server)
	ac.licenseProfile.oauthServer = envFallback("AMBER_OAUTH_SERVER", ac.licenseProfile.oauthServer)
	if ac.licenseProfile.oauthServer == "" {
		ac.licenseProfile.oauthServer = ac.licenseProfile.server
	}
	// verify required profile elements have been created
	if ac.licenseProfile.username == "" {
		return nil, errors.New("missing username in profile")
	}
	if ac.licenseProfile.password == "" {
		return nil, errors.New("missing username in profile")
	}
	if ac.licenseProfile.server == "" {
		return nil, errors.New("missing username in profile")
	}

	ac.oauthParams = &aops.PostOauth2Params{
		PostAuth2Request: &amodels.PostAuth2Request{
			Username: &ac.licenseProfile.username,
			Password: &ac.licenseProfile.password,
		},
	}

	/*
	   // set timeout in milliseconds
	   this.defaultClient.timeout = timeout * 1000

	   // set the proxy
	   this.defaultClient.proxy = process.env.AMBER_PROXY || null

	   // process overrides for the cert and verify
	   this.license_profile.cert = process.env.AMBER_SSL_CERT || cert
	   if (this.license_profile.cert !== null) {
	       console.log("cert specification not implemented yet")
	   }
	   this.license_profile.verify = verify
	   let verify_str = process.env.AMBER_SSL_VERIFY
	   if (verify_str && verify_str.toLowerCase() === "false") {
	       this.license_profile.verify = false
	   }
	   if (this.license_profile.verify === false) {
	       this.defaultClient.verifyTLS = this.license_profile.verify
	   }

	*/

	if timeout == nil {
		ac.timeout = 360
	} else {
		ac.timeout = *timeout
	}

	// set up default transport
	ac.amberServer = aclient.New(httptransport.New(ac.licenseProfile.server, "", nil), strfmt.Default)

	// set up oauth2 transport
	ac.oauthServer = aclient.New(httptransport.New(ac.licenseProfile.oauthServer, "", nil), strfmt.Default)

	return &ac, nil
}

func (a AmberClient) authenticate() bool {
	tIn := time.Now()
	if a.reauthTime.Before(tIn) {
		a.oauthServer.Operations
		response, err := a.oauthServer.ClientService.PostOauth2(a.oauthParams, nil)
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

	if (a.authenticate()) == false {
		return nil, errors.New("authentication failed")
	}

	params := &aops.GetSensorsParams{
		Context: nil,
		HTTPClient: &http.Client{
			Transport:     a.,
			Timeout:       time.Duration(time.Second * a.timeout),
		},
	}
	resp, err := a.amberServer.Operations(operations.AllParams{}, a.authWriter)
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
