package main

import (
	"github.com/Xib1uvXi/rarity-box/pkg/appbackend"
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

	summonersSyncer := box.BuildBox(summonerinfo.NewConcurrentReader(40, rlib))

	taskBuilder := tasker.NewTaskBuilder(summonersSyncer)

	ginServer := appbackend.NewGinServer(appbackend.NewSrv(summonersSyncer, taskBuilder))

	router := gin.Default()

	router.GET("/summoners/:address", ginServer.Summoners)
	router.GET("/tasks/:address", ginServer.Tasks)

	router.Run(":8080")
}
