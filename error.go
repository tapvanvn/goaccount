package goaccount

import "errors"

var ErrRequireDocDB = errors.New("Require atleast a document db")
var ErrRequireAccountProvider = errors.New("Require Account provider")

//MARK: ACCOUNT
var ErrAccountExisted = errors.New("Account existed")
var ErrAccountNotExisted = errors.New("Account is not existed")

//MARK: 2FA
var ErrAccount2FAExisted = errors.New("Account already has 2FA")
var ErrAccount2FANotExisted = errors.New("Account is not active 2FA")
