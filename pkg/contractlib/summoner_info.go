package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (r *RarityLib) getRarityInfo(info *types.Summoner, sid *big.Int, address common.Address) error {
	rawInfo, err := r.rarity.Summoner(&bind.CallOpts{
		From: address,
	}, sid)
	if err != nil {
		return err
	}

	info.ClassID = rawInfo.Class.String()
	info.Xp = rawInfo.Xp.String()
	info.AdventureLog = rawInfo.Log.String()
	info.Level = rawInfo.Level.String()

	nextLevelRequireXp, err := r.rarity.XpRequired(nil, rawInfo.Level)
	if err != nil {
		return err
	}

	info.NextLevelXp = nextLevelRequireXp.String()

	gold, err := r.gold.BalanceOf(&bind.CallOpts{
		From: address,
	}, sid)

	if err != nil {
		return  err
	}

	info.Gold = gold.String()

	goldClaimalbe, err := r.gold.Claimable(&bind.CallOpts{
		From: address,
	}, sid)

	if err != nil {
		return  err
	}

	info.GoldClaimable = goldClaimalbe.String()

	dungeonScout, err := r.dungeon.Scout(nil, sid)
	if err != nil {
		return err
	}

	info.DungeonScout = dungeonScout.String()

	dungeonLog, err := r.dungeon.AdventurersLog(nil, sid)
	if err != nil {
		return err
	}

	info.DungeonLog = dungeonLog.String()

	abscore, err := r.attributes.AbilityScores(nil, sid)
	if err != nil {
		return err
	}

	info.Point = types.Point{
		Strength:     abscore.Strength,
		Dexterity:    abscore.Dexterity,
		Constitution: abscore.Constitution,
		Intelligence: abscore.Intelligence,
		Wisdom:       abscore.Wisdom,
		Charisma:     abscore.Charisma,
	}

	return  nil
}
