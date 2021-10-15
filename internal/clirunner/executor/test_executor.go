package executor

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type TestExecutor struct {
	limit          int
	clib           lib
	RunLevelupTask bool
	RunGoldClaimTask bool
	RunAdventureTask bool
	RunDungeonTask   bool
}

func NewTestExecutor(clib lib) *TestExecutor {
	return &TestExecutor{limit: 100, clib: clib}
}

func (l *TestExecutor) Run(tasks []*types.Task) error {
	for i := range tasks {
		task := tasks[i]

		switch task.TaskType {
		case types.LevelupTask:

			if !l.RunLevelupTask {
				continue
			}

			err := l.limitRun(task.IDs, func(ids []uint64) error {
				return l.clib.Levelup(ids)
			})

			if err != nil {
				log.Logger.Error("limit run Levelup failed")
				return err
			}

		case types.GoldClaimTask:

			if !l.RunGoldClaimTask {
				continue
			}

			err := l.limitRun(task.IDs, func(ids []uint64) error {
				return l.clib.ClaimGold(ids)
			})

			if err != nil {
				log.Logger.Error("limit run ClaimGold failed")
				return err
			}

		case types.AdventureTask:

			if !l.RunAdventureTask {
				continue
			}

			err := l.limitRun(task.IDs, func(ids []uint64) error {
				return l.clib.Adventure(ids)
			})

			if err != nil {
				log.Logger.Error("limit run Adventure failed")
				return err
			}

		case types.DungeonTask:

			if !l.RunDungeonTask {
				continue
			}

			err := l.limitRun(task.IDs, func(ids []uint64) error {
				return l.clib.Dungeon(ids)
			})

			if err != nil {
				log.Logger.Error("limit run Dungeon failed")
				return err
			}
		}
	}

	return nil
}

func (l *TestExecutor) limitRun(ids []uint64, exec func(ids []uint64) error) error {
	if len(ids) < l.limit || len(ids) == l.limit {
		if err := exec(ids); err != nil {
			log.Logger.Error("limit run failed", zap.Uint64s("ids", ids), zap.Error(err))
			return err
		}
	}

	var tmp []uint64

	for i := range ids {
		if len(tmp) > l.limit {
			if err := exec(tmp); err != nil {
				log.Logger.Error("limit run failed", zap.Uint64s("tmp ids", ids), zap.Error(err))
				return err
			}
			tmp = []uint64{}
		}

		tmp = append(tmp, ids[i])
	}

	if len(tmp) > 0 {
		if err := exec(tmp); err != nil {
			log.Logger.Error("limit run failed", zap.Uint64s("tmp ids", ids), zap.Error(err))
			return err
		}
	}

	return nil
}
