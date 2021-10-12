package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
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

const (
	RarityContractAddress     = "0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb"
	AttributesContractAddress = "0xB5F5AF1087A8DA62A23b08C00C6ec9af21F397a1"
	CraftIContractAddress     = "0x2A0F1cB17680161cF255348dDFDeE94ea8Ca196A"
	GoldContractAddress       = "0x2069B76Afe6b734Fb65D1d099E7ec64ee9CC76B2"
)

type RarityLib struct {
	conn       *ethclient.Client
	attributes *attributes.Attributes
	dungeon    *craft_i.CraftI
	gold       *gold.Gold
	rarity     *rarity.Rarity
	txSender   TxSender
}

func NewRarityLib(conn *ethclient.Client, sender TxSender) (*RarityLib, error) {
	rlib := &RarityLib{conn: conn, txSender: sender}

	if err := initRarityLib(rlib); err != nil {
		log.Logger.Error("init rarity lib failed", zap.Error(err))
		return nil, err
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

func initRarityLib(rlib *RarityLib) error {
	rarityContract, err := rarity.NewRarity(common.HexToAddress(RarityContractAddress), rlib.conn)
	if err != nil {
		return err
	}

	rlib.rarity = rarityContract

	attributesContract, err := attributes.NewAttributes(common.HexToAddress(AttributesContractAddress), rlib.conn)
	if err != nil {
		return err
	}

	rlib.attributes = attributesContract

	dungeonContract, err := craft_i.NewCraftI(common.HexToAddress(CraftIContractAddress), rlib.conn)
	if err != nil {
		return err
	}

	rlib.dungeon = dungeonContract

	goldContract, err := gold.NewGold(common.HexToAddress(GoldContractAddress), rlib.conn)
	if err != nil {
		return err
	}

	rlib.gold = goldContract

	return nil
}
