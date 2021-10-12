package tx

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ToPublicKeyErr = errors.New("pk to public key failed")
)

// account Ethereum account
type account struct {
	nonce      uint64
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func NewAccount(pkStr string) (*account, error) {
	privateKey, address, err := initPK(pkStr)
	if err != nil {
		return nil, err
	}
	account := &account{privateKey: privateKey, address: address}

	return account, nil
}

// init private key
func initPK(pkStr string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(pkStr)
	if err != nil {
		return nil, common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, ToPublicKeyErr
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, address, nil
}

func (a *account) PrivateKey() *ecdsa.PrivateKey {
	return a.privateKey
}

func (a *account) Address() common.Address {
	return a.address
}

