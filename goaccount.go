package goaccount

import (
	"github.com/tapvanvn/goauth"
	"github.com/tapvanvn/goauth/eth"
	"github.com/tapvanvn/godbengine/engine"
)

var __engine *engine.Engine = nil
var __auth *goauth.Auth = nil
var __account_provider IAccountProvider = nil

func InitGoAccount(eng *engine.Engine, accountProvider IAccountProvider, config *Config) error {

	docdb := eng.GetDocumentPool()

	if docdb == nil {

		return ErrRequireDocDB
	}
	if accountProvider == nil {

		return ErrRequireAccountProvider
	}
	__account_provider = accountProvider

	__engine = eng
	__auth = goauth.NewAuth(&RepoAuth{})

	if config == nil {

		return nil
	}
	if config.Ethereum != nil {

		client, err := eth.NewOwnerClient(config.Ethereum.PrivateKey)
		if err != nil {
			return err
		}
		__auth.RegClient(client)
	}
	return nil
}
