package storage

import (
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormDao struct {
	*gorm.DB
}

func (dao *gormDao) Save(address string, summoners []*types.Summoner) error {
	var addr types.Addresses

	if err := dao.Model(&types.Addresses{}).Where("address = ?", address).FirstOrCreate(&addr).Error; err != nil {
		return err
	}

	for i := range summoners {
		if err := dao.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "token_id"}},
			UpdateAll: true,
		}).Create(summoners[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
