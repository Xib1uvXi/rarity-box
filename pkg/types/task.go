package types

const (
	AdventureTask = "adventure"
	LevelupTask   = "levelup"
	GoldClaimTask = "goldclaim"
	DungeonTask   = "dungeon"
)

type Task struct {
	TaskType string   `json:"task_type"`
	IDs      []uint64 `json:"ids"`
	Count    int      `json:"count"`
}
