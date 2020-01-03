package ost

import (
	"fmt"
	"time"
)

type UsersService service

type User struct {
	ID                   string `json:"id,omitempty"`
	TokenID              int    `json:"token_id,omitempty"`
	TokenHolderAddress   string `json:"token_holder_address,omitempty"`
	DeviceManagerAddress string `json:"device_manager_address,omitempty"`
	RecoveryAddress      string `json:"recovery_address,omitempty"`
	RecoveryOwnerAddress string `json:"recovery_owner_address,omitempty"`
	Type                 string `json:"type,omitempty"`
	Status               string `json:"status,omitempty"`
	UpdatedTimestamp     int    `json:"updated_timestamp,omitempty"`
}

type UserResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string `json:"result_type,omitempty"`
		User       *User  `json:"user,omitempty"`
		Users      []User `json:"users,omitempty"`
	} `json:"data,omitempty"`
}

type Empty struct{}

func (s *UsersService) Create() (*UserResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	resource := fmt.Sprintf("/users?api_key=%s&api_request_timestamp=%d&api_signature_kind=OST1-HMAC-SHA256", s.client.options.ApiKey, timestamp)
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	params := &QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
		ApiSignature:        signature,
	}
	res := new(UserResponse)
	s.client.base.Post("users").BodyForm(params).Receive(res, err)
	return res, err
}

func (s *UsersService) Get(ID string) (*UserResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	resource := fmt.Sprintf("/users/%s?api_key=%s&api_request_timestamp=%d&api_signature_kind=OST1-HMAC-SHA256", ID, s.client.options.ApiKey, timestamp)
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	params := &QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
		ApiSignature:        signature,
	}
	res := new(UserResponse)
	s.client.base.Path("users/").Get(ID).QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *UsersService) GetList() (*UserResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	resource := fmt.Sprintf("/users?api_key=%s&api_request_timestamp=%d&api_signature_kind=OST1-HMAC-SHA256", s.client.options.ApiKey, timestamp)
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(UserResponse)
	params := &QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
		ApiSignature:        signature,
	}
	s.client.base.Get("users").QueryStruct(params).Receive(res, err)
	return res, err
}
