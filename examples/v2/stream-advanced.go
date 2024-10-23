package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall"
	"time"

	amberClient "github.com/boonlogic/amber-go-sdk/v2"
	amberModels "github.com/boonlogic/amber-go-sdk/v2/models"
)

// streamCsv:
// Streams sensor data from a current sensor.
func streamCsv(modelId string, batch_size int, client *amberClient.AmberClient, records []string) {

	var data string
	streamPayload := amberModels.PostDataRequest{}

	configResponse, err := client.GetModelConfig(modelId)
	if err != nil {
		syscall.Exit(1)
	}

	fCount := len(configResponse.Features)
	sampleCnt := int(*configResponse.Training.BufferingSamples)
	state := ""

	elements := 0
	for idx := 0; idx < len(records); idx += elements {

		switch state {
		case "":
			// send entire sample buffer
			elements = sampleCnt
		case "Autotuning":
			// send a single vector
			elements = fCount
		case "Learning":
			// send in blocks of 200 vectors (to speed things up)
			elements = fCount * 200
		case "Monitoring":
			// send in single vectors
			elements = fCount
		}

		if idx+elements > len(records) {
			// handle final streaming call of data set
			elements = len(records) - idx
		}

		// construct csv data string
		data = strings.Join(records[idx:idx+elements], ",")
		streamPayload.Data = &data

		start := time.Now()
		sr, err := client.PostModelData(modelId, streamPayload)
		if err != nil {
			fmt.Printf("%v\n", err)
			syscall.Exit(1)
		}
		duration := time.Since(start)

		fmt.Printf("Duration: %vms, State: %v, %v%%\n", duration.Milliseconds(), sr.Status.State, *sr.Status.Progress)
		fmt.Printf("    AD: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.Analytics.AD)), ","), "[]"))
		fmt.Printf("    AH: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.Analytics.AH)), ","), "[]"))
		fmt.Printf("    ID: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.Analytics.ID)), ","), "[]"))
		fmt.Printf("    RI: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.Analytics.RI)), ","), "[]"))
		fmt.Printf("    SI: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.Analytics.SI)), ","), "[]"))

		state = string(sr.Status.State)
	}
}

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
		createSensorResponse, aErr := ac.PostModel(modelLabel)
		if aErr != nil {
			panic(aErr)
		}
		modelId = createSensorResponse.ID
		fmt.Printf("created sensor id: %v\n", modelId)
	} else {
		fmt.Printf("using sensor id: %v\n", modelId)
	}

	if configureSensor {
		// configure the sensor
		fmt.Printf("configuring sensor")
		fc := uint16(1)
		fusionRule := "submit"
		fusionTTL := uint64(100)
		var sw amberModels.StreamingWindow = 25
		features := make([]*amberModels.FeatureConfig, fc)
		features[0] = &amberModels.FeatureConfig{
			FusionRule: &fusionRule,
			FusionTTL:  &fusionTTL,
			MaxVal:     1000.0,
			MinVal:     0.0,
			Name:       "Feature_0",
		}
		autotune_pv, autotune_range := true, true
		maxclusters := uint16(2000)
		bufferingsamples := uint32(1000)
		learningsamples := uint64(3000)
		denominator := uint64(1000)
		numerator := uint64(1)

		configRequest := amberModels.PostConfigRequest{
			Autotuning: &amberModels.AutotuneConfig{
				PercentVariation: &autotune_pv,
				Range:            &autotune_range,
			},
			Features: features,
			Training: &amberModels.TrainingConfig{
				BufferingSamples:        &bufferingsamples,
				HistoryWindow:           nil,
				LearningMaxClusters:     &maxclusters,
				LearningMaxSamples:      &learningsamples,
				LearningRateDenominator: &denominator,
				LearningRateNumerator:   &numerator,
			},
			StreamingWindow: &sw,
		}
		configSensorResponse, aErr := ac.PostModelConfig(modelId, configRequest)
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
	records, _ := amberClient.LoadCsvRecords("examples/v2/output_current.csv")

	// stream a csv
	streamCsv(modelId, 5, ac, records)
}
