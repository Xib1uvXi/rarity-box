package executor

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type lib interface {
	Adventure(ids []uint64) error
	Levelup(ids []uint64) error
	ClaimGold(ids []uint64) error
	Dungeon(ids []uint64) error
}

type LimitExecutor struct {
	clib lib
}

func NewLimitExecutor(clib lib) *LimitExecutor {
	return &LimitExecutor{clib: clib}
}

func (l *LimitExecutor) Run(tasks []*types.Task) error {
	for i := range tasks {
		task := tasks[i]

		if len(task.IDs) == 0 {
			continue
		}

		switch task.TaskType {
		case types.LevelupTask:
			err := l.limitRun(249, task.IDs, func(ids []uint64) error {
				return l.clib.Levelup(ids)
			})

			if err != nil {
				log.Logger.Error("limit run Levelup failed")
				return err
			}

		case types.GoldClaimTask:
			err := l.limitRun(100, task.IDs, func(ids []uint64) error {
				return l.clib.ClaimGold(ids)
			})

			if err != nil {
				log.Logger.Error("limit run ClaimGold failed")
				return err
			}

		case types.AdventureTask:
			err := l.limitRun(249, task.IDs, func(ids []uint64) error {
				return l.clib.Adventure(ids)
			})

			if err != nil {
				log.Logger.Error("limit run Adventure failed")
				return err
			}

		case types.DungeonTask:
			err := l.limitRun(100, task.IDs, func(ids []uint64) error {
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

func (l *LimitExecutor) limitRun(limit int, ids []uint64, exec func(ids []uint64) error) error {
	if len(ids) < limit || len(ids) == limit {
		if err := exec(ids); err != nil {
			log.Logger.Error("limit run failed", zap.Uint64s("ids", ids), zap.Error(err))
			return err
		}
	}

	var tmp []uint64

	for i := range ids {
		if len(tmp) > limit {
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
