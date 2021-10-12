package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/caller-contract/caller"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/attributes"
	craft_i "github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/craft-i"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/gold"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/rarity"
	"github.com/ethereum/go-ethereum/common"
)

const (
	RarityContractAddress     = "0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb"
	AttributesContractAddress = "0xB5F5AF1087A8DA62A23b08C00C6ec9af21F397a1"
	CraftIContractAddress     = "0x2A0F1cB17680161cF255348dDFDeE94ea8Ca196A"
	GoldContractAddress       = "0x2069B76Afe6b734Fb65D1d099E7ec64ee9CC76B2"
	CallerContractAddress     = "0x3B8d1f9569ec2B8C3dD76847B944bB1b0c998A38"
)

func InitRarityLib(rlib *RarityLib) error {
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

	callerContract, err := caller.NewCaller(common.HexToAddress(CallerContractAddress), rlib.conn)
	rlib.caller = callerContract

	return nil
}
