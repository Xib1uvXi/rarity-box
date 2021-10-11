package storage

import (
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

func Test_gormDao_Save(t *testing.T) {
	home, err := os.UserHomeDir()
	assert.NoError(t, err)

	dbPath := filepath.Join(home, "test_gorm.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	assert.NoError(t, err)

	orm := gormDao{db}

	_ = orm.AutoMigrate(&types.Addresses{}, &types.Summoner{})

	assert.NoError(t, orm.Save("0x6D81145629f154dBf07fDAb18D2892818626FeCF", []*types.Summoner{{
		Address:       "0x6D81145629f154dBf07fDAb18D2892818626FeCF",
		TokenID:       1,
		ClassID:       "123",
		Level:         "123",
		Xp:            "123",
		NextLevelXp:   "123",
		Gold:          "123",
		GoldClaimable: "123",
		AdventureLog:  "123",
		DungeonLog:    "123",
		DungeonScout:  "123",
	}, {
		Address:       "0x6D81145629f154dBf07fDAb18D2892818626FeCF",
		TokenID:       2,
		ClassID:       "123",
		Level:         "123",
		Xp:            "123",
		NextLevelXp:   "123",
		Gold:          "123",
		GoldClaimable: "123",
		AdventureLog:  "123",
		DungeonLog:    "123",
		DungeonScout:  "123",
	}}))
}
