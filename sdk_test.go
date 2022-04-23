// Copyright 2018, Boon Logic Inc

package amber_client

import (
	"encoding/json"
	"fmt"
	am "github.com/boonlogic/amber-go-sdk/models"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	_ "log"
	"os"
)

// we need a global AmberClient and sensor id
var testClient *AmberClient
var testSensor string

// secrets downloaded from points beyond
func createAmberClient() *AmberClient {
	amberLicenseFile := os.Getenv("AMBER_TEST_LICENSE_FILE")
	amberLicenseId := os.Getenv("AMBER_TEST_LICENSE_ID")
	if amberLicenseId == "" {
		panic("AMBER_TEST_LICENSE_ID is missing in test environment")
	}

	// purge AMBER environment variables for key in Test_01_AmberInstance.saved_env.keys():
	clearEnv()

	var licenseProfile LicenseProfile
	var amberClient *AmberClient
	var err error

	if amberLicenseFile != "" {
		// load license profile using a local license file
		amberClient, err = NewAmberClientFromFile(&amberLicenseId, &amberLicenseFile)
		if err != nil {
			panic(err)
		}
	} else {
		// load license profile from secrets manager
		licenseProfile, err = getUserSecrets(amberLicenseId)
		if err != nil {
			panic(err)
		}
		amberClient, err = NewAmberClientFromProfile(licenseProfile)
	}

	return amberClient
}

func getUserSecrets(licenseId string) (LicenseProfile, error) {

	var lp licenseProfiles

	region := "us-east-1"
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("amber-test-users"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(*result.SecretString), &lp); err != nil {
		return LicenseProfile{}, err
	}

	profile := lp[licenseId]
	if profile.Username == "" {
		fmt.Printf("deployment %v not found\n", licenseId)
		os.Exit(3)
	}

	return profile, nil
}

func clearEnv() {
	os.Unsetenv("AMBER_USERNAME")
	os.Unsetenv("AMBER_PASSWORD")
	os.Unsetenv("AMBER_SERVER")
	os.Unsetenv("AMBER_OAUTH_SERVER")
}

// Runs before each test. Initializes by setting env variables.
func TestNewAmberClientFromProfile(t *testing.T) {
	clearEnv()

	// expected license profile
	profileA := LicenseProfile{
		Username:    "mruser",
		Password:    "mypassword",
		Server:      "https://fakeserver.com/v1",
		OauthServer: "https://fakeserver.com/v1/oauth",
	}
	amberClientA, err := NewAmberClientFromProfile(profileA)
	require.Nil(t, err)
	require.Equal(t, amberClientA.licenseProfile.Username, "mruser")
	require.Equal(t, amberClientA.licenseProfile.Password, "mypassword")
	require.Equal(t, amberClientA.licenseProfile.Server, "https://fakeserver.com/v1")
	require.Equal(t, amberClientA.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth")

	// expected license profile override from environment
	os.Setenv("AMBER_USERNAME", "mruser-env")
	os.Setenv("AMBER_PASSWORD", "mypassword-env")
	os.Setenv("AMBER_SERVER", "https://fakeserver.com/v1/env")
	os.Setenv("AMBER_OAUTH_SERVER", "https://fakeserver.com/v1/oauth/env")
	amberClientB, err := NewAmberClientFromProfile(profileA)
	require.Nil(t, err)
	require.Equal(t, amberClientB.licenseProfile.Username, "mruser-env")
	require.Equal(t, amberClientB.licenseProfile.Password, "mypassword-env")
	require.Equal(t, amberClientB.licenseProfile.Server, "https://fakeserver.com/v1/env")
	require.Equal(t, amberClientB.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth/env")
}

func TestNewAmberClientFromFile(t *testing.T) {
	clearEnv()

	id := "default"
	file := "test/test.Amber.license"

	amberClientA, err := NewAmberClientFromFile(&id, &file)
	require.Nil(t, err)
	require.Equal(t, amberClientA.licenseProfile.Username, "amber-dev-user")
	require.Equal(t, amberClientA.licenseProfile.Password, "phi{Ch2obovoe")
	require.Equal(t, amberClientA.licenseProfile.Server, "http://dev-build-1:5888/v1")
	require.Equal(t, amberClientA.licenseProfile.OauthServer, "https://amber-local.boonlogic.com/qa")

	// expected license profile override from environment
	os.Setenv("AMBER_USERNAME", "mruser-env")
	os.Setenv("AMBER_PASSWORD", "mypassword-env")
	os.Setenv("AMBER_SERVER", "https://fakeserver.com/v1/env")
	os.Setenv("AMBER_OAUTH_SERVER", "https://fakeserver.com/v1/oauth/env")
	amberClientB, err := NewAmberClientFromFile(&id, &file)
	require.Nil(t, err)
	require.Equal(t, amberClientB.licenseProfile.Username, "mruser-env")
	require.Equal(t, amberClientB.licenseProfile.Password, "mypassword-env")
	require.Equal(t, amberClientB.licenseProfile.Server, "https://fakeserver.com/v1/env")
	require.Equal(t, amberClientB.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth/env")
}

func TestNewAmberClientNegative(t *testing.T) {

	id := "default"
	file := "nonexistent-license-file"

	// should error when specified license file does not exist
	clearEnv()
	_, err := NewAmberClientFromFile(&id, &file)
	require.Equal(t, err.Error(), "license file 'nonexistent-license-file' not found")

	// should error when specified license id does not exist
	clearEnv()
	id = "nonexistent-license-id"
	file = "test/test.Amber.license"
	_, err = NewAmberClientFromFile(&id, &file)
	require.Equal(t, err.Error(), "license id 'nonexistent-license-id' not found")
}

func TestAuthenticate(t *testing.T) {
	// should error when specified license file does not exist
	clearEnv()

	amberClient := createAmberClient()

	// test authentication failure
	savePassword := amberClient.licenseProfile.Password
	amberClient.licenseProfile.Password = "bad"
	authResult, aErr := amberClient.authenticate()
	require.False(t, authResult)
	require.Equal(t, 401, int(aErr.Code))
	amberClient.licenseProfile.Password = savePassword

	// test authentication success (will set jwt token)
	authResult, aErr = amberClient.authenticate()
	require.True(t, authResult)
	require.Nil(t, aErr)

	// test authentication when jwt token is expired.  setting it to Now should
	// trigger a reauthentication since it is within the 1 minute reauth window
	timeNow := time.Now()
	amberClient.reauthTime = timeNow
	authResult, aErr = amberClient.authenticate()
	require.True(t, authResult)
	require.Nil(t, aErr)
	require.True(t, timeNow.Before(amberClient.reauthTime))
}

func TestSensor(t *testing.T) {
	// the client and sensor that is created here will be used for the remainder of the tests
	testClient = createAmberClient()

	// test sensor creation
	sensorLabel := "amber-go-sdk-create"
	response, aErr := testClient.CreateSensor(sensorLabel)
	require.Nil(t, aErr)
	require.Equal(t, sensorLabel, *response.Label)
	testSensor = *response.SensorID
}

func TestUpdateSensor(t *testing.T) {
	// test sensor update
	sensorLabel := "amber-go-sdk-test"
	response, aErr := testClient.UpdateLabel(testSensor, sensorLabel)
	require.Nil(t, aErr)
	require.Equal(t, sensorLabel, *response.Label)

	// test sensor update with invalid sensorID
	sensorLabel = "amber-go-sdk-test"
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	response, aErr = testClient.UpdateLabel(notASensor, sensorLabel)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestGetSensor(t *testing.T) {
	// test get sensor
	response, aErr := testClient.GetSensor(testSensor)
	require.Nil(t, aErr)
	require.Equal(t, testSensor, *response.SensorID)

	// test sensor update with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	response, aErr = testClient.GetSensor(notASensor)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestListSensors(t *testing.T) {
	// test list sensors
	response, aErr := testClient.ListSensors()
	require.Nil(t, aErr)
	length := len(*response)
	require.Greater(t, length, 0)
}

func TestConfigureSensor(t *testing.T) {
	// test sensor configuration
	var featureCount uint16 = 1
	var streamingWindowSize uint16 = 25
	postConfigRequest := am.PostConfigRequest{
		FeatureCount:        &featureCount,
		Features:            nil,
		SamplesToBuffer:     nil,
		StreamingWindowSize: &streamingWindowSize,
		StreamingParameters: am.StreamingParameters{
			LearningMaxClusters:     nil,
			LearningMaxSamples:      nil,
			LearningRateDenominator: nil,
			LearningRateNumerator:   nil,
			AnomalyHistoryWindow:    nil,
		},
	}
	response, aErr := testClient.ConfigureSensor(testSensor, postConfigRequest)
	require.Nil(t, aErr)
	require.Equal(t, postConfigRequest.StreamingWindowSize, response.StreamingWindowSize)
	require.Equal(t, postConfigRequest.FeatureCount, response.FeatureCount)
	require.Equal(t, 1, len(response.Features))

	// test sensor configuration with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	response, aErr = testClient.ConfigureSensor(notASensor, postConfigRequest)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)

	// get the sensor config
	getConfigResponse, aErr := testClient.GetConfig(testSensor)
	require.Nil(t, aErr)

	// test get sensor config with invalid sensorID
	getConfigResponse, aErr = testClient.GetConfig(notASensor)
	require.NotNil(t, aErr)
	require.Nil(t, getConfigResponse)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestConfigureFusion(t *testing.T) {
	// setup: configure the sensor for fusion
	fc := uint16(5)
	sw := uint16(1)
	postConfigRequest := am.PostConfigRequest{FeatureCount: &fc, StreamingWindowSize: &sw}
	_, aErr := testClient.ConfigureSensor(testSensor, postConfigRequest)
	if aErr != nil {
		panic(aErr)
	}

	// test fusion configuration
	f := make([]*am.FusionConfig, fc)
	for i := 0; i < int(fc); i++ {
		l := fmt.Sprintf("f%d", i)
		f[i] = &am.FusionConfig{Label: &l, SubmitRule: "submit"}
	}
	request := am.PutConfigRequest{Features: f}
	response, aErr := testClient.ConfigureFusion(testSensor, request)
	require.Nil(t, aErr)
	require.Equal(t, request.Features, response.Features)

	// test fusion configuration with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	response, aErr = testClient.ConfigureFusion(notASensor, request)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)

	// number of features doesn't match configured feature count
	l := "f5"
	fplus := append(f, &am.FusionConfig{Label: &l})
	request = am.PutConfigRequest{Features: fplus}
	response, aErr = testClient.ConfigureFusion(testSensor, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// duplicate feature in configuration
	fbad := make([]*am.FusionConfig, len(f))
	copy(fbad, f)
	fbad[3] = fbad[2]
	request = am.PutConfigRequest{Features: fbad}
	response, aErr = testClient.ConfigureFusion(testSensor, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// unrecognized submit rule in configuration
	copy(fbad, f)
	fbad[2].SubmitRule = "badsubmitrule"
	request = am.PutConfigRequest{Features: fbad}
	response, aErr = testClient.ConfigureFusion(testSensor, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))
}

func TestStreamFusion(t *testing.T) {
	// stream partial vector (202 response)
	var l1, l3 = "f1", "f3"
	var v1, v3 float32 = 2, 4
	v := []*am.PutStreamFeature{{Label: &l1, Value: &v1}, {Label: &l3, Value: &v3}}
	expVec := am.MayContainNullsArray{nil, json.Number("2"), nil, json.Number("4"), nil}
	request := am.PutStreamRequest{Vector: v}
	response, aErr := testClient.StreamFusion(testSensor, request)
	require.Nil(t, aErr)
	require.Equal(t, expVec, response.Vector)
	require.Nil(t, response.Results)
	require.Nil(t, response.VectorCSV)

	// stream full vector (200 response)
	var l0, l2, l4 = "f0", "f2", "f4"
	var v0, v2, v4 float32 = 1, 3, 5
	v = []*am.PutStreamFeature{{Label: &l0, Value: &v0}, {Label: &l2, Value: &v2}, {Label: &l4, Value: &v4}}
	expVec = am.MayContainNullsArray{json.Number("1"), json.Number("2"), json.Number("3"), json.Number("4"), json.Number("5")}
	request = am.PutStreamRequest{Vector: v}
	response, aErr = testClient.StreamFusion(testSensor, request)
	require.Nil(t, aErr)
	require.NotNil(t, response)
	require.NotNil(t, response.Results)
	require.Nil(t, response.VectorCSV)
	require.Equal(t, expVec, response.Vector)

	// fusion vector contains label not in fusion configuration
	l0, l1 = "badfeature", "f3"
	v = []*am.PutStreamFeature{{Label: &l0, Value: &v0}, {Label: &l1, Value: &v1}}
	request = am.PutStreamRequest{Vector: v}
	response, aErr = testClient.StreamFusion(testSensor, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// fusion vector contains duplicate label
	v[0].Label = v[1].Label
	request = am.PutStreamRequest{Vector: v}
	response, aErr = testClient.StreamFusion(testSensor, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// teardown: re-configure the sensor for single feature streaming
	fc := uint16(1)
	sw := uint16(25)
	postConfigRequest := am.PostConfigRequest{FeatureCount: &fc, StreamingWindowSize: &sw}
	_, aErr = testClient.ConfigureSensor(testSensor, postConfigRequest)
	if aErr != nil {
		panic(aErr)
	}
}

func TestPostStream(t *testing.T) {
	// stream the sensor
	data := "1.0,1.2,1.1,3.0"
	saveImage := true
	postStreamRequest := am.PostStreamRequest{
		Data:      &data,
		SaveImage: &saveImage,
	}
	postStreamResponse, aErr := testClient.StreamSensor(testSensor, postStreamRequest)
	require.Nil(t, aErr)
	require.NotNil(t, postStreamResponse)
	require.Equal(t, "Buffering", *postStreamResponse.StreamStatus.State)
	require.Equal(t, 4, len(postStreamResponse.AH))
	require.Equal(t, 4, len(postStreamResponse.AM))
	require.Equal(t, 4, len(postStreamResponse.AW))
	require.Equal(t, 4, len(postStreamResponse.AD))
	require.Equal(t, 4, len(postStreamResponse.ID))

	// stream the sensor with invalid sensor id
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	postStreamResponse, aErr = testClient.StreamSensor(notASensor, postStreamRequest)
	require.NotNil(t, aErr)
	require.Nil(t, postStreamResponse)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestPretrainSensor(t *testing.T) {
	// test get pretrain
	getPretrainResponse, aErr := testClient.GetPretrainState(testSensor)
	require.Nil(t, aErr)
	require.NotNil(t, getPretrainResponse)
	require.Equal(t, "", getPretrainResponse.Message)
	require.Equal(t, "None", *getPretrainResponse.State)

	// stream the sensor with invalid sensor id
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	getPretrainResponse, aErr = testClient.GetPretrainState(notASensor)
	require.NotNil(t, aErr)
	require.Nil(t, getPretrainResponse)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)

	// read entire data csv
	csvRecords, err := LoadCsvRecords("examples/output_current.csv")
	require.Nil(t, err)

	// generate one csv string from csv records
	pretrainData := strings.Join(csvRecords, ",")

	// pretrain the sensor with
	autoTuneConfig := true
	postPretrainRequest := am.PostPretrainRequest{
		AutotuneConfig: &autoTuneConfig,
		Data:           &pretrainData,
	}
	postPretrainResponse, aErr := testClient.PretrainSensor(testSensor, postPretrainRequest)
	require.Nil(t, aErr)
	require.NotNil(t, postPretrainResponse)
	require.Equal(t, "", postPretrainResponse.Message)
	// some implementations of amber block until finished with request.  The state moves directly to "Pretrained"
	pretrainState := *postPretrainResponse.State
	require.True(t, pretrainState == "Pretrained" || pretrainState == "Pretraining")

	for pretrainState == "Pretraining" {
		// wait for 5 seconds between checking pretraining state
		time.Sleep(5 * time.Second)
		getPretrainResponse, aErr = testClient.GetPretrainState(testSensor)
		require.Nil(t, aErr)
		require.NotNil(t, getPretrainResponse)
		pretrainState = *getPretrainResponse.State
		require.True(t, pretrainState == "Pretraining" || pretrainState == "Pretrained" || pretrainState == "None")
	}
	require.Equal(t, "Pretrained", pretrainState)
}

func TestGetRootCause(t *testing.T) {
	// test get rootcause
	clusterIds := "[1]"
	patterns := ""
	getRootCauseResponse, aErr := testClient.GetRootCause(testSensor, &clusterIds, &patterns)
	require.Nil(t, aErr)
	require.NotNil(t, getRootCauseResponse)

	// stream the sensor with invalid sensor id
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	getRootCauseResponse, aErr = testClient.GetRootCause(notASensor, &clusterIds, &patterns)
	require.NotNil(t, aErr)
	require.Nil(t, getRootCauseResponse)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestGetStatus(t *testing.T) {
	// test get status
	response, aErr := testClient.GetStatus(testSensor)
	require.Nil(t, aErr)
	require.NotNil(t, response)
	require.Equal(t, "Monitoring", *response.State)

	// test get status with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	response, aErr = testClient.GetStatus(notASensor)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestEnableLearning(t *testing.T) {
	putConfigRequest := am.PutConfigRequest{
		Streaming: &am.StreamingParameters{
			LearningMaxClusters:     nil,
			LearningMaxSamples:      nil,
			LearningRateDenominator: nil,
			LearningRateNumerator:   nil,
			AnomalyHistoryWindow:    nil,
		},
	}
	response, aErr := testClient.EnableLearning(testSensor, putConfigRequest)
	require.Nil(t, aErr)
	require.Equal(t, putConfigRequest.Streaming.AnomalyHistoryWindow, response.Streaming.AnomalyHistoryWindow)
	require.Equal(t, putConfigRequest.Streaming.LearningMaxSamples, response.Streaming.LearningMaxSamples)

	// test sensor configuration with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	response, aErr = testClient.EnableLearning(notASensor, putConfigRequest)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)

	// failed because not in learning/monitoring
	var featureCount uint16 = 1
	var streamingWindowSize uint16 = 25
	postConfigRequest := am.PostConfigRequest{
		FeatureCount:        &featureCount,
		Features:            nil,
		SamplesToBuffer:     nil,
		StreamingWindowSize: &streamingWindowSize,
		StreamingParameters: am.StreamingParameters{
			LearningMaxClusters:     nil,
			LearningMaxSamples:      nil,
			LearningRateDenominator: nil,
			LearningRateNumerator:   nil,
			AnomalyHistoryWindow:    nil,
		},
	}
	_, aErr = testClient.ConfigureSensor(testSensor, postConfigRequest)
	require.Nil(t, aErr)
	response, aErr = testClient.EnableLearning(testSensor, putConfigRequest)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

}

func TestDeleteSensor(t *testing.T) {
	// test delete sensor
	aErr := testClient.DeleteSensor(testSensor)
	require.Nil(t, aErr)

	// test delete sensor with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	aErr = testClient.DeleteSensor(notASensor)
	require.NotNil(t, aErr)
	require.Equal(t, 404, int(aErr.Code))
	require.Equal(t, "sensor aaaaaaaaaaaaaaaaaaa not found", aErr.Message)
}

func TestGetVersion(t *testing.T) {
	// test get version
	response, aErr := testClient.GetVersion()
	require.Nil(t, aErr)
	require.NotNil(t, response)
	require.Equal(t, "/v1", *response.APIVersion)
}
