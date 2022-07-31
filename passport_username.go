package goaccount

type PassportUsername struct {
	AccountID      Identity `json:"AccountID" bson:"AccountID"`
	Username       string   `json:"Username" bson:"Username"`
	PasswordShadow string   `json:"PasswordShadow" bson:"PasswordShadow"`
	Salt           string   `json:"Salt" bson:"Salt"`
}

//MARK: implement IPassport
func (doc *PassportUsername) GetID() string { //documentdb interface
	return doc.Username
}

func (doc *PassportUsername) GetPassportID() Identity {
	//account id in provider system
	return Identity(doc.Username)
}
func (doc *PassportUsername) GetProvider() Provider {
	return ProviderUsername
}

func (doc *PassportUsername) GetAccountID() Identity {
	return doc.AccountID
}

//set binding to account
func (doc *PassportUsername) SetAccountID(accountID Identity) {
	doc.AccountID = accountID
}

func (doc *PassportUsername) CloneEmpty() IPassport {
	return &PassportUsername{
		Username:       doc.Username,
		PasswordShadow: doc.PasswordShadow,
		Salt:           doc.Salt,
	}
}
