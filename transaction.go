package goaccount

import (
	"context"

	"github.com/tapvanvn/goutil"
)

func SignupFromPassport(passport IPassport, context context.Context) (IAccount, error) {
	clone := passport.CloneEmpty()

	err := GetPassport(passport.GetPassportID(), clone)

	if err == nil || !__engine.GetDocumentPool().IsNoRecordError(err) {

		return nil, ErrAccountExisted
	}
	acc := __account_provider.NewAccount(context)

	if passport.HasTitle() {

		acc.SetTitle(passport.GetTitle())
	}

	passport.SetAccountID(acc.GetAccountID())

	trans := __engine.GetDocumentPool().MakeTransaction()

	trans.Put(CollectionAccount(), acc)

	trans.Put(CollectionPassport(passport.GetProvider()), passport)

	err = trans.Commit()

	if err != nil {

		return nil, err
	}
	return acc, nil
}

//IMPORTANT: the passport must be valid
func GetAccountIDFromPassport(passport IPassport) (Identity, error) {

	clone := passport.CloneEmpty()

	err := GetPassport(passport.GetPassportID(), clone)

	if err != nil {

		if __engine.GetDocumentPool().IsNoRecordError(err) && IsAutoCreateAccount(passport.GetProvider()) {
			//if auto create account
			if acc, err := SignupFromPassport(passport, context.Background()); err != nil {
				return IdentityEmpty, err
			} else {
				return acc.GetAccountID(), nil
			}
		}
		return IdentityEmpty, ErrAccountNotExisted
	}
	return clone.GetAccountID(), nil
}

//MARK: 2FA
func CreateAccount2FA(accountID Identity) (TwoFASecret, error) {
	acc := __account_provider.NewAccount(context.Background())
	if err := GetAccount(accountID, acc); err != nil {
		return "", err
	}
	if acc.Get2FASecret() != "" {
		return "", ErrAccount2FAExisted
	}
	secret := goutil.Gen2FASecret()
	acc.Set2FASecret(secret)
	err := PutAccount(acc)
	return TwoFASecret(secret), err
}

func ValidateAccount2FA(accountID Identity, otp TwoFAOTP) (bool, error) {
	acc := __account_provider.NewAccount(context.Background())
	if err := GetAccount(accountID, acc); err != nil {
		return false, err
	}
	secret := acc.Get2FASecret()
	if secret == "" {
		return false, ErrAccount2FANotExisted
	}
	shouldOTP, err := goutil.Get2FAOTP(secret)
	return TwoFAOTP(shouldOTP) == otp, err
}

func RemoveAccount2FA(accountID Identity) error {
	acc := __account_provider.NewAccount(context.Background())
	if err := GetAccount(accountID, acc); err != nil {
		return err
	}
	if acc.Get2FASecret() != "" {
		return ErrAccount2FAExisted
	}
	acc.Set2FASecret("")
	return PutAccount(acc)
}
