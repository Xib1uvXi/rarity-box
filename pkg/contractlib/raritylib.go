package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/attributes"
	craft_i "github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/craft-i"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/gold"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/rarity"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

const (
	RarityContractAddress     = "0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb"
	AttributesContractAddress = "0xB5F5AF1087A8DA62A23b08C00C6ec9af21F397a1"
	CraftIContractAddress     = "0x2A0F1cB17680161cF255348dDFDeE94ea8Ca196A"
	GoldContractAddress       = "0x2069B76Afe6b734Fb65D1d099E7ec64ee9CC76B2"
)

type RarityLib struct {
	conn       bind.ContractBackend
	attributes *attributes.Attributes
	dungeon    *craft_i.CraftI
	gold       *gold.Gold
	rarity     *rarity.Rarity
}

func (r *RarityLib) SummonerInfo(tokenID uint64) (*types.SummonerInfo, error) {
	info := &types.SummonerInfo{TokenID: tokenID}

	sid := big.NewInt(int64(tokenID))



	return info, nil
}
