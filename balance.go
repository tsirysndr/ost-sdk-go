package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type BalanceService service

type Balance struct {
	UserID           string `json:"user_id,omitempty"`
	TotalBalance     string `json:"total_balance,omitempty"`
	AvailableBalance string `json:"available_balance,omitempty"`
	UnsettledDebit   string `json:"unsettled_debit,omitempty"`
	UpdatedTimestamp int    `json:"updated_timestamp,omitempty"`
}

type BalanceResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string   `json:"result_type,omitempty"`
		Balance    *Balance `json:"balance,omitempty"`
	} `json:"data,omitempty"`
}

func (s *BalanceService) Get(userID string) (*BalanceResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/balance?%s", userID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(BalanceResponse)
	params.ApiSignature = signature
	balance := fmt.Sprintf("users/%s/balance", userID)
	s.client.base.Get(balance).QueryStruct(params).Receive(res, err)
	return res, err
}
