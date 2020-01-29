package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type RecoveryOwnerService service

type RecoveryOwner struct {
	UserID           string `json:"user_id,omitempty"`
	Address          string `json:"address,omitempty"`
	Status           string `json:"status,omitempty"`
	UpdatedTimestamp int    `json:"updated_timestamp,omitempty"`
}

type RecoveryOwnerResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType    string        `json:"result_type,omitempty"`
		RecoveryOwner RecoveryOwner `json:"rules,omitempty"`
	} `json:"data,omitempty"`
}

func (s *RecoveryOwnerService) Get(userID, recoveryOwnerAddress string) (*RecoveryOwnerResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/recovery-owners/%s?%s", userID, recoveryOwnerAddress, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(RecoveryOwnerResponse)
	params.ApiSignature = signature
	recovery_owners := fmt.Sprintf("users/%s/recovery-owners/%s", userID, recoveryOwnerAddress)
	s.client.base.Get(recovery_owners).QueryStruct(params).Receive(res, err)
	return res, err
}
