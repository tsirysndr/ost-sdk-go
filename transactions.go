package ost

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type TransactionsService service

type Transaction struct {
	ID                string        `json:"id,omitempty"`
	TransactionHash   string        `json:"transaction_hash,omitempty"`
	From              string        `json:"from,omitempty"`
	To                string        `json:"to,omitempty"`
	Nonce             int           `json:"nonce,omitempty"`
	Value             string        `json:"value,omitempty"`
	GasPrice          string        `json:"gas_price,omitempty"`
	GasUsed           int           `json:"gas_used,omitempty"`
	TransactionFee    string        `json:"transaction_fee,omitempty"`
	BlockConfirmation int           `json:"block_confirmation,omitempty"`
	Status            string        `json:"status,omitempty"`
	UpdatedTimestamp  int           `json:"updated_timestamp,omitempty"`
	BlockNumber       int           `json:"block_number,omitempty"`
	RuleName          string        `json:"rule_name,omitempty"`
	Transfers         []Transfer    `json:"transfers,omitempty"`
	MetaProperty      *MetaProperty `json:"meta_property,omitempty"`
}

type TransactionResponse struct {
	Success bool `json:"success,omitempty"`
	Data    *struct {
		ResultType   string        `json:"result_type,omitempty"`
		Transaction  *Transaction  `json:"transaction,omitempty"`
		Transactions []Transaction `json:"transactions,omitempty"`
	} `json:"data,omitempty"`
}

type TransactionParams struct {
	QueryParams
	To           string        `url:"to,omitempty"`
	RawCallData  string        `url:"raw_call_data,omitempty"`
	MetaProperty *MetaProperty `url:"meta_property,omitempty"`
}

type MetaProperty struct {
	Name    string `url:"name,omitempty"`
	Type    string `url:"type,omitempty"`
	Details string `url:"details,omitempty"`
}

type Transfer struct {
	From       string `json:"from,omitempty"`
	FromUserID string `json:"from_user_id,omitempty"`
	To         string `json:"to,omitempty"`
	ToUserID   string `json:"to_user_id,omitempty"`
	Amount     string `json:"amount,omitempty"`
	Kind       string `json:"kind,omitempty"`
}

func (s *TransactionsService) Get(userID, transactionID string) (*TransactionResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := &QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/transactions/%s?%s", userID, transactionID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	params.ApiSignature = signature
	res := new(TransactionResponse)
	transaction := fmt.Sprintf("users/%s/transactions/%s", userID, transactionID)
	s.client.base.Get(transaction).QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *TransactionsService) GetList(userID string) (*TransactionResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	params := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/transactions?%s", userID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	res := new(TransactionResponse)
	params.ApiSignature = signature
	transactions := fmt.Sprintf("users/%s/transactions", userID)
	s.client.base.Get(transactions).QueryStruct(params).Receive(res, err)
	return res, err
}

func (s *TransactionsService) Execute(userID string, params TransactionParams) (*TransactionResponse, error) {
	var err error
	timestamp := time.Now().Unix()
	q := QueryParams{
		ApiKey:              s.client.options.ApiKey,
		ApiRequestTimestamp: timestamp,
		ApiSignatureKind:    SIGNATURE_KIND,
	}
	params.QueryParams = q
	v, _ := query.Values(params)
	resource := fmt.Sprintf("/users/%s/transactions?%s", userID, v.Encode())
	signature := SignQueryParams(resource, s.client.options.ApiSecret)
	params.ApiSignature = signature
	res := new(TransactionResponse)
	endpoint := fmt.Sprintf("/users/%s/transactions", userID)
	s.client.base.Post(endpoint).BodyForm(params).Receive(res, err)
	return res, err
}
