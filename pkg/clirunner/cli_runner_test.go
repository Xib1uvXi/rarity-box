package clirunner

import (
	"fmt"
	"github.com/Xib1uvXi/rarity-box/pkg/box"
	"github.com/Xib1uvXi/rarity-box/pkg/box/summonerinfo"
	"github.com/Xib1uvXi/rarity-box/pkg/clirunner/executor"
	"github.com/Xib1uvXi/rarity-box/pkg/common/json"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib"
	"github.com/Xib1uvXi/rarity-box/pkg/tasker"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestRunner_Run(t *testing.T) {
	pk := os.Getenv("FTMPK")
	if pk == "" {
		t.Skip("pk is empty")
	}

	client, err := ethclient.Dial("https://rpc.ftm.tools")
	assert.NoError(t, err)

	clib, err := contractlib.BuildV1CallerLib(pk, client, false)
	assert.NoError(t, err)

	summonersSyncer := box.BuildBox(summonerinfo.NewConcurrentReader(40, clib))

	tExecutor := executor.NewTestExecutor(clib)
	tExecutor.RunGoldClaimTask = true

	runner := NewRunner(tasker.NewTaskBuilder(summonersSyncer), tExecutor, strings.ToLower(clib.TxSender.Address().String()))

	assert.NoError(t, runner.Run())
}

type testExecutor struct {
}

func (t *testExecutor) Run(tasks []*types.Task) error {
	fmt.Println(json.StringifyJson(tasks))
	return nil
}
