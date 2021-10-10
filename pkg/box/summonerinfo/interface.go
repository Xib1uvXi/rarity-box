package summonerinfo

import "github.com/Xib1uvXi/rarity-box/pkg/types"

type chainReader interface {
	SummonerInfo(tokenID uint64, address string) (*types.Summoner, error)
}
