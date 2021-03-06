package clirunner

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/tasker"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"go.uber.org/zap"
)

type Executor interface {
	Run([]*types.Task) error
}

type Runner struct {
	taskBuilder *tasker.TaskBuilder
	executor    Executor
	address     string
}

func NewRunner(taskBuilder *tasker.TaskBuilder, executor Executor, address string) *Runner {
	return &Runner{taskBuilder: taskBuilder, executor: executor, address: address}
}

func (r *Runner) Run() error {
	tasks, err := r.taskBuilder.Tasks(r.address)
	if err != nil {
		log.Logger.Error("runner build task failed", zap.String("address", r.address), zap.Error(err))
		return err
	}

	err = r.executor.Run(tasks)
	if err != nil {
		return err
	}

	return nil
}
