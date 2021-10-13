package contractlib

import (
	"context"
	"fmt"
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
