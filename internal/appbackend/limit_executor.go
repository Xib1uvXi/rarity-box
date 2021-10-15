package appbackend

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
)

type TxBuilder interface {
	Adventure(from string, ids []uint64) (*types.RawTxParam, error)
	Levelup(from string, ids []uint64) (*types.RawTxParam, error)
	ClaimGold(from string, ids []uint64) (*types.RawTxParam, error)
	Dungeon(from string, ids []uint64) (*types.RawTxParam, error)
}

type LimitExecutor struct {
	txBuilder TxBuilder
}

func NewLimitExecutor(clib TxBuilder) *LimitExecutor {
	return &LimitExecutor{txBuilder: clib}
}

func (l *LimitExecutor) Run(from, taskType string, tasks []*types.Task) ([]*types.RawTxParam, error) {
	for i := range tasks {
		task := tasks[i]

		switch task.TaskType {
		case types.LevelupTask:
			if task.TaskType != taskType {
				continue
			}

			params, err := l.limitRun(249, from, task.IDs, func(from string, ids []uint64) (*types.RawTxParam, error) {
				return l.txBuilder.Levelup(from, ids)
			})

			if err != nil {
				log.Logger.Error("limit run Levelup failed")
				return nil, err
			}

			return params, nil

		case types.GoldClaimTask:
			if task.TaskType != taskType {
				continue
			}

			params, err := l.limitRun(100, from, task.IDs, func(from string, ids []uint64) (*types.RawTxParam, error) {
				return l.txBuilder.ClaimGold(from, ids)
			})

			if err != nil {
				log.Logger.Error("limit run ClaimGold failed")
				return nil, err
			}

			return params, nil

		case types.AdventureTask:
			if task.TaskType != taskType {
				continue
			}

			params, err := l.limitRun(249, from, task.IDs, func(from string, ids []uint64) (*types.RawTxParam, error) {
				return l.txBuilder.Adventure(from, ids)
			})

			if err != nil {
				log.Logger.Error("limit run Adventure failed")
				return nil, err
			}

			return params, nil

		case types.DungeonTask:
			if task.TaskType != taskType {
				continue
			}

			params, err := l.limitRun(100, from, task.IDs, func(from string, ids []uint64) (*types.RawTxParam, error) {
				return l.txBuilder.Dungeon(from, ids)
			})

			if err != nil {
				log.Logger.Error("limit run Dungeon failed")
				return nil, err
			}

			return params, nil
		}
	}

	return nil, nil
}

func (l *LimitExecutor) limitRun(limit int, from string, ids []uint64, exec func(from string, ids []uint64) (*types.RawTxParam, error)) ([]*types.RawTxParam, error) {
	var params []*types.RawTxParam

	if len(ids) < limit || len(ids) == limit {
		param, err := exec(from, ids)
		if err != nil {
			return nil, err
		}

		params = append(params, param)
		return params, nil
	}

	var tmp []uint64

	for i := range ids {
		if len(tmp) > limit {

			param, err := exec(from, tmp)
			if err != nil {
				return nil, err
			}

			params = append(params, param)
			tmp = []uint64{}
		}

		tmp = append(tmp, ids[i])
	}

	if len(tmp) > 0 {
		param, err := exec(from, tmp)
		if err != nil {
			return nil, err
		}

		params = append(params, param)
	}

	return params, nil
}
