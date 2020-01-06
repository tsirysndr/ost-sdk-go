package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type TokenService service

type Token struct {
	ID               int              `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Symbol           string           `json:"symbol,omitempty"`
	BaseToken        string           `json:"base_token,omitempty"`
	ConversionFactor int              `json:"conversion_factor,omitempty"`
	TotalSupply      string           `json:"total_supply,omitempty"`
	Decimals         int              `json:"decimals,omitempty"`
	OriginChain      *OriginChain     `json:"origin_chain,omitempty"`
	Stakers          []string         `json:"stakers,omitempty"`
	AuxiliaryChain   []AuxiliaryChain `json:"auxiliary_chains,omitempty"`
	UpdatedTimestamp int              `json:"updated_timestamp,omitempty"`
}

type OriginChain struct {
	ChainID      int    `json:"chain_id,omitempty"`
	BrandedToken string `json:"branded_token,omitempty"`
	Organization *struct {
		Contract string `json:"contract,omitempty"`
		Owner    string `json:"owner,omitempty"`
	} `json:"organization,omitempty"`
}

type AuxiliaryChain struct {
	ChainID             int      `json:"chain_id,omitempty"`
	UtilityBrandedToken string   `json:"utility_branded_token,omitempty"`
	CompanyTokenHolders []string `json:"company_token_holders,omitempty"`
	CompanyUUIDs        []string `json:"company_uuids,omitempty"`
	Organization        *struct {
		Contract string `json:"contract,omitempty"`
		Owner    string `json:"owner,omitempty"`
	} `json:"organization,omitempty"`
}

type TokenResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string `json:"result_type,omitempty"`
		Token      *Token `json:"token,omitempty"`
	} `json:"data,omitempty"`
}

func (s *TokenService) Get() (*TokenResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/tokens?%s", v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(TokenResponse)
	params.ApiSignature = signature
	tokens := fmt.Sprintf("tokens")
	s.client.base.Get(tokens).QueryStruct(params).Receive(res, err)
	return res, err
}
