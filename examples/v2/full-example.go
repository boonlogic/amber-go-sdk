package main

import (
	"encoding/json"
	"fmt"
	"syscall"

	amberClient "github.com/boonlogic/amber-go-sdk/v2"
	amberModels "github.com/boonlogic/amber-go-sdk/v2/models"
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
	listResponse, aErr := ac.GetModels()
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ := json.MarshalIndent(*listResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// Create a new sensor
	fmt.Printf("create sensor\n")
	createSensorResponse, aErr := ac.PostModel("new-go-sdk-sensor")
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*createSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// retain modelId
	modelId := createSensorResponse.ID

	// get sensor info
	fmt.Printf("get sensor\n")
	getSensorResponse, aErr := ac.GetModel(modelId)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// update the label of a sensor
	fmt.Printf("update label\n")
	updateLabelResponse, aErr := ac.PutModel(modelId, "updated-go-sdk-sensor")
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*updateLabelResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// configure sensor
	fmt.Printf("configuring sensor\n")
	fusionRule := "submit"
	fusionTTL := uint64(100)
	features := make([]*amberModels.FeatureConfig, 1)
	features[0] = &amberModels.FeatureConfig{
		MaxVal:     1000.0,
		MinVal:     0.0,
		Name:       "feature_0",
		Weight:     1,
		FusionRule: &fusionRule,
		FusionTTL:  &fusionTTL,
	}
	var streamingWindowSize amberModels.StreamingWindow = 25
	var pv amberModels.PercentVariation = 0.05
	postConfigRequest := amberModels.PostConfigRequest{
		Autotuning:       nil,
		Features:         features,
		StreamingWindow:  &streamingWindowSize,
		PercentVariation: &pv,
		Training:         nil,
	}
	configSensorResponse, aErr := ac.PostModelConfig(modelId, postConfigRequest)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*configSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	// get sensor configuration
	fmt.Printf("get sensor configuration\n")
	getConfigResponse, aErr := ac.GetModelConfig(modelId)
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
	streamPayload := amberModels.PostDataRequest{
		Data:      &data,
		SaveImage: &saveImage,
	}
	streamSensorResponse, aErr := ac.PostModelData(modelId, streamPayload)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*streamSensorResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("get cluster status\n")
	getStatusResponse, aErr := ac.GetModelStatus(modelId)
	if err != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ = json.MarshalIndent(*getStatusResponse, "", "\t")
	fmt.Printf("%v\n", string(formatted))

	fmt.Printf("post outage\n")
	aErr = ac.PostModelOutage(modelId)
	if aErr != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}

	fmt.Printf("delete sensor instance\n")
	aErr = ac.DeleteModel(modelId)
	if aErr != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	fmt.Printf("%v deleted\n", modelId)
}
