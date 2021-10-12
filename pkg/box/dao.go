package box

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type Storage interface {
	Save(address string, summoners []*types.Summoner) error
	SummonersByAddress(address string) ([]*types.Summoner, error)
}

type dao struct {
	db Storage
}

func NewDao(db Storage) *dao {
	return &dao{db: db}
}

func (d *dao) Summoners(address string) ([]*types.Summoner, error) {
	summoners, err := d.db.SummonersByAddress(address)
	if err != nil {
		log.Logger.Error("storage get summoners failed", zap.String("address", address), zap.Error(err))
		return nil, err
	}

	return summoners, nil
}

func (d *dao) Save(address string, summoners []*types.Summoner) error {
	err := d.db.Save(address, summoners)
	if err != nil {
		log.Logger.Error("storage save summoners failed", zap.String("address", address), zap.Error(err))
		return err
	}

	return nil
}
