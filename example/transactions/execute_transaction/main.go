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
	r, _ := client.Transactions.Execute("b6504ca4-9263-4998-8036-c90e648c48de", ost.TransactionParams{
		To:          "0xced071365a63ba20ca8b50391c62c23aa1831422",
		RawCallData: `{"method":"directTransfers","parameters":[["0xced071365a63ba20ca8b50391c62c23aa1831422"],["1000000000000000000"]]}`,
		MetaProperty: &ost.MetaProperty{
			Name:    "Name for this transaction",
			Type:    "company_to_user",
			Details: "Details about the transaction",
		},
	})
	response, _ := json.Marshal(r)
	fmt.Println(string(response))
}
