package taskfactory

import (
	"github.com/Xib1uvXi/rarity-box/pkg/tasker/summonerstatuschecker"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
)

type Checker interface {
	Status(summoner *types.Summoner) *types.SummonerStatus
}

type callerTaskFactory struct {
	checker Checker
}

func NewCallerTaskFactory() *callerTaskFactory {
	return &callerTaskFactory{checker: &summonerstatuschecker.Checkers{}}
}

func (c *callerTaskFactory) Task(summoners []*types.Summoner) []*types.Task {
	var (
		adventureTask = &types.Task{TaskType: types.AdventureTask, IDs: []uint64{}}
		levelupTask   = &types.Task{TaskType: types.LevelupTask, IDs: []uint64{}}
		goldclaimTask = &types.Task{TaskType: types.GoldClaimTask, IDs: []uint64{}}
		dungeonTask   = &types.Task{TaskType: types.DungeonTask, IDs: []uint64{}}
	)

	for i := range summoners {
		summoner := summoners[i]

		status := c.checker.Status(summoner)

		if status.CanAdventure {
			adventureTask.IDs = append(adventureTask.IDs, summoner.TokenID)
		}

		if status.CanLevelUp {
			levelupTask.IDs = append(levelupTask.IDs, summoner.TokenID)
		}

		if status.GoldClaimable {
			goldclaimTask.IDs = append(goldclaimTask.IDs, summoner.TokenID)
		}

		if status.CanDungeonAdventure {
			dungeonTask.IDs = append(dungeonTask.IDs, summoner.TokenID)
		}
	}

	adventureTask.Count = len(adventureTask.IDs)
	levelupTask.Count = len(levelupTask.IDs)
	goldclaimTask.Count = len(goldclaimTask.IDs)
	dungeonTask.Count = len(dungeonTask.IDs)

	var result []*types.Task

	result = append(result, levelupTask, goldclaimTask, adventureTask, dungeonTask)

	return result
}
