package box

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type SummonerReader interface {
	Summoners(address string, tokenIDs []uint64) ([]*types.Summoner, error)
}

type AssetsIDSynchronizer interface {
	Sync(address string) ([]uint64, error)
}

type syncer struct {
	summonerInfoGetter SummonerReader
	tokenIDsGetter     AssetsIDSynchronizer
}

func (s *syncer) SetTokenIDsGetter(tokenIDsGetter AssetsIDSynchronizer) {
	s.tokenIDsGetter = tokenIDsGetter
}

func (s *syncer) SetSummonerInfoGetter(summonerInfoGetter SummonerReader) {
	s.summonerInfoGetter = summonerInfoGetter
}

func (s *syncer) Sync(address string) ([]*types.Summoner, error) {
	ids, err := s.tokenIDsGetter.Sync(address)

	if err != nil {
		log.Logger.Error("syncer get address summoner token id failed", zap.String("address", address), zap.Error(err))
		return nil, err
	}

	summoners, err := s.summonerInfoGetter.Summoners(address, ids)

	if err != nil {
		log.Logger.Error("syncer get summoner failed", zap.String("address", address), zap.Error(err))
		return nil, err
	}

	return summoners, nil
}
