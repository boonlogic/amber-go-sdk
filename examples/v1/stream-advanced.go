package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall"
	"time"

	amberClient "github.com/boonlogic/amber-go-sdk/v1"
	amberModels "github.com/boonlogic/amber-go-sdk/v1/models"
)

// streamCsv:
// Streams sensor data from a current sensor.
func streamCsv(sensorId string, batch_size int, client *amberClient.AmberClient, records []string) {

	var data string
	streamPayload := amberModels.PostStreamRequest{}

	configResponse, err := client.GetConfig(sensorId)
	if err != nil {
		syscall.Exit(1)
	}

	fCount := int(*configResponse.FeatureCount)
	sampleCnt := int(*configResponse.SamplesToBuffer)
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
		sr, err := client.StreamSensor(sensorId, streamPayload)
		if err != nil {
			fmt.Printf("%v\n", err)
			syscall.Exit(1)
		}
		duration := time.Since(start)

		fmt.Printf("Duration: %vms, State: %v, %v%%\n", duration.Milliseconds(), *sr.StreamStatus.State, *sr.StreamStatus.Progress)
		fmt.Printf("    AD: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.AD)), ","), "[]"))
		fmt.Printf("    AH: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.AH)), ","), "[]"))
		fmt.Printf("    AM: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.AM)), ","), "[]"))
		fmt.Printf("    ID: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.ID)), ","), "[]"))
		fmt.Printf("    RI: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.RI)), ","), "[]"))
		fmt.Printf("    SI: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sr.SI)), ","), "[]"))

		state = *sr.State
	}
}

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
		fmt.Printf("configuring sensor")
		featureCount := uint16(1)
		streamingWindowSize := uint16(25)
		learningMaxSamples := uint64(3000)
		samplesToBuffer := uint32(1000)
		learningRateNumerator := uint64(1)
		learningRateDenominator := uint64(1000)
		configRequest := amberModels.PostConfigRequest{
			StreamingParameters: amberModels.StreamingParameters{
				AnomalyHistoryWindow:    nil,
				LearningMaxClusters:     nil,
				LearningMaxSamples:      &learningMaxSamples,
				LearningRateDenominator: &learningRateDenominator,
				LearningRateNumerator:   &learningRateNumerator,
			},
			FeatureCount:        &featureCount,
			Features:            nil,
			SamplesToBuffer:     &samplesToBuffer,
			StreamingWindowSize: &streamingWindowSize,
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
	records, _ := amberClient.LoadCsvRecords("examples/v1/output_current.csv")

	// stream a csv
	streamCsv(sensorId, 5, ac, records)
}
