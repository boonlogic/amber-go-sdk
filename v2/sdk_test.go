// Copyright 2018, Boon Logic Inc

package amber_client_v2

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	am "github.com/boonlogic/amber-go-sdk/v2/models"
	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	_ "log"
	"os"
)

// we need a global AmberClient and sensor id
var testClient *AmberClient
var testModelID string

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
		amberClient, _ = NewAmberClientFromProfile(licenseProfile)
	}

	return amberClient
}

func getUserSecrets(licenseId string) (LicenseProfile, error) {

	var lp licenseProfiles

	region := "us-east-1"
	svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("amber-test-licenses"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(*result.SecretString), &lp); err != nil {
		return LicenseProfile{}, err
	}

	profile := lp[licenseId]
	if profile.LicenseKey == "" {
		fmt.Printf("deployment %v not found\n", licenseId)
		os.Exit(3)
	}

	return profile, nil
}

func clearEnv() {
	os.Unsetenv("AMBER_LICENSE_KEY")
	os.Unsetenv("AMBER_SECRET_KEY")
	os.Unsetenv("AMBER_SERVER")
	os.Unsetenv("AMBER_OAUTH_SERVER")
}

// Runs before each test. Initializes by setting env variables.
func TestNewAmberClientFromProfile(t *testing.T) {
	clearEnv()

	// expected license profile
	profileA := LicenseProfile{
		LicenseKey:  "mylicensekey",
		SecretKey:   "mysecretkey",
		Server:      "https://fakeserver.com/v1",
		OauthServer: "https://fakeserver.com/v1/oauth",
	}
	amberClientA, err := NewAmberClientFromProfile(profileA)
	require.Nil(t, err)
	require.Equal(t, amberClientA.licenseProfile.LicenseKey, "mylicensekey")
	require.Equal(t, amberClientA.licenseProfile.SecretKey, "mysecretkey")
	require.Equal(t, amberClientA.licenseProfile.Server, "https://fakeserver.com/v1")
	require.Equal(t, amberClientA.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth")

	// expected license profile override from environment
	os.Setenv("AMBER_LICENSE_KEY", "mylicensekey-env")
	os.Setenv("AMBER_SECRET_KEY", "mysecretkey-env")
	os.Setenv("AMBER_SERVER", "https://fakeserver.com/v1/env")
	os.Setenv("AMBER_OAUTH_SERVER", "https://fakeserver.com/v1/oauth/env")
	amberClientB, err := NewAmberClientFromProfile(profileA)
	require.Nil(t, err)
	require.Equal(t, amberClientB.licenseProfile.LicenseKey, "mylicensekey-env")
	require.Equal(t, amberClientB.licenseProfile.SecretKey, "mysecretkey-env")
	require.Equal(t, amberClientB.licenseProfile.Server, "https://fakeserver.com/v1/env")
	require.Equal(t, amberClientB.licenseProfile.OauthServer, "https://fakeserver.com/v1/oauth/env")
}

func TestNewAmberClientFromFile(t *testing.T) {
	clearEnv()

	id := "default"
	file := "test/test.Amber.license"

	amberClientA, err := NewAmberClientFromFile(&id, &file)
	require.Nil(t, err)
	require.Equal(t, amberClientA.licenseProfile.LicenseKey, "amber-dev-user")
	require.Equal(t, amberClientA.licenseProfile.SecretKey, "phi{Ch2obovoe")
	require.Equal(t, amberClientA.licenseProfile.Server, "http://dev-build-1:5888/v1")
	require.Equal(t, amberClientA.licenseProfile.OauthServer, "https://amber-local.boonlogic.com/qa")

	// expected license profile override from environment
	os.Setenv("AMBER_LICENSE_KEY", "mylicensekey-env")
	os.Setenv("AMBER_SECRET_KEY", "mysecretkey-env")
	os.Setenv("AMBER_SERVER", "https://fakeserver.com/v1/env")
	os.Setenv("AMBER_OAUTH_SERVER", "https://fakeserver.com/v1/oauth/env")
	amberClientB, err := NewAmberClientFromFile(&id, &file)
	require.Nil(t, err)
	require.Equal(t, amberClientB.licenseProfile.LicenseKey, "mylicensekey-env")
	require.Equal(t, amberClientB.licenseProfile.SecretKey, "mysecretkey-env")
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
	saveSecretKey := amberClient.licenseProfile.SecretKey
	amberClient.licenseProfile.SecretKey = "bad"
	authResult, aErr := amberClient.authenticate()
	require.False(t, authResult)
	require.Equal(t, 401, int(aErr.Code))
	amberClient.licenseProfile.SecretKey = saveSecretKey

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

func TestModel(t *testing.T) {
	// the client and sensor that is created here will be used for the remainder of the tests
	testClient = createAmberClient()

	// test sensor creation
	modelLabel := "amber-go-sdk-create"
	response, aErr := testClient.PostModel(modelLabel)
	require.Nil(t, aErr)
	require.Equal(t, modelLabel, response.Label)
	testModelID = response.ID
}

func TestUpdateModel(t *testing.T) {
	// test sensor update
	modelLabel := "amber-go-sdk-test"
	response, aErr := testClient.PutModel(testModelID, modelLabel)
	require.Nil(t, aErr)
	require.Equal(t, modelLabel, response.Label)

	// test sensor update with invalid sensorID
	modelLabel = "amber-go-sdk-test"
	notAModel := "aaaaaaaaaaaaaaaaaaa"
	_, aErr = testClient.PutModel(notAModel, modelLabel)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Invalid ID", aErr.Message)
}

func TestGetModel(t *testing.T) {
	// test get sensor
	response, aErr := testClient.GetModel(testModelID)
	require.Nil(t, aErr)
	require.Equal(t, testModelID, response.ID)

	// test sensor update with invalid sensorID
	notAModel := "aaaaaaaaaaaaaaaaaaa"
	_, aErr = testClient.GetModel(notAModel)
	require.NotNil(t, aErr)
	require.Equal(t, 500, int(aErr.Code))
}

func TestListModels(t *testing.T) {
	// test list sensors
	response, aErr := testClient.GetModels()
	require.Nil(t, aErr)
	length := len(response.ModelList)
	require.Greater(t, length, 0)
}

func TestConfigureModel(t *testing.T) {
	// test sensor configuration
	fusionRule := "submit"
	fusionTTL := uint64(100)
	features := make([]*am.FeatureConfig, 1)
	features[0] = &am.FeatureConfig{
		FusionRule: &fusionRule,
		FusionTTL:  &fusionTTL,
		MaxVal:     1000.0,
		MinVal:     0.0,
		Name:       "feature_0",
	}

	var streamingWindowSize am.StreamingWindow = 25
	var pv am.PercentVariation = 0.05
	postConfigRequest := am.PostConfigRequest{
		Autotuning:       nil,
		Features:         features,
		StreamingWindow:  &streamingWindowSize,
		PercentVariation: &pv,
		Training:         nil,
	}
	response, aErr := testClient.PostModelConfig(testModelID, postConfigRequest)
	require.Nil(t, aErr)
	require.Equal(t, *postConfigRequest.StreamingWindow, response.StreamingWindow)
	require.Equal(t, len(postConfigRequest.Features), len(response.Features))
	require.Equal(t, 1, len(response.Features))

	// test sensor configuration with invalid sensorID
	notAModel := "aaaaaaaaaaaaaaaaaaa"
	_, aErr = testClient.PostModelConfig(notAModel, postConfigRequest)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Invalid ID", aErr.Message)

	// get the sensor config
	_, aErr = testClient.GetModelConfig(testModelID)
	require.Nil(t, aErr)

	// test get sensor config with invalid sensorID
	getConfigResponse, aErr := testClient.GetModelConfig(notAModel)
	require.NotNil(t, aErr)
	require.Nil(t, getConfigResponse)
	require.Equal(t, 500, int(aErr.Code))
}

func TestConfigureFusion(t *testing.T) {
	// setup: configure the sensor for fusion
	fc := uint16(5)
	var sw am.StreamingWindow = 1
	features := make([]*am.FeatureConfig, fc)

	fusionRule := "submit"
	fusionTTL := uint64(100)
	for i := 0; i < int(fc); i++ {
		label := fmt.Sprintf("f%d", i)
		features[i] = &am.FeatureConfig{
			FusionRule: &fusionRule,
			FusionTTL:  &fusionTTL,
			MaxVal:     1000.0,
			MinVal:     0.0,
			Name:       label,
		}
	}

	postConfigRequest := am.PostConfigRequest{Features: features, StreamingWindow: &sw}
	_, aErr := testClient.PostModelConfig(testModelID, postConfigRequest)
	if aErr != nil {
		panic(aErr)
	}

	// duplicate feature in configuration
	fbad := make([]*am.FeatureConfig, len(features))
	copy(fbad, features)
	fbad[3] = fbad[2]
	request := am.PostConfigRequest{Features: fbad, StreamingWindow: &sw}
	_, aErr = testClient.PostModelConfig(testModelID, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// unrecognized submit rule in configuration
	badsubmitrule := "badsubmitrule"
	copy(fbad, features)
	fbad[2].FusionRule = &badsubmitrule
	request = am.PostConfigRequest{Features: fbad, StreamingWindow: &sw}
	_, aErr = testClient.PostModelConfig(testModelID, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))
}

func TestStreamFusion(t *testing.T) {
	// stream partial vector (202 response)
	var l1, l3 = "f1", "f3"
	var v1, v3 float32 = 2, 4
	v := []*am.FusionFeature{{Name: &l1, Value: &v1}, {Name: &l3, Value: &v3}}
	expVec := map[string]float32{
		"f1": 2,
		"f3": 4,
	}
	request := am.PutDataRequest{Vector: v}
	response, aErr := testClient.PutModelData(testModelID, request)
	require.Nil(t, aErr)
	require.Equal(t, expVec, response.Vector)
	require.Nil(t, response.Analytics)

	// stream full vector (200 response)
	var l0, l2, l4 = "f0", "f2", "f4"
	var v0, v2, v4 float32 = 1, 3, 5
	v = []*am.FusionFeature{{Name: &l0, Value: &v0}, {Name: &l2, Value: &v2}, {Name: &l4, Value: &v4}}
	expVec = map[string]float32{
		"f0": 1,
		"f1": 2,
		"f2": 3,
		"f3": 4,
		"f4": 5,
	}
	request = am.PutDataRequest{Vector: v}
	response, aErr = testClient.PutModelData(testModelID, request)
	require.Nil(t, aErr)
	require.NotNil(t, response)
	require.NotNil(t, response.Analytics)
	require.Equal(t, expVec, response.Vector)

	// fusion vector contains label not in fusion configuration
	l0, l1 = "badfeature", "f3"
	v = []*am.FusionFeature{{Name: &l0, Value: &v0}, {Name: &l1, Value: &v1}}
	request = am.PutDataRequest{Vector: v}
	_, aErr = testClient.PutModelData(testModelID, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// fusion vector contains duplicate label
	v[0].Name = v[1].Name
	request = am.PutDataRequest{Vector: v}
	_, aErr = testClient.PutModelData(testModelID, request)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))

	// teardown: re-configure the sensor for single feature streaming
	fc := uint16(1)
	var sw am.StreamingWindow = 25
	fusionRule := "submit"
	fusionTTL := uint64(100)
	features := make([]*am.FeatureConfig, fc)
	features[0] = &am.FeatureConfig{
		FusionRule: &fusionRule,
		FusionTTL:  &fusionTTL,
		MaxVal:     1000.0,
		MinVal:     0.0,
		Name:       "Feature_0",
	}

	postConfigRequest := am.PostConfigRequest{Features: features, StreamingWindow: &sw}
	_, aErr = testClient.PostModelConfig(testModelID, postConfigRequest)
	if aErr != nil {
		panic(aErr)
	}
}

func TestPostStream(t *testing.T) {
	// stream the sensor
	data := "1.0,1.2,1.1,3.0"
	saveImage := true
	postDataRequest := am.PostDataRequest{
		Data:      &data,
		SaveImage: &saveImage,
	}
	postDataResponse, aErr := testClient.PostModelData(testModelID, postDataRequest)
	require.Nil(t, aErr)
	require.NotNil(t, postDataResponse)
	require.Equal(t, am.AmberState("Buffering"), postDataResponse.Status.State)
	require.Equal(t, 4, len(postDataResponse.Analytics.NI))
	require.Equal(t, 4, len(postDataResponse.Analytics.NS))
	require.Equal(t, 4, len(postDataResponse.Analytics.NW))
	require.Equal(t, 4, len(postDataResponse.Analytics.RI))
	require.Equal(t, 4, len(postDataResponse.Analytics.ID))

	// stream the sensor with invalid sensor id
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	postDataResponse, aErr = testClient.PostModelData(notASensor, postDataRequest)
	require.NotNil(t, aErr)
	require.Nil(t, postDataResponse)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Invalid ID", aErr.Message)
}

func TestPretrainModel(t *testing.T) {
	// teardown: re-configure the sensor for single feature streaming
	fc := uint16(1)
	var sw am.StreamingWindow = 25
	fusionRule := "submit"
	fusionTTL := uint64(100)
	features := make([]*am.FeatureConfig, fc)
	features[0] = &am.FeatureConfig{
		Weight:     1,
		FusionRule: &fusionRule,
		FusionTTL:  &fusionTTL,
		MaxVal:     1000.0,
		MinVal:     0.0,
		Name:       "feature_0",
	}
	autotune_pv, autotune_range := true, true
	autotuning := am.AutotuneConfig{
		PercentVariation: &autotune_pv,
		Range:            &autotune_range,
	}

	maxclusters := uint16(1000)
	bufferingsamples := uint32(10000)
	learningsamples := uint64(20000)
	window := uint32(1000)
	denominator := uint64(10000)
	numerator := uint64(1)
	training := am.TrainingConfig{
		BufferingSamples:        &bufferingsamples,
		HistoryWindow:           &window,
		LearningMaxClusters:     &maxclusters,
		LearningMaxSamples:      &learningsamples,
		LearningRateDenominator: &denominator,
		LearningRateNumerator:   &numerator,
	}

	postConfigRequest := am.PostConfigRequest{Autotuning: &autotuning, Features: features, StreamingWindow: &sw, Training: &training}
	_, aErr := testClient.PostModelConfig(testModelID, postConfigRequest)
	if aErr != nil {
		panic(aErr)
	}

	// test get pretrain
	getPretrainResponse, aErr := testClient.GetPretrainState(testModelID)
	require.Nil(t, aErr)
	require.NotNil(t, getPretrainResponse)
	require.Equal(t, "", getPretrainResponse.Message)
	require.Equal(t, "None", getPretrainResponse.Status)

	// stream the sensor with invalid sensor id
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	getPretrainResponse, aErr = testClient.GetPretrainState(notASensor)
	require.NotNil(t, aErr)
	require.Nil(t, getPretrainResponse)
	require.Equal(t, 500, int(aErr.Code))

	// read entire data csv
	csvRecords, err := LoadCsvRecords("../examples/v2/output_current.csv")
	require.Nil(t, err)

	// generate one csv string from csv records
	pretrainData := strings.Join(csvRecords, ",")

	// pretrain the sensor with
	postPretrainRequest := am.PostPretrainRequest{
		Data: &pretrainData,
	}
	postPretrainResponse, aErr := testClient.PretrainModel(testModelID, postPretrainRequest)
	require.Nil(t, aErr)
	require.NotNil(t, postPretrainResponse)
	require.Equal(t, "", postPretrainResponse.Message)
	// some implementations of amber block until finished with request.  The state moves directly to "Pretrained"
	pretrainState := postPretrainResponse.Status
	require.True(t, pretrainState == "Pretrained" || pretrainState == "Pretraining")

	for pretrainState == "Pretraining" {
		// wait for 5 seconds between checking pretraining state
		time.Sleep(5 * time.Second)
		getPretrainResponse, aErr = testClient.GetPretrainState(testModelID)
		require.Nil(t, aErr)
		require.NotNil(t, getPretrainResponse)
		pretrainState = getPretrainResponse.Status
		require.True(t, pretrainState == "Pretraining" || pretrainState == "Pretrained" || pretrainState == "None")
	}
	require.Equal(t, "Pretrained", pretrainState)
}

func TestGetRootCause(t *testing.T) {
	// test get rootcause error (neigher param specified)
	clusterIds := []int32{}
	patterns := [][]float32{}
	getRootCauseResponse, aErr := testClient.GetRootCause(testModelID, clusterIds, patterns)
	require.NotNil(t, aErr)
	require.Nil(t, getRootCauseResponse)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Must specify either patterns or cluster IDs for analysis", aErr.Message)

	// test get rootcause error (both params specified)
	clusterIds = []int32{1}
	patterns = [][]float32{{1.0}}
	getRootCauseResponse, aErr = testClient.GetRootCause(testModelID, clusterIds, patterns)
	require.NotNil(t, aErr)
	require.Nil(t, getRootCauseResponse)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Cannot specify both patterns and cluster IDs for analysis", aErr.Message)

	// test get rootcause for success
	patterns = [][]float32{}
	getRootCauseResponse, aErr = testClient.GetRootCause(testModelID, clusterIds, patterns)
	require.Nil(t, aErr)
	require.NotNil(t, getRootCauseResponse)

	// stream the sensor with invalid sensor id
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	getRootCauseResponse, aErr = testClient.GetRootCause(notASensor, clusterIds, patterns)
	require.NotNil(t, aErr)
	require.Nil(t, getRootCauseResponse)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Invalid ID", aErr.Message)
}

func TestGetStatus(t *testing.T) {
	// test get status
	response, aErr := testClient.GetModelStatus(testModelID)
	require.Nil(t, aErr)
	require.NotNil(t, response)
	require.Equal(t, am.AmberState("Monitoring"), response.State)

	// test get status with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	_, aErr = testClient.GetModelStatus(notASensor)
	require.NotNil(t, aErr)
	require.Equal(t, 500, int(aErr.Code))
}

func TestEnableLearning(t *testing.T) {
	// get the sensor config
	_, aErr := testClient.GetModelConfig(testModelID)
	require.Nil(t, aErr)

	maxclusters := uint16(1000)
	learningsamples := uint64(1000)
	denominator := uint64(10000)
	numerator := uint64(1)
	training := &am.TrainingConfig{
		LearningMaxClusters:     &maxclusters,
		LearningMaxSamples:      &learningsamples,
		LearningRateDenominator: &denominator,
		LearningRateNumerator:   &numerator,
	}
	state := "Learning"
	putConfigRequest := am.PostLearningRequest{
		Training: training,
		State:    &state,
	}

	response, aErr := testClient.PostModelLearning(testModelID, putConfigRequest)
	require.Nil(t, aErr)
	require.Equal(t, *putConfigRequest.Training.LearningMaxSamples, *response.Training.LearningMaxSamples)
	require.Equal(t, *putConfigRequest.Training.LearningRateNumerator, *response.Training.LearningRateNumerator)

	// test sensor configuration with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	_, aErr = testClient.PostModelLearning(notASensor, putConfigRequest)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))
	require.Equal(t, "Invalid ID", aErr.Message)

	// failed because not in learning/monitoring
	fc := uint16(1)
	var sw am.StreamingWindow = 25
	fusionRule := "submit"
	fusionTTL := uint64(100)
	bufferingsamples := uint32(1000)
	features := make([]*am.FeatureConfig, fc)
	features[0] = &am.FeatureConfig{
		FusionRule: &fusionRule,
		FusionTTL:  &fusionTTL,
		MaxVal:     1000.0,
		MinVal:     0.0,
		Name:       "feature_0",
	}

	postConfigRequest := am.PostConfigRequest{
		Features:        features,
		StreamingWindow: &sw,
		Training: &am.TrainingConfig{
			BufferingSamples:        &bufferingsamples,
			LearningMaxClusters:     &maxclusters,
			LearningMaxSamples:      &learningsamples,
			LearningRateDenominator: &denominator,
			LearningRateNumerator:   &numerator,
		},
	}
	_, aErr = testClient.PostModelConfig(testModelID, postConfigRequest)
	require.Nil(t, aErr)
	_, aErr = testClient.PostModelLearning(testModelID, putConfigRequest)
	require.NotNil(t, aErr)
	require.Equal(t, 400, int(aErr.Code))
}

func TestPostOutage(t *testing.T) {
	aErr := testClient.PostModelOutage(testModelID)
	require.Nil(t, aErr)

	// test sensor configuration with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	aErr = testClient.PostModelOutage(notASensor)
	require.NotNil(t, aErr)
	require.Equal(t, 500, int(aErr.Code))
}

func TestDeleteSensor(t *testing.T) {
	// test delete sensor
	aErr := testClient.DeleteModel(testModelID)
	require.Nil(t, aErr)

	// test delete sensor with invalid sensorID
	notASensor := "aaaaaaaaaaaaaaaaaaa"
	aErr = testClient.DeleteModel(notASensor)
	require.NotNil(t, aErr)
	require.Equal(t, 500, int(aErr.Code))
}

func TestGetVersion(t *testing.T) {
	// test get version
	response, aErr := testClient.GetVersion()
	require.Nil(t, aErr)
	require.NotNil(t, response)
	require.Equal(t, "/v2", response.APIVersion)
}
