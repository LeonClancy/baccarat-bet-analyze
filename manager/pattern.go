package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

// pattrenA 如果遇到莊就打閒，遇到閒就打莊
func (analyzeManger *AnalyzeManager) PatternAInBigRoad(bigRoad *roadmap.BigRoad) {
	if len(bigRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(bigRoad.Columns) ; colIndex++ {
		for blockIndex := 0; blockIndex < len(bigRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				bigRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				bigRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := bigRoad.Columns[len(bigRoad.Columns) - 1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManger.Predictions.BigRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManger.Predictions.BigRoad.BetArea = 2
		} else {
			analyzeManger.Predictions.BigRoad.BetArea = 1
		}
	}
}

func (analyzeManger *AnalyzeManager) PatternAInBigEyeRoad(bigEyeRoad *roadmap.BigEyeRoad) {
	if len(bigEyeRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(bigEyeRoad.Columns) ; colIndex++ {
		for blockIndex := 0; blockIndex < len(bigEyeRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				bigEyeRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				bigEyeRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := bigEyeRoad.Columns[len(bigEyeRoad.Columns) - 1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManger.Predictions.BigEyeRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManger.Predictions.BigEyeRoad.BetArea = 2
		} else {
			analyzeManger.Predictions.BigEyeRoad.BetArea = 1
		}
	}
}

func (analyzeManger *AnalyzeManager) PatternAInSmallRoad(smallRoad *roadmap.SmallRoad) {
	if len(smallRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(smallRoad.Columns) ; colIndex++ {
		for blockIndex := 0; blockIndex < len(smallRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				smallRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				smallRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := smallRoad.Columns[len(smallRoad.Columns) - 1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManger.Predictions.SmallRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManger.Predictions.SmallRoad.BetArea = 2
		} else {
			analyzeManger.Predictions.SmallRoad.BetArea = 1
		}
	}
}

func (analyzeManger *AnalyzeManager) PatternAInCockroachRoad(cockroachRoad *roadmap.CockroachRoad) {
	if len(cockroachRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(cockroachRoad.Columns) ; colIndex++ {
		for blockIndex := 0; blockIndex < len(cockroachRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				cockroachRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				cockroachRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := cockroachRoad.Columns[len(cockroachRoad.Columns) - 1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManger.Predictions.CockroachRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManger.Predictions.CockroachRoad.BetArea = 2
		} else {
			analyzeManger.Predictions.CockroachRoad.BetArea = 1
		}
	}
}

func (*AnalyzeManager) PatternBInBigRoad(bigRoad *roadmap.BigRoad) {

}

func (*AnalyzeManager) PatternBInBigEyeRoad(bigEyeRoad *roadmap.BigEyeRoad) {

}

func (*AnalyzeManager) PatternBInSmallRoad(smallRoad *roadmap.SmallRoad) {

}

func (*AnalyzeManager) PatternBInCockroachRoad(cockroachRoad *roadmap.CockroachRoad) {

}
