package main

import (
	"encoding/json"
	"fmt"
	amberClient "github.com/boonlogic/amber-go-sdk"
	amberModels "github.com/boonlogic/amber-go-sdk/models"
	"syscall"
)

func main() {

	ac, err := amberClient.NewAmberClientFromFile(nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}

	// Get version
	fmt.Printf("get version\n")
	versionResponse, aErr := ac.GetVersion()
	if aErr != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formmatted, _ := json.MarshalIndent(*versionResponse, "", "\t")
	fmt.Printf("%v\n", string(formmatted))

	// List all sensors belonging to current user
	fmt.Printf("listing sensors\n")
	listResponse, aErr := ac.ListSensors()
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ := json.MarshalIndent(*listResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// Create a new sensor
	fmt.Printf("create sensor\n")
	createSensorResponse, aErr := ac.CreateSensor("new-go-sdk-sensor")
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*createSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// retain sensorId
	sensorId := *createSensorResponse.SensorID

	// get sensor info
	fmt.Printf("get sensor\n")
	getSensorResponse, aErr := ac.GetSensor(sensorId)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// update the label of a sensor
	fmt.Printf("update label\n")
	updateLabelResponse, aErr := ac.UpdateLabel(sensorId, "updated-go-sdk-sensor")
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*updateLabelResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// configure sensor
	fmt.Printf("configuring sensor\n")
	var featureCount uint16 = 1
	var streamingWindowSize uint16 = 25
	postConfigRequest := amberModels.PostConfigRequest{
		StreamingParameters: amberModels.StreamingParameters{
			AnomalyHistoryWindow:    nil,
			LearningMaxClusters:     nil,
			LearningMaxSamples:      nil,
			LearningRateDenominator: nil,
			LearningRateNumerator:   nil,
		},
		FeatureCount:        &featureCount,
		Features:            nil,
		SamplesToBuffer:     nil,
		StreamingWindowSize: &streamingWindowSize,
	}
	configSensorResponse, aErr := ac.ConfigureSensor(sensorId, postConfigRequest)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*configSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// get sensor configuration
	fmt.Printf("get sensor configuration\n")
	getConfigResponse, aErr := ac.GetConfig(sensorId)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getConfigResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// amber stream
	fmt.Printf("stream data\n")
	data := "0.1,0.1,0.3"
	saveImage := true
	streamPayload := amberModels.PostStreamRequest{
		Data:      &data,
		SaveImage: &saveImage,
	}
	streamSensorResponse, aErr := ac.StreamSensor(sensorId, streamPayload)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*streamSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("get cluster status\n")
	getStatusResponse, aErr := ac.GetStatus(sensorId)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getStatusResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("getting root cause\n")
	/*
		clusterId := "[1,2]"
		getRootCause, err := ac.GetRootCause(sensorId, &clusterId, nil)
		if err != nil {
			fmt.Printf("%v\n", aErr)
			syscall.Exit(1)
		}
		formatted, _ = json.MarshalIndent(*getRootCause, "", "\t")
		fmt.Printf("%v\n", string(formatted))
	*/

	fmt.Printf("post outage\n")
	postOutageResponse, aErr := ac.PostOutage(sensorId)
	if aErr != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*postOutageResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("delete sensor instance\n")
	aErr = ac.DeleteSensor(sensorId)
	if aErr != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	fmt.Printf("%v deleted\n", sensorId)
}
