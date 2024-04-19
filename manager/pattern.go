package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

// PatternAInBigRoad patternA 如果遇到莊就打閒，遇到閒就打莊
func (analyzeManager *AnalyzeManager) PatternAInBigRoad(bigRoad *roadmap.BigRoad) {
	if len(bigRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(bigRoad.Columns); colIndex++ {
		for blockIndex := 0; blockIndex < len(bigRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				bigRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				bigRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := bigRoad.Columns[len(bigRoad.Columns)-1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManager.Predictions.BigRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.BigRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.BigRoad.BetArea = 1
		}
	}
}

func (analyzeManager *AnalyzeManager) PatternAInBigEyeRoad(bigEyeRoad *roadmap.BigEyeRoad) {
	if len(bigEyeRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(bigEyeRoad.Columns); colIndex++ {
		for blockIndex := 0; blockIndex < len(bigEyeRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				bigEyeRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				bigEyeRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := bigEyeRoad.Columns[len(bigEyeRoad.Columns)-1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManager.Predictions.BigEyeRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.BigEyeRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.BigEyeRoad.BetArea = 1
		}
	}
}

func (analyzeManager *AnalyzeManager) PatternAInSmallRoad(smallRoad *roadmap.SmallRoad) {
	if len(smallRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(smallRoad.Columns); colIndex++ {
		for blockIndex := 0; blockIndex < len(smallRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				smallRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				smallRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := smallRoad.Columns[len(smallRoad.Columns)-1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManager.Predictions.SmallRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.SmallRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.SmallRoad.BetArea = 1
		}
	}
}

func (analyzeManager *AnalyzeManager) PatternAInCockroachRoad(cockroachRoad *roadmap.CockroachRoad) {
	if len(cockroachRoad.Columns) == 0 {
		return
	}
	for colIndex := 0; colIndex < len(cockroachRoad.Columns); colIndex++ {
		for blockIndex := 0; blockIndex < len(cockroachRoad.Columns[colIndex].Blocks); blockIndex++ {
			if blockIndex > 1 {
				cockroachRoad.Columns[colIndex].Blocks[blockIndex].Result = -1
			}
			if colIndex > 0 && blockIndex == 0 {
				cockroachRoad.Columns[colIndex].Blocks[blockIndex].Result = 1
			}
		}
	}
	lastColumn := cockroachRoad.Columns[len(cockroachRoad.Columns)-1]
	if len(lastColumn.Blocks) > 1 {
		analyzeManager.Predictions.CockroachRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.CockroachRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.CockroachRoad.BetArea = 1
		}
	}
}

// PatternBInBigRoad patternB 從第2排開始打,打反邊(開藍打紅,開紅打藍)
func (analyzeManager *AnalyzeManager) PatternBInBigRoad(bigRoad *roadmap.BigRoad) {
	if len(bigRoad.Columns) <= 1 {
		return
	}
	if len(bigRoad.Columns[0].Blocks) >= 2 && len(bigRoad.Columns[1].Blocks) >= 2 {
		bigRoad.Columns[1].Blocks[1].Result = -1
	}
	for colIndex := 2; colIndex < len(bigRoad.Columns); colIndex++ {
		// 前兩行 block 大於等於 2 並且前一行 block 等於 1 的時候，第一顆 Result 等於 1
		if len(bigRoad.Columns[colIndex-2].Blocks) >= 2 && len(bigRoad.Columns[colIndex-1].Blocks) == 1 {
			bigRoad.Columns[colIndex].Blocks[0].Result = 1
		}
		if len(bigRoad.Columns[colIndex-1].Blocks) >= 2 && len(bigRoad.Columns[colIndex].Blocks) >= 2 {
			bigRoad.Columns[colIndex].Blocks[1].Result = -1
		}
	}
	lastColumn := bigRoad.Columns[len(bigRoad.Columns)-1]
	prevColumn := bigRoad.Columns[len(bigRoad.Columns)-2]
	if len(lastColumn.Blocks) == 1 && len(prevColumn.Blocks) > 1 {
		analyzeManager.Predictions.BigRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.BigRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.BigRoad.BetArea = 1
		}
	}
}

func (analyzeManager *AnalyzeManager) PatternBInBigEyeRoad(bigEyeRoad *roadmap.BigEyeRoad) {
	if len(bigEyeRoad.Columns) <= 1 {
		return
	}
	if len(bigEyeRoad.Columns[0].Blocks) >= 2 && len(bigEyeRoad.Columns[1].Blocks) >= 2 {
		bigEyeRoad.Columns[1].Blocks[1].Result = -1
	}
	for colIndex := 2; colIndex < len(bigEyeRoad.Columns); colIndex++ {
		// 前兩行 block 大於等於 2 並且前一行 block 等於 1 的時候，第一顆 Result 等於 1
		if len(bigEyeRoad.Columns[colIndex-2].Blocks) >= 2 && len(bigEyeRoad.Columns[colIndex-1].Blocks) == 1 {
			bigEyeRoad.Columns[colIndex].Blocks[0].Result = 1
		}
		if len(bigEyeRoad.Columns[colIndex-1].Blocks) >= 2 && len(bigEyeRoad.Columns[colIndex].Blocks) >= 2 {
			bigEyeRoad.Columns[colIndex].Blocks[1].Result = -1
		}
	}
	lastColumn := bigEyeRoad.Columns[len(bigEyeRoad.Columns)-1]
	prevColumn := bigEyeRoad.Columns[len(bigEyeRoad.Columns)-2]
	if len(lastColumn.Blocks) == 1 && len(prevColumn.Blocks) > 1 {
		analyzeManager.Predictions.BigEyeRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.BigEyeRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.BigEyeRoad.BetArea = 1
		}
	}
}

func (analyzeManager *AnalyzeManager) PatternBInSmallRoad(smallRoad *roadmap.SmallRoad) {
	if len(smallRoad.Columns) <= 1 {
		return
	}
	if len(smallRoad.Columns[0].Blocks) >= 2 && len(smallRoad.Columns[1].Blocks) >= 2 {
		smallRoad.Columns[1].Blocks[1].Result = -1
	}
	for colIndex := 2; colIndex < len(smallRoad.Columns); colIndex++ {
		// 前兩行 block 大於等於 2 並且前一行 block 等於 1 的時候，第一顆 Result 等於 1
		if len(smallRoad.Columns[colIndex-2].Blocks) >= 2 && len(smallRoad.Columns[colIndex-1].Blocks) == 1 {
			smallRoad.Columns[colIndex].Blocks[0].Result = 1
		}
		if len(smallRoad.Columns[colIndex-1].Blocks) >= 2 && len(smallRoad.Columns[colIndex].Blocks) >= 2 {
			smallRoad.Columns[colIndex].Blocks[1].Result = -1
		}
	}
	lastColumn := smallRoad.Columns[len(smallRoad.Columns)-1]
	prevColumn := smallRoad.Columns[len(smallRoad.Columns)-2]
	if len(lastColumn.Blocks) == 1 && len(prevColumn.Blocks) > 1 {
		analyzeManager.Predictions.SmallRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.SmallRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.SmallRoad.BetArea = 1
		}

	}
}

func (analyzeManager *AnalyzeManager) PatternBInCockroachRoad(cockroachRoad *roadmap.CockroachRoad) {
	if len(cockroachRoad.Columns) <= 1 {
		return
	}
	if len(cockroachRoad.Columns[0].Blocks) >= 2 && len(cockroachRoad.Columns[1].Blocks) >= 2 {
		cockroachRoad.Columns[1].Blocks[1].Result = -1
	}
	for colIndex := 2; colIndex < len(cockroachRoad.Columns); colIndex++ {
		// 前兩行 block 大於等於 2 並且前一行 block 等於 1 的時候，第一顆 Result 等於 1
		if len(cockroachRoad.Columns[colIndex-2].Blocks) >= 2 && len(cockroachRoad.Columns[colIndex-1].Blocks) == 1 {
			cockroachRoad.Columns[colIndex].Blocks[0].Result = 1
		}
		if len(cockroachRoad.Columns[colIndex-1].Blocks) >= 2 && len(cockroachRoad.Columns[colIndex].Blocks) >= 2 {
			cockroachRoad.Columns[colIndex].Blocks[1].Result = -1
		}
	}
	lastColumn := cockroachRoad.Columns[len(cockroachRoad.Columns)-1]
	prevColumn := cockroachRoad.Columns[len(cockroachRoad.Columns)-2]
	if len(lastColumn.Blocks) == 1 && len(prevColumn.Blocks) > 1 {
		analyzeManager.Predictions.CockroachRoad.Bet = 1
		if lastColumn.Blocks[0].Symbol == roadmap.Symbol_Banker {
			analyzeManager.Predictions.CockroachRoad.BetArea = 2
		} else {
			analyzeManager.Predictions.CockroachRoad.BetArea = 1
		}
	}
}
