package summonerstatuschecker

import (
	"errors"
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
	"math/big"
	"time"
)

var BigIntSetStringErr = errors.New("SetString: error")

type Checkers struct {
}

func (c *Checkers) Status(summoner *types.Summoner) *types.SummonerStatus {
	status := &types.SummonerStatus{TokenID: summoner.TokenID}

	// can adventure
	status.CanAdventure = c.canAdventure(summoner)

	// can levelup
	status.CanLevelUp = c.canLevelup(summoner)

	// gold claimable
	status.GoldClaimable = c.goldClaimable(summoner)

	// attributes is set
	status.AttributesIsSet = !summoner.PointNotSet()

	// can dungeon adventure
	status.CanDungeonAdventure = c.canDungeon(summoner)

	return status
}

func (c *Checkers) canAdventure(summoner *types.Summoner) bool {
	alog, err := c.convertStringToBigInt(summoner.AdventureLog)
	if err != nil {
		log.Logger.Error("adventure log convert failed", zap.String("address", summoner.Address), zap.Uint64("token id", summoner.TokenID), zap.String("raw value", summoner.AdventureLog))
		return false
	}

	// can adventure
	// Just need to judge whether the log is less than the current timestamp,
	// the adventure log in the contract records the timestamp 24 hours later
	// special cases, New Summoner adventure log is zero
	if big.NewInt(time.Now().Unix()).Cmp(alog) == 1 {
		return true
	}

	return false
}

func (c *Checkers) canLevelup(summoner *types.Summoner) bool {
	nextLevelRequireXp, err := c.convertStringToBigInt(summoner.NextLevelXp)
	if err != nil {
		log.Logger.Error("next level require convert failed", zap.String("address", summoner.Address), zap.Uint64("token id", summoner.TokenID), zap.String("raw value", summoner.NextLevelXp))
		return false
	}

	xp, err := c.convertStringToBigInt(summoner.Xp)
	if err != nil {
		log.Logger.Error("xp require convert failed", zap.String("address", summoner.Address), zap.Uint64("token id", summoner.TokenID), zap.String("raw value", summoner.Xp))
		return false
	}

	if nextLevelRequireXp.Cmp(xp) == 0 || nextLevelRequireXp.Cmp(xp) == -1 {
		// can levelup
		return true
	}

	return false
}

func (c *Checkers) goldClaimable(summoner *types.Summoner) bool {
	claimable, err := c.convertStringToBigInt(summoner.GoldClaimable)
	if err != nil {
		log.Logger.Error("gold claimable convert failed", zap.String("address", summoner.Address), zap.Uint64("token id", summoner.TokenID), zap.String("raw value", summoner.GoldClaimable))
		return false
	}

	if claimable.Cmp(big.NewInt(0)) == 0 {
		return false
	}

	return true
}

func (c *Checkers) canDungeon(summoner *types.Summoner) bool {
	scout, err := c.convertStringToBigInt(summoner.DungeonScout)
	if err != nil {
		log.Logger.Error("dungeon scout convert failed", zap.String("address", summoner.Address), zap.Uint64("token id", summoner.TokenID), zap.String("raw value", summoner.DungeonScout))
		return false
	}

	if scout.Cmp(big.NewInt(0)) != 1 {
		return false
	}

	alog, err := c.convertStringToBigInt(summoner.DungeonLog)
	if err != nil {
		log.Logger.Error("dungeon log convert failed", zap.String("address", summoner.Address), zap.Uint64("token id", summoner.TokenID), zap.String("raw value", summoner.DungeonLog))
		return false
	}

	// can adventure
	// Just need to judge whether the log is less than the current timestamp,
	// the adventure log in the contract records the timestamp 24 hours later
	// special cases, New Summoner adventure log is zero
	if big.NewInt(time.Now().Unix()).Cmp(alog) == 1 {
		return true
	}

	return false
}

func (c *Checkers) convertStringToBigInt(data string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(data, 10)

	if !ok {
		return nil, BigIntSetStringErr
	}

	return n, nil
}
