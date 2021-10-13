package tokenidsync

import (
	"fmt"
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

const thegraphURL = "https://api.thegraph.com/subgraphs/name/rarity-adventure/rarity"

// thegraphTokenIDSynchronizer use The Graph which is an indexing protocol for querying networks like Ethereum and IPFS
// Anyone can build and publish open APIs, called subgraphs, making data easily accessible
// more info: https://thegraph.com/en/
type thegraphTokenIDSynchronizer struct {
	client *resty.Client
}

func NewThegraphTokenIDSynchronizer() *thegraphTokenIDSynchronizer {
	return &thegraphTokenIDSynchronizer{client: resty.New()}
}

func (ds *thegraphTokenIDSynchronizer) Sync(address string) ([]uint64, error) {
	var response ThegraphResp

	_, err := ds.client.R().
		SetResult(&response).
		SetHeader("Accept", "application/json").
		SetBody([]byte(ds.genGraphql(address))).
		Post(thegraphURL)

	if err != nil {
		log.Logger.Error("request thegraph api failed", zap.Error(err))
		return nil, err
	}

	ids, err := response.dto()

	if err != nil {
		log.Logger.Error("thegraph api resp dto failed", zap.Error(err))
		return nil, err
	}

	return ids, nil
}

func (ds *thegraphTokenIDSynchronizer) genGraphql(address string) string {
	return fmt.Sprintf(`{"variables":{},"query":"{\n  summoners(\n    first: 1000\n    where: {owner: \"%s\"}\n  ) {\n    id\n    owner\n    _class\n    _level\n    __typename\n  }\n}\n"}`, address)
}

type ThegraphResp struct {
	Data struct {
		Summoners []struct {
			ID    string `json:"id"`
			Owner string `json:"owner"`
		} `json:"summoners"`
	} `json:"data"`
}

// data to object
func (r *ThegraphResp) dto() ([]uint64, error) {
	var ids []uint64

	for i := range r.Data.Summoners {
		hexID := r.Data.Summoners[i].ID

		id, err := hexutil.DecodeUint64(hexID)

		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}
