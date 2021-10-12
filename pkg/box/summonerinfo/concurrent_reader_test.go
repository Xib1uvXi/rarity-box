package summonerinfo

import (
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_concurrentReader_SummonerInfo(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	rib, err := contractlib.NewRarityLib(client, &testSender{})
	assert.NoError(t, err)

	cr := NewConcurrentReader(2, rib)

	info, err := cr.Summoners("0x6D81145629f154dBf07fDAb18D2892818626FeCF", []uint64{2062009, 2063971, 2222628})
	assert.NoError(t, err)

	assert.Len(t, info, 3)
}

type testSender struct {
}

func (t *testSender) SendTx(txExecutor func(opts *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
	panic("implement me")
}

func (t *testSender) SendTxWaitConfirm(txExecutor func(opts *bind.TransactOpts) (*types.Transaction, error)) error {
	panic("implement me")
}

func (t *testSender) Address() common.Address {
	panic("implement me")
}
