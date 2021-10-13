package clirunner

import (
	"github.com/Xib1uvXi/rarity-box/pkg/box"
	"github.com/Xib1uvXi/rarity-box/pkg/box/summonerinfo"
	"github.com/Xib1uvXi/rarity-box/pkg/clirunner/executor"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib"
	"github.com/Xib1uvXi/rarity-box/pkg/tasker"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

func Build(pkStr string, hasFee bool) (*Runner, error) {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	if err != nil {
		return nil, err
	}

	clib, err := contractlib.BuildV1CallerLib(pkStr, client, hasFee)

	if err != nil {
		return nil, err
	}

	limitExecutor := executor.NewLimitExecutor(clib)

	summonersSyncer := box.BuildBox(summonerinfo.NewConcurrentReader(40, clib))

	runner := NewRunner(tasker.NewTaskBuilder(summonersSyncer), limitExecutor, strings.ToLower(clib.TxSender.Address().String()))

	return runner, nil
}
