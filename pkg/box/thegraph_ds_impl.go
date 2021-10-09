package box

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
)

const thegraphURL = "https://api.thegraph.com/subgraphs/name/rarity-adventure/rarity"

// thegraphDataSynchronizer use The Graph which is an indexing protocol for querying networks like Ethereum and IPFS
// Anyone can build and publish open APIs, called subgraphs, making data easily accessible
// more info: https://thegraph.com/en/
type thegraphDataSynchronizer struct {
	dao dao
}

func newThegraphDataSynchronizer(dao dao) *thegraphDataSynchronizer {
	return &thegraphDataSynchronizer{dao: dao}
}

func (ds *thegraphDataSynchronizer) Sync(address string) error {
	var response ThegraphResp

	client := resty.New()
	_, err := client.R().
		SetResult(&response).
		SetHeader("Accept", "application/json").
		SetBody([]byte(ds.genGraphql(address))).
		Post(thegraphURL)

	if err != nil {
		return err
	}

	ids, err := response.dto()

	if err != nil {
		return err
	}

	return ds.dao.SaveOrUpdate(address, ids)
}

func (ds *thegraphDataSynchronizer) genGraphql(address string) string {
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
