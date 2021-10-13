package tasker

import (
	"github.com/Xib1uvXi/rarity-box/pkg/tasker/taskfactory"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"strings"
)

type SummonersSyncer interface {
	SyncSummoners(address string) ([]*types.Summoner, error)
}

type TaskFactory interface {
	Task([]*types.Summoner) []*types.Task
}

type TaskBuilder struct {
	summonersSyncer SummonersSyncer
	taskFactory     TaskFactory
}

func NewTaskBuilder(summonersSyncer SummonersSyncer) *TaskBuilder {
	return &TaskBuilder{summonersSyncer: summonersSyncer, taskFactory: taskfactory.NewCallerTaskFactory()}
}

func (t *TaskBuilder) Tasks(address string) ([]*types.Task, error) {
	summoners, err := t.summonersSyncer.SyncSummoners(strings.ToLower(address))

	if err != nil {
		return nil, err
	}

	return t.taskFactory.Task(summoners), nil
}
