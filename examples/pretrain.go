package main

import (
	"encoding/json"
	"fmt"
	amberClient "github.com/boonlogic/amber-go-sdk"
	amberModels "github.com/boonlogic/amber-go-sdk/models"
	"strings"
	"time"
)

func main() {

	// create the client
	ac, err := amberClient.NewAmberClientFromFile(nil, nil)
	if err != nil {
		panic(err)
	}

	// Fill in the sensorId here if one is already known.
	sensorId := ""
	// Set configureSensor to false to skip reconfiguring sensor
	configureSensor := true

	if sensorId == "" {
		// create a sensor when none specified
		createSensorResponse, aErr := ac.CreateSensor("new-go-sdk-sensor")
		if aErr != nil {
			panic(aErr)
		}
		sensorId = *createSensorResponse.SensorID
		fmt.Printf("created sensor id: %v\n", sensorId)
	} else {
		fmt.Printf("using sensor id: %v\n", sensorId)
	}

	if configureSensor {
		// configure the sensor
		featureCount := uint16(1)
		streamingWindowSize := uint16(25)
		learningMaxSamples := uint64(3000)
		samplesToBuffer := uint32(1000)
		learningRateNumerator := uint64(1)
		learningRateDenominator := uint64(1000)
		configRequest := amberModels.PostConfigRequest{
			FeatureCount:            &featureCount,
			StreamingWindowSize:     &streamingWindowSize,
			LearningMaxSamples:      &learningMaxSamples,
			LearningRateNumerator:   &learningRateNumerator,
			LearningRateDenominator: &learningRateDenominator,
			SamplesToBuffer:         &samplesToBuffer,
		}
		configSensorResponse, aErr := ac.ConfigureSensor(sensorId, configRequest)
		if aErr != nil {
			panic(aErr)
		}
		formatted, err := json.MarshalIndent(*configSensorResponse, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", string(formatted))
	}

	// read in entire csv file
	records, _ := amberClient.LoadCsvRecords("examples/output_current.csv")

	// pretrain the request
	pretrainData := strings.Join(records, ",")
	autotuneConfig := true
	pretrainRequest := amberModels.PostPretrainRequest{
		AutotuneConfig: &autotuneConfig,
		Data:           &pretrainData,
	}

	pretrainResponse, aErr := ac.PretrainSensor(sensorId, pretrainRequest)
	if aErr != nil {
		panic(err)
	}
	result := *pretrainResponse.State
	for result == "Pretraining" {
		fmt.Printf("pretrainResponse = %v\n", result)
		time.Sleep(5 * time.Second)
		getResponse, aErr := ac.GetPretrainState(sensorId)
		if aErr != nil {
			panic(err)
		}
		result = *getResponse.State
	}
	fmt.Printf("pretrainResponse = %v\n", result)
}
