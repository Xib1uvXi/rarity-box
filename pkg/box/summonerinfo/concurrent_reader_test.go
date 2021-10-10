package summonerinfo

import (
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_concurrentReader_SummonerInfo(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	rib, err := contractlib.NewRarityLib(client)
	assert.NoError(t, err)

	cr := NewConcurrentReader(2, rib)

	info, err := cr.Summoners("0x6D81145629f154dBf07fDAb18D2892818626FeCF", []uint64{2062009, 2063971, 2222628})
	assert.NoError(t, err)

	assert.Len(t, info, 3)
}
