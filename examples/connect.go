package main

import (
	amber_client "amber-go-sdk"
	"fmt"
)

func main() {

	amberClient, err := amber_client.NewAmberClient(nil, nil, nil, nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	response, err := amberClient.ListSensors()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", response)
}