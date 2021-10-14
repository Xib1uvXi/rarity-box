package appbackend

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/hnlq715/golang-lru"
	"go.uber.org/zap"
	"time"
)

type SummonersSyncer interface {
	SyncSummoners(address string) ([]*types.Summoner, error)
}

type Tasker interface {
	Tasks(address string) ([]*types.Task, error)
}

type Operator interface {
	IsOperator(address string) (bool, error)
	SetOperator() (*types.RawTxParam, error)
}

type TxBuilder interface {
	Adventure(from string, ids []uint64) (*types.RawTxParam, error)
	Levelup(from string, ids []uint64) (*types.RawTxParam, error)
	ClaimGold(from string, ids []uint64) (*types.RawTxParam, error)
	Dungeon(from string, ids []uint64) (*types.RawTxParam, error)
}

type srv struct {
	summonersSyncer SummonersSyncer
	summonersCache *lru.ARCCache

	tasker        Tasker
	operator      Operator
	operatorCache *lru.ARCCache
}

func NewSrv(summonersSyncer SummonersSyncer, tasker Tasker, operator Operator) *srv {
	operatorCache, err := lru.NewARCWithExpire(100, 15*time.Second)
	if err != nil {
		panic(err)
	}

	summonersCache, err := lru.NewARCWithExpire(100, 10*time.Second)
	if err != nil {
		panic(err)
	}

	return &srv{summonersSyncer: summonersSyncer, tasker: tasker, operatorCache: operatorCache, operator: operator, summonersCache: summonersCache}
}

func (s *srv) Summoners(address string) ([]*types.Summoner, error) {
	r, ok := s.summonersCache.Get(address)
	if ok {
		return r.([]*types.Summoner), nil
	}

	summoners, err := s.summonersSyncer.SyncSummoners(address)
	if err != nil {
		log.Logger.Error("appbackend get Summoners failed", zap.String("address", address), zap.Error(err))
		return nil, err
	}

	s.summonersCache.Add(address, summoners)

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

func (s *srv) IsOperator(address string) (bool, error) {
	r, ok := s.operatorCache.Get(address)
	if ok {
		return r.(bool), nil
	}

	result, err := s.operator.IsOperator(address)
	if err != nil {
		log.Logger.Error("appbackend get is operator failed", zap.String("address", address), zap.Error(err))
		return false, err
	}

	s.operatorCache.Add(address, result)

	return result, nil
}

func (s *srv) SetOperator() (*types.RawTxParam, error) {
	param, err := s.operator.SetOperator()
	if err != nil {
		log.Logger.Error("appbackend set operator failed", zap.Error(err))
		return nil, err
	}

	return param, nil
}
