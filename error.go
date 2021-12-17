package goaccount

import "errors"

var ErrRequireDocDB = errors.New("Require atleast a document db")
var ErrRequireAccountProvider = errors.New("Require Account provider")

//MARK: ACCOUNT

var ErrAccountExisted = errors.New("Account existed")
var ErrAccountNotExisted = errors.New("Account is not existed")
