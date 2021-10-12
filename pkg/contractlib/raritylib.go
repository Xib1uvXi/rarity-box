package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/caller-contract/caller"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/attributes"
	craft_i "github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/craft-i"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/gold"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/rarity"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
)

type TxSender interface {
	SendTx(txExecutor func(opts *bind.TransactOpts) (*ethTypes.Transaction, error)) (*ethTypes.Transaction, error)
	SendTxWaitConfirm(txExecutor func(opts *bind.TransactOpts) (*ethTypes.Transaction, error)) error
	Address() common.Address
}

type RarityLib struct {
	conn       *ethclient.Client
	attributes *attributes.Attributes
	dungeon    *craft_i.CraftI
	gold       *gold.Gold
	rarity     *rarity.Rarity
	caller     *caller.Caller
	txSender   TxSender
}

func NewRarityLib(conn *ethclient.Client, sender TxSender) (*RarityLib, error) {
	rlib := &RarityLib{conn: conn, txSender: sender}

	if err := InitRarityLib(rlib); err != nil {
		log.Logger.Error("init rarity lib failed", zap.Error(err))
		return nil, err
	}

	callerIsOperator, err := rlib.checkCallerIsOperator(rlib.txSender.Address().String())
	if err != nil {
		return nil, err
	}

	if !callerIsOperator {
		if err := rlib.callerSetOperator(); err != nil {
			return nil, err
		}
	}

	return rlib, nil
}

func (r *RarityLib) SummonerInfo(tokenID uint64, address string) (*types.Summoner, error) {
	info := &types.Summoner{TokenID: tokenID, Address: address}

	sid := big.NewInt(int64(tokenID))

	if err := r.getRarityInfo(info, sid, common.HexToAddress(address)); err != nil {
		log.Logger.Error("get rarity info failed", zap.Error(err))
		return nil, err
	}

	return info, nil
}
