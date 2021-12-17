package goaccount

func NewPassportEthereum(walletAddress string) *PassportEthereum {

	return &PassportEthereum{

		WalletAddress: walletAddress,
	}
}

type PassportEthereum struct {
	WalletAddress string   `json:"WalletAddress" bson:"WalletAddress"`
	AccountID     Identity `json:"AccountID" bson:"AccountID"`
}

func (doc *PassportEthereum) GetID() string {
	return doc.WalletAddress
}
func (doc *PassportEthereum) GetPassportID() Identity {
	return Identity(doc.WalletAddress)
}
func (doc *PassportEthereum) GetProvider() Provider {
	return ProviderEthereum
}

func (doc *PassportEthereum) GetAccountID() Identity {
	return doc.AccountID
}
func (doc *PassportEthereum) SetAccountID(accountID Identity) {
	doc.AccountID = accountID
}

func (doc *PassportEthereum) CloneEmpty() IPassport {
	return &PassportEthereum{
		WalletAddress: doc.WalletAddress,
	}
}
