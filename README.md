![Logo](https://github.com/boonlogic/amber-go-sdk/blob/master/docs/BoonLogic.png?raw=true)

# Boon Amber GO SDK

A Go SDK for Boon Amber sensor analytics

- __Website__: [boonlogic.com](https://boonlogic.com)
- __Documentation__: [Boon Docs Main Page](https://docs.boonlogic.com)
- __SDK Functional Breakdown__: [amber-go-sdk functional breakdown](https://boonlogic.github.io/amber-go-sdk/docs/user-docs.html)

## Installation

```
go get github.com/boonlogic/amber-go-sdk
```

## Credentials setup

Note: An account in the Boon Amber cloud must be obtained from Boon Logic to use the Amber SDK.

The username and password should be placed in a file named _~/.Amber.license_ whose contents are the following:

```json
{
    "default": {
        "username": "AMBER-ACCOUNT-USERNAME",
        "password": "AMBER-ACCOUNT-PASSWORD",
        "server": "https://amber.boonlogic.com/v1"
    }
}
```

The _~/.Amber.license_ file will be consulted by the Amber SDK to find and authenticate your account credentials with the Amber server. Credentials may optionally be provided instead via the environment variables `AMBER_USERNAME` and `AMBER_PASSWORD`.

## Connectivity test

The following is a simple connectivity testing using the GetVersion function.

[connect.go](https://github.com/boonlogic/amber-go-sdk/blob/master/examples/connect.go)

```go
package main

import (
	"encoding/json"
	"fmt"
	amberClient "github.com/boonlogic/amber-go-sdk/v1"
	"syscall"
)

func main() {

	ac, err := amberClient.NewAmberClientFromFile(nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	response, err := ac.GetVersion()
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ := json.MarshalIndent(*response, "", "\t")
	fmt.Printf("%v\n", string(formatted))
}
```

Running the connect-example.go script should yield output like the following:

`$ go run examples/connect.go`
```json
{
    "release": "0.0.405",
    "api-version": "/v1",
    "builder": "ec74f421",
    "expert-api": "dee23681",
    "expert-common": "300a588e",
    "nano-secure": "61c431e2",
    "swagger-ui": "914af396"
}
```

## Full Example

The following Go routine will demonstrate each API call in the Amber go SDK.

[full-example.go](https://github.com/boonlogic/amber-go-sdk/blob/master/examples/full-example.go)

```go
package main

import (
	"encoding/json"
	"fmt"
	amberClient "github.com/boonlogic/amber-go-sdk/v1"
	amberModels "github.com/boonlogic/amber-go-sdk/v1/models"
	"syscall"
)

func main() {

	ac, err := amberClient.NewAmberClientFromfile(nil, nil)
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
```


## Advanced CSV file processor

The following will process a CSV file using batch-style streaming requests.  Full Amber analytic results will be displayed after each streaming request.  

[stream-advanced.go](https://github.com/boonlogic/amber-go-sdk/blob/master/examples/stream-advanced.go) <br>
[output_current.csv](https://github.com/boonlogic/amber-go-sdk/blob/master/examples/output_current.csv)
