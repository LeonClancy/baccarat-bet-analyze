package manager

import "github.com/LeonClancy/baccarat-bet-analyze/roadmap"

func (r *RoadmapManager) AskRoad(symbol roadmap.Symbol) roadmap.AskRoadResult {
	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 2 {
		return roadmap.AskRoadResult{
			BigEyeRoadNext:    &roadmap.Block{Symbol: roadmap.Symbol_BlockDefault},
			SmallRoadNext:     &roadmap.Block{Symbol: roadmap.Symbol_BlockDefault},
			CockroachRoadNext: &roadmap.Block{Symbol: roadmap.Symbol_BlockDefault},
		}
	}

	var copyRoadmapManager = NewRoadmapManager("temp")
	copyBigRoad := copyRoadmapManager.Roadmaps.BigRoad
	copyTotalRoad := copyRoadmapManager.Roadmaps.TotalRoad
	copyBigRoad.Columns = make([]*roadmap.Column, len(r.Roadmaps.BigRoad.Columns))
	copyTotalRoad.Columns = make([]*roadmap.Column, len(r.Roadmaps.TotalRoad.Columns))
	for i := range r.Roadmaps.BigRoad.Columns {
		copyBigRoad.Columns[i] = &roadmap.Column{Blocks: []*roadmap.Block{}}
		copyTotalRoad.Columns[i] = &roadmap.Column{Blocks: []*roadmap.Block{}}
		copyBigRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.Roadmaps.BigRoad.Columns[i].Blocks))
		copyTotalRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.Roadmaps.TotalRoad.Columns[i].Blocks))
		copy(copyBigRoad.Columns[i].Blocks, r.Roadmaps.BigRoad.Columns[i].Blocks)
		copy(copyTotalRoad.Columns[i].Blocks, r.Roadmaps.TotalRoad.Columns[i].Blocks)
	}

	copyRoadmapManager.drawBigRoad(symbol)

	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]
	bigRoadLatestBlock := bigRoadLatestColumn.Blocks[len(bigRoadLatestColumn.Blocks)-1]

	if bigRoadLatestBlock.Symbol == roadmap.Symbol_OnlyResult ||
		bigRoadLatestBlock.Symbol == roadmap.Symbol_OnlyResultAndNewLine {
		copyRoadmapManager.restoreBigRoad()
	}

	bigEyeRoadNext := copyRoadmapManager.askBigEyeRoad(symbol)
	smallRoadNext := copyRoadmapManager.askSmallEyeRoad(symbol)
	cockroachRoadNext := copyRoadmapManager.askCockroachRoad(symbol)

	return roadmap.AskRoadResult{
		BigEyeRoadNext:    &roadmap.Block{Symbol: bigEyeRoadNext},
		SmallRoadNext:     &roadmap.Block{Symbol: smallRoadNext},
		CockroachRoadNext: &roadmap.Block{Symbol: cockroachRoadNext},
	}
}

func (r *RoadmapManager) askBigEyeRoad(symbol roadmap.Symbol) roadmap.Symbol {
	if symbol == roadmap.Symbol_Tie ||
		symbol == roadmap.Symbol_TieAndBankerPair ||
		symbol == roadmap.Symbol_TieAndPlayerPair ||
		symbol == roadmap.Symbol_TieAndBothPair {
		return roadmap.Symbol_BlockDefault
	}

	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 2 {
		return roadmap.Symbol_BlockDefault
	}

	if len(bigRoad.Columns[1].Blocks) == 1 && len(bigRoad.Columns) == 2 {
		return roadmap.Symbol_BlockDefault
	}

	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	// 換列
	if len(bigRoadLatestColumn.Blocks) == 1 {
		// 比對前一列和前二列的結果位置是齊整，則於大眼路畫紅圈。
		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
			len(bigRoad.Columns[len(bigRoad.Columns)-3].Blocks) {
			return roadmap.Symbol_Banker
		} else {
			// 以大路最新的結果，比對前一列與前二列的結果位置是不齊整，則於大眼路畫藍圈
			return roadmap.Symbol_Player
		}
	} else {
		// 向下
		diff := len(bigRoadLatestColumn.Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks)
		if diff == 1 {
			// 以大路最新的結果，水平方向跟前一列作對比，如前一列無結果，則於大眼路畫藍圈
			return roadmap.Symbol_Player
		}
		// 以大路最新的結果，水平方向跟前一列作對比，如前一列有結果，則於大眼路畫紅圈
		// 以大路最新的結果，水平方向跟前一列作對比，如前一列的前二行或以上都無結果，則於大眼路畫紅圈
		return roadmap.Symbol_Banker
	}
}

func (r *RoadmapManager) askSmallEyeRoad(symbol roadmap.Symbol) roadmap.Symbol {
	if symbol == roadmap.Symbol_Tie ||
		symbol == roadmap.Symbol_TieAndBankerPair ||
		symbol == roadmap.Symbol_TieAndPlayerPair ||
		symbol == roadmap.Symbol_TieAndBothPair {
		return roadmap.Symbol_BlockDefault
	}

	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 3 {
		return roadmap.Symbol_BlockDefault
	}

	if len(bigRoad.Columns[2].Blocks) == 1 && len(bigRoad.Columns) == 3 {
		return roadmap.Symbol_BlockDefault
	}

	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	if len(bigRoadLatestColumn.Blocks) == 1 {
		// 換列
		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
			len(bigRoad.Columns[len(bigRoad.Columns)-4].Blocks) {
			// 以大路最新的結果，比對前一列與前三列結果位置是齊整，則於小路畫紅點。
			return roadmap.Symbol_Banker
		} else {
			// 以大路最新的結果，對比前一列與前三列的位置是不齊整，則於小路畫藍點。
			return roadmap.Symbol_Player
		}
	} else {
		// 向下
		diff := len(bigRoad.Columns[len(bigRoad.Columns)-1].Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-3].Blocks)
		if diff == 1 {
			// 以大路最新的結果，對比前一列與前三列的位置是不齊整，則於小路畫藍點。
			return roadmap.Symbol_Player
		} else {
			// 以大路最新的結果，水平方向跟前二列作對比，前二列有結果時，則於小路畫紅點。
			// 以大路最新的結果，水平方向跟前二列作對比，如前二列的前二行或以上都無結果，則於小路畫紅點。
			return roadmap.Symbol_Banker
		}
	}

}

func (r *RoadmapManager) askCockroachRoad(block roadmap.Symbol) roadmap.Symbol {
	if block == roadmap.Symbol_Tie ||
		block == roadmap.Symbol_TieAndBankerPair ||
		block == roadmap.Symbol_TieAndPlayerPair ||
		block == roadmap.Symbol_TieAndBothPair {
		return roadmap.Symbol_BlockDefault
	}

	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 4 {
		return roadmap.Symbol_BlockDefault
	}

	if len(bigRoad.Columns[3].Blocks) == 1 && len(bigRoad.Columns) == 4 {
		return roadmap.Symbol_BlockDefault
	}

	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	if len(bigRoadLatestColumn.Blocks) == 1 {
		// 換列 對比前一列與前四列結果位置
		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
			len(bigRoad.Columns[len(bigRoad.Columns)-5].Blocks) {
			// 以大路最新的結果，對比前一列與前四列結果位置是齊整，則於小強路畫紅色斜線。
			return roadmap.Symbol_Banker
		} else {
			// 以大路最新的結果，對比前一列與前四列結果位置是不齊整，則於大眼路畫藍色斜線。
			return roadmap.Symbol_Player
		}
	} else {
		// 向下
		diff := len(bigRoad.Columns[len(bigRoad.Columns)-1].Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-4].Blocks)
		if diff == 1 {
			// 以大路最新的結果，水平方向跟前三列作對比，如前三列無結果，則於小強路畫藍色斜線。
			return roadmap.Symbol_Player
		} else {
			// 以大路最新的結果，水平方向跟前三列作對比，如前三列的前二行或以上都無結果，則於小強路畫紅色斜線。
			// 以大路最新的結果，水平方向跟前三列作對比，前三列有結果時，則於小強路畫紅色斜線。
			return roadmap.Symbol_Banker
		}
	}
}
