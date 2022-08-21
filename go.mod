module github.com/tapvanvn/goaccount

require (
	github.com/tapvanvn/goauth v0.0.1-build.41
	github.com/tapvanvn/godbengine v1.4.9-build.35
)

require (
	github.com/tapvanvn/gomomo v0.0.1-build.3 // indirect
	github.com/tapvanvn/gorouter/v2 v2.0.9-build.12 // indirect
)

replace (
	github.com/tapvanvn/goauth => ../goauth
	github.com/tapvanvn/gomomo => ../../2022/gomomo
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/go-ethereum v1.10.21 // indirect
	github.com/miguelmota/go-solidity-sha3 v0.1.1 // indirect
	github.com/tapvanvn/gocondition v1.0.0-alpha.1 // indirect
	github.com/tapvanvn/goutil v0.0.18-build.20 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
)

go 1.17
