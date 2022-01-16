package main

import (
	"encoding/json"
	"fmt"
	amberClient "github.com/boonlogic/amber-go-sdk"
	amberModels "github.com/boonlogic/amber-go-sdk/models"
	"syscall"
)

func main() {

	ac, err := amberClient.NewAmberClient(nil, nil, nil, nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// Get version
	fmt.Printf("get version\n")
	versionResponse, err := ac.GetVersion()
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formmatted, _ := json.MarshalIndent(*versionResponse, "", "\t")
	fmt.Printf("%v\n", string(formmatted))

	// List all sensors belonging to current user
	fmt.Printf("listing sensors\n")
	listResponse, err := ac.ListSensors()
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ := json.MarshalIndent(*listResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// Create a new sensor
	fmt.Printf("create sensor\n")
	createSensorResponse, err := ac.CreateSensor("new-go-sdk-sensor")
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*createSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// retain sensorId
	sensorId := *createSensorResponse.SensorID

	// get sensor info
	fmt.Printf("get sensor\n")
	getSensorResponse, err := ac.GetSensor(sensorId)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// update the label of a sensor
	fmt.Printf("update label\n")
	updateLabelResponse, err := ac.UpdateLabel(sensorId, "updated-go-sdk-sensor")
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*updateLabelResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// configure sensor
	fmt.Printf("configuring sensor\n")
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
	configSensorResponse, err := ac.ConfigureSensor(sensorId, postConfigRequest)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*configSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// get sensor configuration
	fmt.Printf("get sensor configuration\n")
	getConfigResponse, err := ac.GetConfig(sensorId)
	if err != nil {
		fmt.Printf("%v\n", err)
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
	streamSensorResponse, err := ac.StreamSensor(sensorId, streamPayload)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*streamSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("get cluster status\n")
	getStatusResponse, err := ac.GetStatus(sensorId)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getStatusResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("getting root cause\n")
	/*
		clusterId := "[1,2]"
		getRootCause, err := ac.GetRootCause(sensorId, &clusterId, nil)
		if err != nil {
			fmt.Printf("%v\n", err)
			syscall.Exit(1)
		}
		formatted, _ = json.MarshalIndent(*getRootCause, "", "\t")
		fmt.Printf("%v\n", string(formatted))
	*/

	fmt.Printf("delete sensor instance\n")
	err = ac.DeleteSensor(sensorId)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	fmt.Printf("%v deleted\n", sensorId)
}
