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
	"go.uber.org/zap"
	"math/big"
)

type RarityLib struct {
	conn       bind.ContractBackend
	attributes *attributes.Attributes
	dungeon    *craft_i.CraftI
	gold       *gold.Gold
	rarity     *rarity.Rarity
}

func NewRarityLib(conn bind.ContractBackend) (*RarityLib, error) {
	rlib := &RarityLib{conn: conn}

	if err := InitRarityLib(rlib); err != nil {
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
