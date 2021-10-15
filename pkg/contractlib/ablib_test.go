package contractlib

import (
	"context"
	"fmt"
	"github.com/Xib1uvXi/rarity-box/pkg/common/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestAbLib_SetOperator(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	rlib, err := NewRarityLib(client, &testSender{})
	assert.NoError(t, err)

	ablib, err := NewAbLib(rlib)
	assert.NoError(t, err)

	param, err := ablib.SetOperator()
	assert.NoError(t, err)

	fmt.Println(param.Input)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x294d28cfc630a6d552219bdf2208fdfe84482a83977120d82567ff9983ac57fc"))
	assert.NoError(t, err)

	fmt.Println(hexutil.Encode(tx.Data()))

	fmt.Println(param.Input == hexutil.Encode(tx.Data()))

}

func TestAbLib_Adventure(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	rlib, err := NewRarityLib(client, &testSender{})
	assert.NoError(t, err)

	ablib, err := NewAbLib(rlib)
	assert.NoError(t, err)

	r, err := ablib.Adventure("0x6D81145629f154dBf07fDAb18D2892818626FeCF", []uint64{3574968, 3575125})

	assert.NoError(t, err)

	fmt.Println(json.StringifyJson(r))
}

func TestAbLib_Dungeon(t *testing.T) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	rlib, err := NewRarityLib(client, &testSender{})
	assert.NoError(t, err)

	ablib, err := NewAbLib(rlib)
	assert.NoError(t, err)

	r, err := ablib.Dungeon("0x6D81145629f154dBf07fDAb18D2892818626FeCF", []uint64{3574968, 3575125})

	assert.NoError(t, err)

	fmt.Println(json.StringifyJson(r))
}