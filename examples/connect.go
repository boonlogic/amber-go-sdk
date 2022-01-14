package main

import (
	amber_client "amber-go-sdk"
	"encoding/json"
	"fmt"
	"syscall"
)

func main() {

	amberClient, err := amber_client.NewAmberClient(nil, nil, nil, nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	response, err := amberClient.ListSensors()
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}
	formatted, _ := json.MarshalIndent(*response, "", "\t")
	fmt.Printf("%v\n", string(formatted))
}