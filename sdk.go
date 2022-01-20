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
	if profile.OauthServer == "" {
		profile.OauthServer = profile.Server
	}
	client.licenseProfile = profile

	// override from environment
	client.loadFromEnv()

	if profile.Username == "" {
		return nil, errors.New("missing username in profile")
	}
	if profile.Password == "" {
		return nil, errors.New("missing password in profile")
	}
	if profile.Server == "" {
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
		blob, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(blob, &lp); err != nil {
			return nil, err
		}
		profile = lp[id]
	}
	return NewAmberClientFromProfile(profile)
}

func (a *AmberClient) SetNoVerify(verify bool) error {
	a.verify = !verify
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

func (a *AmberClient) GetSensor(sensorId string) (*amberModels.GetSensorResponse, error) {
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

func (a *AmberClient) CreateSensor(label string) (*amberModels.PostSensorResponse, error) {
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

func (a *AmberClient) UpdateLabel(sensorId string, label string) (*amberModels.PutSensorResponse, error) {
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

func (a *AmberClient) ConfigureSensor(sensorId string, payload amberModels.PostConfigRequest) (*amberModels.PostConfigResponse, error) {
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

func (a *AmberClient) GetConfig(sensorId string) (*amberModels.GetConfigResponse, error) {
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

func (a *AmberClient) DeleteSensor(sensorId string) error {
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

func (a *AmberClient) StreamSensor(sensorId string, payload amberModels.PostStreamRequest) (*amberModels.PostStreamResponse, error) {
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

func (a *AmberClient) GetStatus(sensorId string) (*amberModels.GetStatusResponse, error) {
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

func (a *AmberClient) PretrainSensor(sensorId string, payload amberModels.PostPretrainRequest) (*amberModels.PostPretrainResponse, error) {
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

func (a *AmberClient) GetPretrainState(sensorId string) (*amberModels.GetPretrainResponse, error) {
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

func (a *AmberClient) GetRootCause(sensorId string, clusterId *string, pattern *string) (*amberModels.GetRootCauseResponse, error) {
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

func (a *AmberClient) GetVersion() (*amberModels.Version, error) {
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

func (a *AmberClient) updateHttpClients() {

	// set default verify and cert
	a.tlsOptions.InsecureSkipVerify = !a.verify
	a.tlsOptions.Certificate = a.cert

	// set server http client
	scheme, host, basePath, _ := parseServer(a.licenseProfile.Server)
	if scheme == "https" {
		httpClient, _ := httptransport.TLSClient(a.tlsOptions)
		a.amberServer = amberClient.New(httptransport.NewWithClient(host, basePath, []string{scheme}, httpClient), strfmt.Default)
	} else {
		a.amberServer = amberClient.New(httptransport.New(host, basePath, []string{scheme}), strfmt.Default)
	}

	// set up oauth http client
	_, host, basePath, _ = parseServer(a.licenseProfile.OauthServer)
	if scheme == "https" {
		httpClient, _ := httptransport.TLSClient(a.tlsOptions)
		a.oauthServer = amberClient.New(httptransport.NewWithClient(host, basePath, []string{scheme}, httpClient), strfmt.Default)
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
