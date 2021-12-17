package goaccount

import (
	"context"
	"fmt"
)

var __prefix string = ""

type Provider string
type Identity string

const (
	ProviderEthereum     = Provider("ethereum")
	ProviderEmail        = Provider("email")
	ProviderUsernamePass = Provider("userpass")
	ProviderApple        = Provider("apple")
	ProviderGoogle       = Provider("google")

	IdentityEmpty = Identity("")
)

func CollectionPassport(provider Provider) string {

	return fmt.Sprintf("%s_passport_%s", __prefix, provider)
}
func CollectionAccount() string {
	return fmt.Sprintf("%s_account", __prefix)
}

type IPassport interface {
	GetID() string           //documentdb interface
	GetPassportID() Identity //account id in provider system
	GetProvider() Provider   //

	GetAccountID() Identity
	SetAccountID(accountID Identity) //set binding to account

	CloneEmpty() IPassport //clone provider but setting others attribute to empty
}

type IAccount interface {
	GetID() string          //documentdb interface
	GetAccountID() Identity //account id in host system
}

type IAccountProvider interface {
	NewAccount(context context.Context) IAccount //create new account
}
