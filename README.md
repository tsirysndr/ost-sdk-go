# OST Go SDK

<p>
  <a href="https://github.com/tsirysndr/ost-sdk-go/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/ost-sdk-go.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/ost-sdk-go">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/ost-sdk-go">
  <a href="https://github.com/tsirysndr/ost-sdk-go/issues?q=is%3Apr+is%3Aclosed">
    <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/ost-sdk-go">
  </a>
  <a href="https://github.com/tsirysndr/ost-sdk-go/pulls">
    <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/ost-sdk-go">
  </a>
  <a href="https://github.com/tsirysndr/ost-sdk-go/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/ost-sdk-go">
  </a>
  <a href="https://github.com/tsirysndr/ost-sdk-go/graphs/contributors">
    <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/ost-sdk-go">
  </a>
  <a href="https://github.com/tsirysndr/ost-sdk-go/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-MIT-green.svg" target="_blank" />
  </a>
</p>

OST Platform SDK for Go

## Introduction

OST is a complete technology solution enabling mainstream businesses 
to easily launch blockchain-based economies without 
requiring blockchain development.

At the core of OST is the concept of OST-powered Brand Tokens (BTs). 
BTs are white-label cryptocurrency tokens with utility representations 
running on highly-scalable Ethereum-based side blockchains, 
backed by OST tokens staked on Ethereum mainnet. Within a business’s 
token economy, BTs can only be transferred to whitelisted user addresses. 
This ensures that they stay within the token economy.

The OST technology stack is designed to give businesses everything they need 
to integrate, test, and deploy BTs. Within the OST suite of products, developers 
can use OST Platform to create, test, and launch Brand Tokens backed by OST. 

OST APIs and server-side SDKs make it simple and easy for developers to 
integrate blockchain tokens into their apps.

## Requirements

Integrating an OST SDK into your application can begin as soon as you create an account 
with OST Platform, requiring only three steps:
1. Sign-up on [https://platform.ost.com](https://platform.ost.com).
2. Create your Brand Token in OST Platform.
3. Obtain an API Key and API Secret from [https://platform.ost.com/mainnet/developer](https://platform.ost.com/mainnet/developer).

## Documentation

[https://dev.ost.com/](https://dev.ost.com/)

## Installation

```bash
> go get github.com/tsirysndr/ost-sdk-go
```

## Getting Started

Import the package into your project.

```Go
import "github.com/tsirysndr/ost-sdk-go"
```

Construct a new OST client, then use the various services on the client to access different parts of the OST API. For example:

```Go
// the latest valid API endpoint is "https://api.ost.com/mainnet/v2/"
config := ost.Config{
  Endpoint:  "<API_ENDPOINT>",
  ApiKey:    "<YOUR_API_KEY>",
  ApiSecret: "<YOUR_API_SECRET>",
}
client := ost.NewClient(config)
```

## SDK Modules

If a user's private key is lost, they could lose access 
to their tokens. To tackle this risk, OST promotes a 
mobile-first approach and provides mobile (client) and server SDKs. 


* The server SDKs enable you to register users with OST Platform.
* The client SDKs provide the additional support required for 
the ownership and management of Brand Tokens by users so 
that they can create keys and control their tokens. 

### Users Module 

To register users with OST Platform, you can use the services provided in the Users module.

Create a User with OST Platform:

```Go
r, _ := client.Users.Create()
res, _ := json.Marshal(r)
fmt.Println(string(res))
```

Get User Detail:

```Go
r, _ := client.Users.Get("b6504ca4-9263-4998-8036-c90e648c48de")
res, _ := json.Marshal(r)
fmt.Println(string(res))
```

Get Users List:

```Go
r, _ := client.Users.GetList()
res, _ := json.Marshal(r)
fmt.Println(string(res))
```

##  Coverage

Currently the following services are supported:

- [x] Create User
- [x] Get User
- [x] List Users
- [x] Register Device
- [x] Get Device
- [ ] List Devices in the Economy
- [x] List Devices of a User
- [x] Get Session
- [ ] List Sessions in the Economy
- [x] List Sessions of a User
- [ ] List All Rules
- [ ] Get Price Point
- [ ] Execute Company to User Transaction
- [ ] Get Transaction
- [ ] List Transactions in the economy
- [ ] List Transactions of a user
- [ ] Get Balance
- [ ] Get Token
- [ ] Get Recovery Owner
- [ ] Get Chain Information
- [ ] Get Device Manager
- [ ] Get Available Base Token
- [ ] Create WebHook
- [ ] Get WebHook
- [ ] List WebHooks
- [ ] Update WebHook
- [ ] Delete WebHook


## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
