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
