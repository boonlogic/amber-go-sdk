// Copyright 2018, Boon Logic Inc

package amber_client

import (
	"encoding/json"
	"fmt"
	amberModels "github.com/boonlogic/amber-go-sdk/models"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"

	// "fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	// "io/ioutil"
	_ "log"
	"os"
)

var licenseProfile LicenseProfile
var testClient *AmberClient
var testSensor string

func init() {
	var err error
	licenseProfile, err = getUserSecrets()
	if err != nil {
		fmt.Printf("getUserSecrets failed\n")
		os.Exit(3)
	}
}

func getUserSecrets() (LicenseProfile, error) {

	var lp licenseProfiles

	// retrieve the deployment from the environment.  If not set, default to 'qa'
	deployment := os.Getenv("AMBER_TEST_PROFILE")
	if deployment == "" {
		deployment = "qa"
	}

	region := "us-east-1"
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
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

	profile := lp[deployment]
	if profile.Username == "" {
		fmt.Printf("deployment %v not found\n", deployment)
		os.Exit(3)
	}

	return profile, nil
}

func loadCredentialsIntoEnv() {
	os.Setenv("AMBER_USERNAME", licenseProfile.Username)
	os.Setenv("AMBER_PASSWORD", licenseProfile.Password)
	os.Setenv("AMBER_SERVER", licenseProfile.Server)
	os.Setenv("AMBER_OAUTH_SERVER", licenseProfile.OauthServer)
}

func clearEnv() {
	os.Unsetenv("AMBER_USERNAME")
	os.Unsetenv("AMBER_PASSWORD")
	os.Unsetenv("AMBER_SERVER")
	os.Unsetenv("AMBER_OAUTH_SERVER")
}

func restoreEnv(savedProfile LicenseProfile) {
	clearEnv()
	os.Setenv("AMBER_USERNAME", savedProfile.Username)
	os.Setenv("AMBER_PASSWORD", savedProfile.Password)
	os.Setenv("AMBER_SERVER", savedProfile.Server)
	os.Setenv("AMBER_OAUTH_SERVER", savedProfile.OauthServer)
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

	amberClient, err := NewAmberClientFromProfile(licenseProfile)
	require.Nil(t, err)

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

func TestSensor(T *testing.T) {

	// the client and sensor that is created here will be used for the remainder of the tests
	var err error
	testClient, err = NewAmberClientFromProfile(licenseProfile)
	require.Nil(T, err)

	// test sensor creation
	sensorLabel := "amber-go-sdk-create"
	response, aErr := testClient.CreateSensor(sensorLabel)
	require.Nil(T, aErr)
	require.Equal(T, sensorLabel, *response.Label)
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
	postConfigRequest := amberModels.PostConfigRequest{
		AnomalyHistoryWindow:    nil,
		FeatureCount:            &featureCount,
		Features:                nil,
		LearningMaxClusters:     nil,
		LearningMaxSamples:      nil,
		LearningRateDenominator: nil,
		LearningRateNumerator:   nil,
		SamplesToBuffer:         nil,
		StreamingWindowSize:     &streamingWindowSize,
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

func TestPostStream(t *testing.T) {

	// stream the sensor
	data := "1.0,1.2,1.1,3.0"
	saveImage := true
	postStreamRequest := amberModels.PostStreamRequest{
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
	postPretrainRequest := amberModels.PostPretrainRequest{
		AutotuneConfig: &autoTuneConfig,
		Data:           &pretrainData,
	}
	postPretrainResponse, aErr := testClient.PretrainSensor(testSensor, postPretrainRequest)
	require.Nil(t, aErr)
	require.NotNil(t, postPretrainResponse)
	require.Equal(t, "", postPretrainResponse.Message)
	require.Equal(t, "Pretraining", *postPretrainResponse.State)

	pretrainState := *postPretrainResponse.State
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
