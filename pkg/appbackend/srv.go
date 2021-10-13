package appbackend

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type SummonersSyncer interface {
	SyncSummoners(address string) ([]*types.Summoner, error)
}

type Tasker interface {
	Tasks(address string) ([]*types.Task, error)
}

type srv struct {
	summonersSyncer SummonersSyncer
	tasker          Tasker
}

func NewSrv(summonersSyncer SummonersSyncer, tasker Tasker) *srv {
	return &srv{summonersSyncer: summonersSyncer, tasker: tasker}
}

func (s *srv) Summoners(address string) ([]*types.Summoner, error) {
	summoners, err := s.summonersSyncer.SyncSummoners(address)
	if err != nil {
		log.Logger.Error("appbackend get Summoners failed", zap.String("address", address), zap.Error(err))
		return nil, err
	}

	return summoners, nil
}

func (s *srv) Tasks(address string) ([]*types.Task, error) {
	tasks, err := s.tasker.Tasks(address)
	if err != nil {
		log.Logger.Error("appbackend get tasks failed", zap.String("address", address), zap.Error(err))
		return nil, err
	}

	for i := range tasks {
		tasks[i].IDs = []uint64{}
	}

	return tasks, nil
}
