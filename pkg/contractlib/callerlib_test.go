package contractlib

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewCallerLib(t *testing.T) {
	pk := os.Getenv("FTMPK")
	if pk == "" {
		t.Skip("pk is empty")
	}

	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	clib, err := BuildV1CallerLib(pk, client, false)
	assert.NoError(t, err)

	assert.NoError(t, clib.Adventure([]uint64{2062009, 2063971, 2222628}))
}

func TestNewCallerLib_Dungeon(t *testing.T) {
	pk := os.Getenv("FTMPK")
	if pk == "" {
		t.Skip("pk is empty")
	}

	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	clib, err := BuildV1CallerLib(pk, client, false)
	assert.NoError(t, err)

	assert.NoError(t, clib.Dungeon([]uint64{2062009, 2222628, 2841525}))
}
