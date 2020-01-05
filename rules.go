package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type RulesService service

type Rule struct {
	ID      int         `json:"id,omitempty"`
	TokenID int         `json:"token_id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Address string      `json:"address,omitempty"`
	ABI     interface{} `json:"abi,omitempty"`
}

type RuleResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string `json:"result_type,omitempty"`
		Rules      []Rule `json:"rules,omitempty"`
	} `json:"data,omitempty"`
}

func (s *RulesService) GetList() (*RuleResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/rules?%s", v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(RuleResponse)
	params.ApiSignature = signature
	s.client.base.Get("rules").QueryStruct(params).Receive(res, err)
	return res, err
}
