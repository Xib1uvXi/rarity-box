package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"math/big"
)

func (r *RarityLib) getRarityInfo(info *types.SummonerInfo, sid *big.Int) (*types.SummonerInfo, error) {
	rawInfo, err := r.rarity.Summoner(nil, sid)
	if err != nil {
		return nil, err
	}

	info.ClassID = rawInfo.Class.String()
	info.Xp = rawInfo.Xp.String()
	info.AdventureLog = rawInfo.Log.String()
	info.Level = rawInfo.Level.String()

	// can levelup
	nextLevelRequireXp, err := r.rarity.XpRequired(nil, rawInfo.Level)
	if err != nil {
		return nil, err
	}

	if nextLevelRequireXp.Cmp(rawInfo.Xp) == 0 || nextLevelRequireXp.Cmp(rawInfo.Xp) == -1 {
		// can levelup
		info.CanLevelUp = true
	}
}
