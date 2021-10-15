package main

import (
	appbackend2 "github.com/Xib1uvXi/rarity-box/internal/appbackend"
	"github.com/Xib1uvXi/rarity-box/pkg/box"
	"github.com/Xib1uvXi/rarity-box/pkg/box/summonerinfo"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib"
	"github.com/Xib1uvXi/rarity-box/pkg/tasker"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := ethclient.Dial("https://rpc.ftm.tools")
	if err != nil {
		panic(err)
	}

	rlib, err := contractlib.BuildV1RarityLibWithoutSender(client)

	if err != nil {
		panic(err)
	}

	ablib, err := contractlib.NewAbLib(rlib)

	if err != nil {
		panic(err)
	}

	summonersSyncer := box.BuildBox(summonerinfo.NewConcurrentReader(40, ablib))

	taskBuilder := tasker.NewTaskBuilder(summonersSyncer)

	executor := appbackend2.NewLimitExecutor(ablib)

	ginServer := appbackend2.NewGinServer(appbackend2.NewSrv(summonersSyncer, taskBuilder, ablib, executor))

	router := gin.Default()

	router.GET("/summoners/:address", ginServer.Summoners)
	router.GET("/tasks/:address", ginServer.Tasks)
	router.POST("/approve/check", ginServer.IsOperator)
	router.GET("/tx/approve", ginServer.SetOperator)
	router.POST("/tx/task/run", ginServer.RunTask)

	router.Run(":8080")
}
