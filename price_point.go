package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type PricePointService service

type PricePoint struct {
	OST *struct {
		USD       float64 `json:"USD,omitempty"`
		EUR       float64 `json:"EUR,omitempty"`
		GBP       float64 `json:"GBP,omitempty"`
		Decimals  int     `json:"decimals,omitempty"`
		UpdatedAt int     `json:"updated_at,omitempty"`
	} `json:"OST,omitempty"`
	USDC *struct {
		USD       float64 `json:"USD,omitempty"`
		EUR       float64 `json:"EUR,omitempty"`
		GBP       float64 `json:"GBP,omitempty"`
		Decimals  int     `json:"decimals,omitempty"`
		UpdatedAt int     `json:"updated_at,omitempty"`
	} `json:"USDC,omitempty"`
}

type PricePointResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string      `json:"result_type,omitempty"`
		PricePoint *PricePoint `json:"price_point,omitempty"`
	} `json:"data,omitempty"`
}

func (s *PricePointService) Get(chainID int) (*PricePointResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/chains/%d/price-points?%s", chainID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(PricePointResponse)
	params.ApiSignature = signature
	price_points := fmt.Sprintf("chains/%d/price-points", chainID)
	s.client.base.Get(price_points).QueryStruct(params).Receive(res, err)
	return res, err
}
