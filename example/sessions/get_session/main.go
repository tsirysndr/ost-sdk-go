package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/ost-sdk-go"
)

func main() {
	options := ost.Options{
		Endpoint:  "<API_ENDPOINT>",
		ApiKey:    "<YOUR_API_KEY>",
		ApiSecret: "<YOUR_API_SECRET>",
	}
	client := ost.NewClient(options)
	r, _ := client.Sessions.Get("b6504ca4-9263-4998-8036-c90e648c48de", "0x48c8fd8bb4376e3abb62a6ffaf7c26db58b310fd")
	response, _ := json.Marshal(r)
	fmt.Println(string(response))
}
