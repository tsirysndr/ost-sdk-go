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
	r, _ := client.RecoveryOwner.Get("b6504ca4-9263-4998-8036-c90e648c48de", "0xe37906219ad67cc1301b970539c9860f9ce8d991")
	response, _ := json.Marshal(r)
	fmt.Println(string(response))
}
