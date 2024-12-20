package main

import (
	"encoding/json"
	"fmt"
	"syscall"

	amberClient "github.com/boonlogic/amber-go-sdk/v1"
)

func main() {

	ac, err := amberClient.NewAmberClientFromFile(nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		syscall.Exit(1)
	}

	response, aErr := ac.GetVersion()
	if aErr != nil {
		fmt.Printf("%v\n", aErr)
		syscall.Exit(1)
	}
	formatted, _ := json.MarshalIndent(*response, "", "\t")
	fmt.Printf("%v\n", string(formatted))
}
