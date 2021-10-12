package contractlib

import (
	"fmt"
	"github.com/Xib1uvXi/rarity-box/pkg/common/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRarityLib_SummonerInfo(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	rlib, err := NewRarityLib(client, &testSender{})
	assert.NoError(t, err)

	info, err := rlib.SummonerInfo(2062009, "0x6D81145629f154dBf07fDAb18D2892818626FeCF")

	assert.NoError(t, err)

	fmt.Println(json.StringifyJson(info))
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
