package goaccount

import (
	"context"
	"fmt"
)

var __prefix string = ""

type Provider string
type Identity string
type TwoFASecret string
type TwoFAOTP string

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
	HasTitle() bool          //incase this passport can provide a title for account
	GetTitle() string        //if passport has title then provide it

	GetAccountID() Identity
	SetAccountID(accountID Identity) //set binding to account

	CloneEmpty() IPassport //clone provider but setting other attributes to empty
}

type IAccount interface {
	GetID() string          //documentdb interface
	GetAccountID() Identity //account id in host system
	SetTitle(string)        //Set account title
	//2FA
	Get2FASecret() string
	Set2FASecret(secret string)
}

type IAccountProvider interface {
	NewAccount(context context.Context) IAccount //create new account
}
