package box

import "github.com/Xib1uvXi/rarity-box/pkg/box/tokenidsync"

func BuildSyncer(summonerReader SummonerReader) *syncer {
	return NewSyncer(summonerReader, tokenidsync.NewThegraphTokenIDSynchronizer())
}

func BuildBox(summonerReader SummonerReader) *Box {
	return NewBox(BuildSyncer(summonerReader))
}
