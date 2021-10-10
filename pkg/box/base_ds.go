package box

import "github.com/Xib1uvXi/rarity-box/pkg/types"

type SummonerReader interface {
	Summoners(address string, tokenIDs []uint64) ([]*types.Summoner, error)
}

type baseDataSynchronizer struct {
	summonerInfoGetter SummonerReader
}

func (b *baseDataSynchronizer) SetSummonerInfoGetter(summonerInfoGetter SummonerReader) {
	b.summonerInfoGetter = summonerInfoGetter
}
