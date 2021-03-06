package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type DevicesService service

type Device struct {
	UserID           string `json:"user_id,omitempty"`
	Address          string `json:"address,omitempty"`
	LinkedAddress    string `json:"linked_address,omitempty"`
	ApiSignerAddress string `json:"api_signer_address,omitempty"`
	Status           string `json:"status,omitempty"`
	UpdatedTimestamp int    `json:"updated_timestamp,omitempty"`
}

type NewDevice struct {
	QueryParams
	Address          string `json:"address,omitempty"`
	ApiSignerAddress string `json:"api_signer_address,omitempty"`
}

type DeviceResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType string   `json:"result_type,omitempty"`
		Device     *Device  `json:"device,omitempty"`
		Devices    []Device `json:"devices,omitempty"`
	} `json:"data,omitempty"`
}

type DeviceParams struct {
	QueryParams
	Addresses []string `url:"addresses,omitempty"`
	Limit     int      `url:"limit,omitempty"`
}

func (s *DevicesService) Get(userID, address string) (*DeviceResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/devices/%s?%s", userID, address, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(DeviceResponse)
	params.ApiSignature = signature
	device := fmt.Sprintf("users/%s/devices/%s", userID, address)
	s.client.base.Get(device).QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *DevicesService) GetList(userID string, params DeviceParams) (*DeviceResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	q := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	params.QueryParams = q
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/devices?%s", userID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(DeviceResponse)
	params.ApiSignature = signature
	devices := fmt.Sprintf("users/%s/devices", userID)
	s.client.base.Get(devices).QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *DevicesService) Create(userID string, params NewDevice) (*DeviceResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	q := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	params.QueryParams = q
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/devices?%s", userID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(DeviceResponse)
	params.ApiSignature = signature
	devices := fmt.Sprintf("users/%s/devices", userID)
	s.client.base.Post(devices).BodyForm(params).Receive(res, err)
	return res, err
}
