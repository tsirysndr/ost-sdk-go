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
	r, _ := client.Transactions.Get("b6504ca4-9263-4998-8036-c90e648c48de", "3b905a4f-db65-42f1-98b3-3ba256a4f704")
	response, _ := json.Marshal(r)
	fmt.Println(string(response))
}
