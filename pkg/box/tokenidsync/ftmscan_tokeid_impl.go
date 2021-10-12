package tokenidsync

import (
	"fmt"
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"strconv"
)

type ftmscanTokenIDSynchronizer struct {
	client   *resty.Client
	contract string
	apikey   string
}

func NewFtmscanTokenIDSynchronizer() *ftmscanTokenIDSynchronizer {
	return &ftmscanTokenIDSynchronizer{contract: "0xce761D788DF608BD21bdd59d6f4B54b2e27F25Bb",
		apikey: "7HRPF8KK6XZ38GX8KRKJ3AXNG6RA4KKW15", client: resty.New()}
}

func (f *ftmscanTokenIDSynchronizer) Sync(address string) ([]uint64, error) {
	var response FtmscanResp

	_, err := f.client.R().
		SetResult(&response).
		SetHeader("Accept", "application/json").
		Get(f.apiUrl(address))

	if err != nil {
		log.Logger.Error("request ftmscan api failed", zap.Error(err))
		return nil, err
	}

	var tokenIDs []uint64

	for i := range response.Result {

		element := response.Result[i]

		id, err := strconv.ParseUint(element.TokenID, 10, 64)

		if err != nil {
			log.Logger.Error("convert string to uint64 failed", zap.String("raw value", element.TokenID), zap.Error(err))
			return nil, err
		}

		if element.To != address {
			tokenIDs = filterSlice(tokenIDs, func(x uint64) bool {
				return x == id
			})
		} else {
			tokenIDs = append(tokenIDs, id)
		}
	}

	return tokenIDs, nil
}

func (f *ftmscanTokenIDSynchronizer) apiUrl(address string) string {
	return fmt.Sprintf("https://api.ftmscan.com/api?module=account&action=tokennfttx&contractaddress=%s&address=%s&apikey=%s", f.contract, address, f.apikey)
}

type FtmscanResp struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Result  []*Result `json:"result"`
}

type Result struct {
	From    string `json:"from"`
	To      string `json:"to"`
	TokenID string `json:"tokenID"`
}

func filterSlice(s []uint64, filter func(x uint64) bool) []uint64 {
	newS := s[:0]
	for _, x := range s {
		if !filter(x) {
			newS = append(newS, x)
		}
	}
	return newS
}
