package types

import (
	"gorm.io/gorm"
)

type Addresses struct {
	gorm.Model
	Address string `json:"address"`
}

type SummonerInfo struct {
	gorm.Model
	Address      string `json:"address"`
	TokenID      uint64 `json:"token_id"`
	ClassID      string `json:"class_id"`
	Level        string `json:"level"`
	Xp           string `json:"xp"`
	Gold         string `json:"gold"`
	AdventureLog string `json:"adventure_log"`
	DungeonLog   string `json:"dungeon_log"`
	Point
}

type SummonerStatus struct {
	SummonerID          uint   `json:"summoner_id"`
	TokenID             uint64 `json:"token_id"`
	CanAdventure        bool   `json:"can_adventure"`
	CanLevelUp          bool   `json:"can_level_up"`
	GoldClaimable       bool   `json:"gold_claimable"`
	CanDungeonAdventure bool   `json:"can_dungeon_adventure"`
	AttributesIsSet     bool   `json:"attributes_is_set"`
}

type Point struct {
	Strength     uint32 `json:"strength"`
	Dexterity    uint32 `json:"dexterity"`
	Constitution uint32 `json:"constitution"`
	Intelligence uint32 `json:"intelligence"`
	Wisdom       uint32 `json:"wisdom"`
	Charisma     uint32 `json:"charisma"`
}