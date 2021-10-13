package box

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type Box struct {
	//dao    *dao
	syncer *syncer
}

func NewBox(syncer *syncer) *Box {
	return &Box{syncer: syncer}
}

func (b *Box) SyncSummoners(address string) ([]*types.Summoner, error) {
	// sync summoner
	summoners, err := b.syncer.Sync(address)

	if err != nil {
		log.Logger.Error("box sync summoner failed", zap.String("address", address), zap.Error(err))
	}

	log.Logger.Info("box sync success", zap.String("address", address), zap.Int("summoner amount", len(summoners)))

	// Currently, there is no need to consider the handling of storage, so we will not deal with it temporarily
	//_ = b.dao.Save(address, summoners)

	return summoners, nil
}
