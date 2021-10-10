package summonerinfo

import (
	"fmt"
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
	"sync"
)

type concurrentReader struct {
	maxConcurrent int
	reader        chainReader
}

func NewConcurrentReader(maxConcurrent int, reader chainReader) *concurrentReader {
	return &concurrentReader{maxConcurrent: maxConcurrent, reader: reader}
}

func (g *concurrentReader) Summoners(address string, tokenIDs []uint64) ([]*types.Summoner, error) {
	resultC := make(chan *types.Summoner, len(tokenIDs))

	wg := sync.WaitGroup{}
	ch := make(chan struct{}, g.maxConcurrent)

	for i := range tokenIDs {
		wg.Add(1)
		ch <- struct{}{}
		go func(tid uint64) {
			defer func() {
				wg.Done()
				<-ch
			}()

			summonerInfo, err := g.reader.SummonerInfo(tid, address)
			if err != nil {
				log.Logger.Error("async get summoners failed", zap.Uint64("tokenID", tid), zap.Error(err))
				return
			}
			resultC <- summonerInfo

		}(tokenIDs[i])
	}
	wg.Wait()

	close(resultC)
	var infos []*types.Summoner

	for i := range resultC {
		infos = append(infos, i)
	}

	if len(tokenIDs) != len(infos) {
		return nil, fmt.Errorf("concurrent get summoners has error, tokenids len: %d, summoners len: %d", len(tokenIDs), len(infos))
	}

	return infos, nil
}
