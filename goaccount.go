package goaccount

import (
	"strings"

	"github.com/tapvanvn/goauth"
	"github.com/tapvanvn/goauth/eth"
	"github.com/tapvanvn/goauth/momo"
	"github.com/tapvanvn/godbengine/engine"
)

var __engine *engine.Engine = nil
var __auth *goauth.Auth = nil
var __account_provider IAccountProvider = nil
var __auto_create_account = map[Provider]bool{}

func InitGoAccount(eng *engine.Engine, dbPrefix string, accountProvider IAccountProvider, config *Config) error {

	docdb := eng.GetDocumentPool()
	dbPrefix = strings.TrimSpace(dbPrefix)

	if len(dbPrefix) > 0 {

		setCollectionPrefix(dbPrefix)
	}

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
		__auto_create_account[ProviderEthereum] = config.Ethereum.AutoCreateAccount
		__auth.RegClient(client)
	}

	if config.MomoMiniApp != nil {

		client, err := momo.NewMiniappClient(config.MomoMiniApp.AppID)
		if err != nil {
			return err
		}
		__auto_create_account[ProviderMomo] = config.MomoMiniApp.AutoCreateAccount
		__auth.RegClient(client)
	}

	return nil
}

func IsAutoCreateAccount(passportProvider Provider) bool {
	if val, has := __auto_create_account[passportProvider]; has {
		return val
	}
	return false
}

func GetClientType(provider Provider) goauth.ClientType {
	switch provider {
	case ProviderEthereum:
		return goauth.ClientTypeEthereum
	case ProviderMomo:
		return goauth.ClientTypeMomoMiniapp
	default:
		return goauth.ClientTypeUnknown
	}
}

func StartSession(passport IPassport) (goauth.ISession, error) {

	//fmt.Println("start session", clientType, clientAccountID)
	authClientType := GetClientType(passport.GetProvider())
	if authClientType == goauth.ClientTypeUnknown {
		return nil, goauth.ErrClientNotFound
	}
	return __auth.BeginSession(authClientType, goauth.AccountID(passport.GetPassportID()))
}

func VerifySession(session goauth.ISession, response goauth.IResponse) (bool, error) {

	return __auth.VerifySession(session.GetClientType(), session, response)
}

func Authenticate(passport IPassport, response goauth.IResponse) (bool, error) {
	authClientType := GetClientType(passport.GetProvider())
	if authClientType == goauth.ClientTypeUnknown {
		return false, goauth.ErrClientNotFound
	}
	return __auth.Authenticate(authClientType, goauth.AccountID(passport.GetPassportID()), response)
}
