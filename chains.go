package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type ChainsService service

type Chain struct {
	ID               int    `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	BlockHeight      int    `json:"block_height,omitempty"`
	BlockTime        int    `json:"block_time,omitempty"`
	UpdatedTimestamp int    `json:"updted_timestamp,omitempty"`
}

type ChainResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string `json:"result_type,omitempty"`
		Chain      *Chain `json:"chain,omitempty"`
	} `json:"data,omitempty"`
}

func (s *ChainsService) Get(ID int) (*ChainResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/chains/%d?%s", ID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(ChainResponse)
	params.ApiSignature = signature
	price_points := fmt.Sprintf("chains/%d", ID)
	s.client.base.Get(price_points).QueryStruct(params).Receive(res, err)
	return res, err
}
