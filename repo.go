package goaccount

import (
	"strconv"
	"time"

	"github.com/tapvanvn/goauth"
)

func GetPassport(identity Identity, passport IPassport) error {

	docdb := __engine.GetDocumentPool()
	//fmt.Println("passport", CollectionPassport(passport.GetProvider()), identity)
	return docdb.Get(CollectionPassport(passport.GetProvider()), string(identity), passport)
}

func PutPassport(passport IPassport) error {

	docdb := __engine.GetDocumentPool()
	return docdb.Put(CollectionPassport(passport.GetProvider()), passport)
}

func GetAccount(identity Identity, account IAccount) error {

	docdb := __engine.GetDocumentPool()
	return docdb.Get(CollectionAccount(), string(identity), account)
}

func PutAccount(account IAccount) error {

	docdb := __engine.GetDocumentPool()
	return docdb.Put(CollectionAccount(), account)
}

//MARK: repo auth
type RepoAuth struct {
}

func (repo *RepoAuth) NewSessionID() goauth.SessionID {
	//TODO: apply gosession if needed here
	now := time.Now().UnixNano()
	return goauth.SessionID(strconv.FormatInt(now, 10))
}
