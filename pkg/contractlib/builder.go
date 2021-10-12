package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/tx"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BuildV1RarityLib(pkStr string, conn *ethclient.Client) (*RarityLib, error) {
	txSender, err := tx.NewFTMTXSession(pkStr, conn)
	if err != nil {
		return nil, err
	}

	return NewRarityLib(conn, txSender)
}

func BuildV1CallerLib(pkStr string, conn *ethclient.Client, hasFee bool) (*CallerLib, error) {
	rlib, err := BuildV1RarityLib(pkStr, conn)
	if err != nil {
		return nil, err
	}

	return NewCallerLib(rlib, hasFee)
}
