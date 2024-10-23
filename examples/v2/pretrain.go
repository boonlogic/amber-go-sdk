package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	amberClient "github.com/boonlogic/amber-go-sdk/v2"
	amberModels "github.com/boonlogic/amber-go-sdk/v2/models"
)

func main() {

	// create the client
	ac, err := amberClient.NewAmberClientFromFile(nil, nil)
	if err != nil {
		panic(err)
	}

	// Set configureSensor to false to skip reconfiguring sensor
	configureSensor := true

	// Fill in the modelId here if one is already known.
	modelId := ""
	modelLabel := "new-go-sdk-sensor"

	existingSensorResponse, aErr := ac.GetModels()
	if aErr != nil {
		panic(aErr)
	}
	for _, element := range existingSensorResponse.ModelList {
		if element.Label == modelLabel {
			modelId = element.ID
			break
		}
	}

	if modelId == "" {
		// create a sensor when none specified
		fmt.Printf("PostModel: %v\n", modelLabel)
		createSensorResponse, aErr := ac.PostModel(modelLabel)
		if aErr != nil {
			panic(aErr)
		}
		modelId = createSensorResponse.ID
		fmt.Printf("created model id: %v\n", modelId)
	} else {
		fmt.Printf("using model id: %v\n", modelId)
	}

	fusionRule := "submit"
	fusionTTL := uint64(100)
	if configureSensor {
		fmt.Println("configureSensor")
		// configure the sensor
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
		autotune_pv, autotune_range := true, true
		maxclusters := uint16(1000)
		bufferingsamples := uint32(10000)
		learningsamples := uint64(20000)
		window := uint32(1000)
		denominator := uint64(10000)
		numerator := uint64(1)
		// set just the basics as pretraining sets the rest
		configRequest := amberModels.PostConfigRequest{
			Autotuning: &amberModels.AutotuneConfig{
				PercentVariation: &autotune_pv,
				Range:            &autotune_range,
			},
			Features:        features,
			StreamingWindow: &streamingWindowSize,
			Training: &amberModels.TrainingConfig{
				BufferingSamples:        &bufferingsamples,
				HistoryWindow:           &window,
				LearningMaxClusters:     &maxclusters,
				LearningMaxSamples:      &learningsamples,
				LearningRateDenominator: &denominator,
				LearningRateNumerator:   &numerator,
			},
		}

		configSensorResponse, aErr := ac.PostModelConfig(modelId, configRequest)
		if aErr != nil {
			panic(aErr)
		}
		formatted, err := json.MarshalIndent(*configSensorResponse, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Config Response: %v\n", string(formatted))
	}

	// read in entire csv file
	records, _ := amberClient.LoadCsvRecords("examples/v2/output_current.csv")

	// pretrain the requestc
	pretrainData := strings.Join(records, ",")
	pretrainRequest := amberModels.PostPretrainRequest{
		Data: &pretrainData,
	}

	pretrainResponse, aErr := ac.PretrainModel(modelId, pretrainRequest)
	if aErr != nil {
		panic(aErr)
	}
	result := pretrainResponse.Status
	for result == "Pretraining" {
		fmt.Printf("State: %v \n", result)
		time.Sleep(5 * time.Second)
		getResponse, aErr := ac.GetPretrainState(modelId)
		if aErr != nil {
			panic(err)
		}
		result = getResponse.Status
	}
	fmt.Printf("State: %v \n", result)

	aErr = ac.DeleteModel(modelId)
	if aErr != nil {
		panic(aErr)
	}
}
