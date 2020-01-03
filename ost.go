package ost

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/dghubble/sling"
)

const SIGNATURE_KIND = "OST1-HMAC-SHA256"

type Client struct {
	base          *sling.Sling
	options       Options
	common        service
	Balance       *BalanceService
	BaseToken     *BaseTokenService
	Chains        *ChainsService
	DeviceManager *DeviceManagerService
	Devices       *DevicesService
	PricePoint    *PricePointService
	RecoveryOwner *RecoveryOwnerService
	Rules         *RulesService
	Sessions      *SessionsService
	Token         *TokenService
	Transactions  *TransactionsService
	Users         *UsersService
	Webhooks      *WebhooksService
}

type service struct {
	client *Client
}

type Options struct {
	Endpoint  string
	ApiKey    string
	ApiSecret string
}

type QueryParams struct {
	ApiKey              string `url:"api_key"`
	ApiRequestTimestamp int64  `url:"api_request_timestamp"`
	ApiSignature        string `url:"api_signature"`
	ApiSignatureKind    string `url:"api_signature_kind"`
}

func NewClient(o Options) *Client {
	base := sling.New().Base(o.Endpoint)
	c := &Client{}
	c.base = base
	c.options = o
	c.common.client = c
	c.Balance = (*BalanceService)(&c.common)
	c.BaseToken = (*BaseTokenService)(&c.common)
	c.Chains = (*ChainsService)(&c.common)
	c.DeviceManager = (*DeviceManagerService)(&c.common)
	c.Devices = (*DevicesService)(&c.common)
	c.PricePoint = (*PricePointService)(&c.common)
	c.RecoveryOwner = (*RecoveryOwnerService)(&c.common)
	c.Rules = (*RulesService)(&c.common)
	c.Sessions = (*SessionsService)(&c.common)
	c.Token = (*TokenService)(&c.common)
	c.Transactions = (*TransactionsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)
	return c
}

func SignQueryParams(resource, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(resource))
	return hex.EncodeToString(h.Sum(nil))
}
