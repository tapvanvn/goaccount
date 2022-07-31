package goaccount

type PassportMomo struct {
	AccountID Identity `json:"AccountID" bson:"AccountID"`
	//AppID     string
	MiniAppUserId string `json:"MiniAppUserId" bson:"MiniAppUserId"`
	//PartnerUserID string `json:"PartnerUserId" bson:"PartnerUserId"`
	//AuthCode string `json:"AuthCode" bson:"AuthCode"`
}

//MARK: implement IPassport
func (doc *PassportMomo) GetID() string { //documentdb interface
	return doc.MiniAppUserId
}

func (doc *PassportMomo) GetPassportID() Identity {
	//account id in provider system
	return Identity(doc.MiniAppUserId)
}
func (doc *PassportMomo) GetProvider() Provider {
	return ProviderMomo
}

func (doc *PassportMomo) GetAccountID() Identity {
	return doc.AccountID
}

//set binding to account
func (doc *PassportMomo) SetAccountID(accountID Identity) {
	doc.AccountID = accountID
}

func (doc *PassportMomo) CloneEmpty() IPassport {
	return &PassportMomo{
		MiniAppUserId: doc.MiniAppUserId,
		//AuthCode:      doc.AuthCode,
		//AppID:         doc.AppID,
	}
}
