package roadmap

type AnalyzeLevel struct {
	Levels       []int32
	Boom         bool
	LevelCounter int
}

type LevelManager struct {
	Level AnalyzeLevel
}

func NewLevelManager() *LevelManager {
	return &LevelManager{
		Level: AnalyzeLevel{
			Levels:       []int32{3, 2, 1},
			Boom:         false,
			LevelCounter: 0,
		},
	}
}

func (LevelManager *LevelManager) ResetLevel() {
	LevelManager.Level.LevelCounter = 0
}

func (LevelManager *LevelManager) IncreaseLevel() {
	LevelManager.Level.LevelCounter++
}

func (LevelManager *LevelManager) GetLevel() int32 {
	if LevelManager.Level.LevelCounter >= len(LevelManager.Level.Levels) {
		return 0
	}
	return LevelManager.Level.Levels[LevelManager.Level.LevelCounter]
}
