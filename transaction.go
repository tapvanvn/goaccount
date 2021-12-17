package goaccount

import "context"

func SignupFromPassport(passport IPassport, context context.Context) (IAccount, error) {
	clone := passport.CloneEmpty()

	err := GetPassport(passport.GetPassportID(), clone)
	if err == nil || !__engine.GetDocumentPool().IsNoRecordError(err) {

		return nil, ErrAccountExisted
	}
	acc := __account_provider.NewAccount(context)
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

func GetAccountIDFromPassport(passport IPassport) (Identity, error) {

	clone := passport.CloneEmpty()

	err := GetPassport(passport.GetPassportID(), clone)

	if err != nil || __engine.GetDocumentPool().IsNoRecordError(err) {

		return IdentityEmpty, ErrAccountNotExisted
	}
	return clone.GetAccountID(), nil
}
