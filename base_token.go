package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type BaseTokenService service

type BaseTokens struct {
	OST struct {
		Name                                 string `json:"name,omitempty"`
		Decimals                             int    `json:"decimals"`
		OriginChainErc20tokenContractAddress string `json:"origin_chain_erc20token_contract_address,omitempty"`
	} `json:"OST,omitempty"`
	USDC struct {
		Name                                 string `json:"name,omitempty"`
		Decimals                             int    `json:"decimals"`
		OriginChainErc20tokenContractAddress string `json:"origin_chain_erc20token_contract_address,omitempty"`
	} `json:"USDC,omitempty"`
}

type BaseTokenResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string      `json:"result_type,omitempty"`
		BaseTokens *BaseTokens `json:"base_tokens,omitempty"`
	} `json:"data,omitempty"`
}

func (s *BaseTokenService) Get() (*BaseTokenResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/base-tokens?%s", v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(BaseTokenResponse)
	params.ApiSignature = signature
	s.client.base.Get("base-tokens").QueryStruct(params).Receive(res, err)
	return res, err
}
