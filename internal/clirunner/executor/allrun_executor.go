package executor

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
)

type AllRunExecutor struct {
	clib             lib
	RunLevelupTask   bool
	RunGoldClaimTask bool
	RunAdventureTask bool
	RunDungeonTask   bool
}

func NewAllRunExecutor(clib lib) *AllRunExecutor {
	return &AllRunExecutor{clib: clib}
}

func (l *AllRunExecutor) Run(tasks []*types.Task) error {
	for i := range tasks {
		task := tasks[i]

		switch task.TaskType {
		case types.LevelupTask:

			if !l.RunLevelupTask {
				continue
			}

			err := l.clib.Levelup(task.IDs)

			if err != nil {
				log.Logger.Error("limit run Levelup failed")
				return err
			}

		case types.GoldClaimTask:

			if !l.RunGoldClaimTask {
				continue
			}

			err := l.clib.ClaimGold(task.IDs)

			if err != nil {
				log.Logger.Error("limit run ClaimGold failed")
				return err
			}

		case types.AdventureTask:

			if !l.RunAdventureTask {
				continue
			}

			err := l.clib.Adventure(task.IDs)

			if err != nil {
				log.Logger.Error("limit run Adventure failed")
				return err
			}

		case types.DungeonTask:

			if !l.RunDungeonTask {
				continue
			}

			err := l.clib.Dungeon(task.IDs)

			if err != nil {
				log.Logger.Error("limit run Dungeon failed")
				return err
			}
		}
	}

	return nil
}
