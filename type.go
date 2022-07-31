package goaccount

import (
	"context"
	"fmt"
)

var __prefix string = ""

type Provider string
type Identity string

const (
	ProviderEthereum = Provider("ethereum")
	ProviderEmail    = Provider("email")
	ProviderUsername = Provider("username")
	ProviderApple    = Provider("apple")
	ProviderGoogle   = Provider("google")
	ProviderMomo     = Provider("momo")

	IdentityEmpty = Identity("")
)

var (
	__prefix_passport = "passport"
	__prefix_account  = "account"
)

func setCollectionPrefix(prefix string) {
	__prefix_passport = fmt.Sprintf("%s_passport", prefix)
	__prefix_account = fmt.Sprintf("%s_account", prefix)
	__prefix = prefix
}

func CollectionPassport(provider Provider) string {

	return fmt.Sprintf("%s_%s", __prefix_passport, provider)
}
func CollectionAccount() string {
	return __prefix_account
}

type IPassport interface {
	GetID() string           //documentdb interface
	GetPassportID() Identity //account id in provider system
	GetProvider() Provider   //

	GetAccountID() Identity
	SetAccountID(accountID Identity) //set binding to account

	CloneEmpty() IPassport //clone provider but setting other attributes to empty
}

type IAccount interface {
	GetID() string          //documentdb interface
	GetAccountID() Identity //account id in host system
}

type IAccountProvider interface {
	NewAccount(context context.Context) IAccount //create new account
}
