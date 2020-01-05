package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type SessionsService service

type Session struct {
	UserID                    string `json:"user_id,omitempty"`
	Address                   string `json:"address,omitempty"`
	ExpirationHeight          int    `json:"expiration_height,omitempty"`
	ApproxExpirationTimestamp int    `json:"approx_expiration_timestamp,omitempty"`
	SpendingLimit             string `json:"spending_limit,omitempty"`
	Nonce                     int    `json:"nonce,omitempty"`
	Status                    string `json:"status,omitempty"`
	UpdatedTimestamp          int    `json:"updated_timestamp,omitempty"`
}

type SessionsResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string    `json:"result_type,omitempty"`
		Session    *Session  `json:"session,omitempty"`
		Sessions   []Session `json:"sessions,omitempty"`
	} `json:"data,omitempty"`
}

func (s *SessionsService) Get(userID, sessionAddress string) (*SessionsResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/sessions/%s?%s", userID, sessionAddress, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(SessionsResponse)
	params.ApiSignature = signature
	sessions := fmt.Sprintf("users/%s/sessions/%s", userID, sessionAddress)
	s.client.base.Get(sessions).QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *SessionsService) GetList(userID string) (*SessionsResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/sessions?%s", userID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(SessionsResponse)
	params.ApiSignature = signature
	sessions := fmt.Sprintf("users/%s/sessions", userID)
	s.client.base.Get(sessions).QueryStruct(params).Receive(res, err)
	return res, err
}
