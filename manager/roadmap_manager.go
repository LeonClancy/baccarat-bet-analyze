package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/dealer"
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
	"github.com/gogf/gf/container/garray"
	"sync"
)

type RoadmapManager struct {
	mutex         sync.RWMutex
	Roadmaps      *roadmap.Roadmap `json:"roadmaps"`
	previousMaps  *roadmap.Roadmap
	Name          string `json:"name"`
	ResultCounter *roadmap.RoadmapsResultCount
}

func NewRoadmapManager(name string) *RoadmapManager {
	return &RoadmapManager{
		Name:  name,
		mutex: sync.RWMutex{},
		Roadmaps: &roadmap.Roadmap{
			BeadPlate: &roadmap.BeadPlate{
				Blocks: []*roadmap.Block{},
			},
			BigRoad: &roadmap.BigRoad{
				Columns: []*roadmap.Column{},
			},
			BigEyeRoad: &roadmap.BigEyeRoad{
				Columns: []*roadmap.Column{
					&roadmap.Column{
						Blocks: []*roadmap.Block{},
					},
				},
			},
			SmallRoad: &roadmap.SmallRoad{
				Columns: []*roadmap.Column{
					&roadmap.Column{
						Blocks: []*roadmap.Block{},
					},
				},
			},
			CockroachRoad: &roadmap.CockroachRoad{
				Columns: []*roadmap.Column{
					&roadmap.Column{
						Blocks: []*roadmap.Block{},
					},
				},
			},
		},
		previousMaps: &roadmap.Roadmap{
			BeadPlate: &roadmap.BeadPlate{
				Blocks: []*roadmap.Block{},
			},
			BigRoad: &roadmap.BigRoad{
				Columns: []*roadmap.Column{},
			},
			BigEyeRoad: &roadmap.BigEyeRoad{
				Columns: []*roadmap.Column{
					&roadmap.Column{
						Blocks: []*roadmap.Block{},
					},
				},
			},
			SmallRoad: &roadmap.SmallRoad{
				Columns: []*roadmap.Column{
					&roadmap.Column{
						Blocks: []*roadmap.Block{},
					},
				},
			},
			CockroachRoad: &roadmap.CockroachRoad{
				Columns: []*roadmap.Column{
					&roadmap.Column{
						Blocks: []*roadmap.Block{},
					},
				},
			},
		},
		ResultCounter: &roadmap.RoadmapsResultCount{
			BigRoadCounts: &roadmap.Result{
				TieCount:    0,
				PlayerCount: 0,
				BankerCount: 0,
			},
			BigEyeRoadCounts: &roadmap.Result{
				TieCount:    0,
				PlayerCount: 0,
				BankerCount: 0,
			},
			SmallRoadCounts: &roadmap.Result{
				TieCount:    0,
				PlayerCount: 0,
				BankerCount: 0,
			},
			CockroachRoadCounts: &roadmap.Result{
				TieCount:    0,
				PlayerCount: 0,
				BankerCount: 0,
			},
		},
	}
}

func (r *RoadmapManager) Draw(results []dealer.Result) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	symbol := r.convertResultsToSymbol(results)

	r.previousMaps.BeadPlate.Blocks = make([]*roadmap.Block, len(r.Roadmaps.BeadPlate.Blocks))
	copy(r.previousMaps.BeadPlate.Blocks, r.Roadmaps.BeadPlate.Blocks)

	for i := range r.Roadmaps.BigRoad.Columns {
		if len(r.previousMaps.BigRoad.Columns)-1 < i {
			r.previousMaps.BigRoad.Columns = append(r.previousMaps.BigRoad.Columns, &roadmap.Column{})
		}
		r.previousMaps.BigRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.Roadmaps.BigRoad.Columns[i].Blocks))
		copy(r.previousMaps.BigRoad.Columns[i].Blocks, r.Roadmaps.BigRoad.Columns[i].Blocks)
	}

	for i := range r.Roadmaps.BigEyeRoad.Columns {
		if len(r.previousMaps.BigEyeRoad.Columns)-1 < i {
			r.previousMaps.BigEyeRoad.Columns = append(r.previousMaps.BigEyeRoad.Columns, &roadmap.Column{})
		}
		r.previousMaps.BigEyeRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.Roadmaps.BigEyeRoad.Columns[i].Blocks))
		copy(r.previousMaps.BigEyeRoad.Columns[i].Blocks, r.Roadmaps.BigEyeRoad.Columns[i].Blocks)
	}

	for i := range r.Roadmaps.SmallRoad.Columns {
		if len(r.previousMaps.SmallRoad.Columns)-1 < i {
			r.previousMaps.SmallRoad.Columns = append(r.previousMaps.SmallRoad.Columns, &roadmap.Column{})
		}
		r.previousMaps.SmallRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.Roadmaps.SmallRoad.Columns[i].Blocks))
		copy(r.previousMaps.SmallRoad.Columns[i].Blocks, r.Roadmaps.SmallRoad.Columns[i].Blocks)
	}

	for i := range r.Roadmaps.CockroachRoad.Columns {
		if len(r.previousMaps.CockroachRoad.Columns)-1 < i {
			r.previousMaps.CockroachRoad.Columns = append(r.previousMaps.CockroachRoad.Columns, &roadmap.Column{})
		}
		r.previousMaps.CockroachRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.Roadmaps.CockroachRoad.Columns[i].Blocks))
		copy(r.previousMaps.CockroachRoad.Columns[i].Blocks, r.Roadmaps.CockroachRoad.Columns[i].Blocks)
	}

	r.drawBeadPlate(symbol)
	r.drawBigRoad(symbol)
	r.drawBigEyeRoad(symbol)
	r.drawSmallEyeRoad(symbol)
	r.drawCockroachRoad(symbol)
}

func (r *RoadmapManager) Print() *roadmap.Roadmap {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.Roadmaps
}

func (r *RoadmapManager) Reset() {
	r.Roadmaps = &roadmap.Roadmap{
		BeadPlate: &roadmap.BeadPlate{
			Blocks: []*roadmap.Block{},
		},
		BigRoad: &roadmap.BigRoad{
			Columns: []*roadmap.Column{},
		},
		BigEyeRoad: &roadmap.BigEyeRoad{
			Columns: []*roadmap.Column{
				&roadmap.Column{
					Blocks: []*roadmap.Block{},
				},
			},
		},
		SmallRoad: &roadmap.SmallRoad{
			Columns: []*roadmap.Column{
				&roadmap.Column{
					Blocks: []*roadmap.Block{},
				},
			},
		},
		CockroachRoad: &roadmap.CockroachRoad{
			Columns: []*roadmap.Column{
				&roadmap.Column{
					Blocks: []*roadmap.Block{},
				},
			},
		},
	}
}

func (r *RoadmapManager) drawBeadPlate(symbol roadmap.Symbol) {
	plate := r.Roadmaps.BeadPlate

	//if len(plate.Blocks) == 60 {
	//	plate.Blocks = []roadmap.Block{}
	//}

	if len(plate.Blocks) == 0 {
		plate.Blocks = append(plate.Blocks, &roadmap.Block{
			Symbol:   symbol,
			TieCount: 0,
		})
		return
	}

	plate.Blocks = append(plate.Blocks, &roadmap.Block{
		Symbol:   symbol,
		TieCount: 0,
	})
}

func (r *RoadmapManager) convertResultsToSymbol(results []dealer.Result) roadmap.Symbol {
	var CopyResult garray.Array

	for i := range results {
		CopyResult.Append(results[i])
	}

	if CopyResult.Contains(dealer.Result_BankerPair) &&
		CopyResult.Contains(dealer.Result_PlayerPair) {
		if CopyResult.Contains(dealer.Result_Banker) {
			return roadmap.Symbol_BankerAndBothPair
		}
		if CopyResult.Contains(dealer.Result_Player) {
			return roadmap.Symbol_PlayerAndBothPair
		}
		if CopyResult.Contains(dealer.Result_Tie) {
			return roadmap.Symbol_TieAndBothPair
		}
	}

	if CopyResult.Contains(dealer.Result_BankerPair) {
		if CopyResult.Contains(dealer.Result_Banker) {
			return roadmap.Symbol_BankerAndBankerPair
		}
		if CopyResult.Contains(dealer.Result_Player) {
			return roadmap.Symbol_PlayerAndBankerPair
		}
		if CopyResult.Contains(dealer.Result_Tie) {
			return roadmap.Symbol_TieAndBankerPair
		}
	}

	if CopyResult.Contains(dealer.Result_PlayerPair) {
		if CopyResult.Contains(dealer.Result_Banker) {
			return roadmap.Symbol_BankerAndPlayerPair
		}
		if CopyResult.Contains(dealer.Result_Player) {
			return roadmap.Symbol_PlayerAndPlayerPair
		}
		if CopyResult.Contains(dealer.Result_Tie) {
			return roadmap.Symbol_TieAndPlayerPair
		}
	}

	if CopyResult.Contains(dealer.Result_Banker) {
		return roadmap.Symbol_Banker
	}
	if CopyResult.Contains(dealer.Result_Player) {
		return roadmap.Symbol_Player
	}
	if CopyResult.Contains(dealer.Result_Tie) {
		return roadmap.Symbol_Tie
	}

	return roadmap.Symbol_BlockDefault
}

func (r *RoadmapManager) drawBigRoad(symbol roadmap.Symbol) {
	bigRoad := r.Roadmaps.BigRoad

	//if len(bigRoad.Columns) == 50 {
	//	bigRoad.Columns = []*roadmap.Column{}
	//}

	if len(bigRoad.Columns) == 0 {
		bigRoad.Columns = append(bigRoad.Columns,
			&roadmap.Column{
				Blocks: []*roadmap.Block{},
			})
		if symbol == roadmap.Symbol_Tie ||
			symbol == roadmap.Symbol_TieAndPlayerPair ||
			symbol == roadmap.Symbol_TieAndBankerPair ||
			symbol == roadmap.Symbol_TieAndBothPair {
			r.ResultCounter.BigRoadCounts.TieCount++
			bigRoad.Columns[0].Blocks = append(bigRoad.Columns[0].Blocks, &roadmap.Block{
				Symbol:   symbol,
				TieCount: 1,
			})
			return
		}
		bigRoad.Columns[0].Blocks = append(bigRoad.Columns[0].Blocks, &roadmap.Block{
			Symbol:   symbol,
			TieCount: 0,
		})
		return
	}

	lastColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	if symbol == roadmap.Symbol_Tie ||
		symbol == roadmap.Symbol_TieAndPlayerPair ||
		symbol == roadmap.Symbol_TieAndBankerPair ||
		symbol == roadmap.Symbol_TieAndBothPair {

		r.ResultCounter.BigRoadCounts.TieCount++

		lastBlock := lastColumn.Blocks[len(lastColumn.Blocks)-1]
		lastBlock.TieCount++

		if lastBlock.Symbol == roadmap.Symbol_Tie ||
			lastBlock.Symbol == roadmap.Symbol_BankerAndTie ||
			lastBlock.Symbol == roadmap.Symbol_BankerAndBankerPairAndTie ||
			lastBlock.Symbol == roadmap.Symbol_BankerAndPlayerPairAndTie ||
			lastBlock.Symbol == roadmap.Symbol_BankerAndBothPair ||
			lastBlock.Symbol == roadmap.Symbol_PlayerAndTie ||
			lastBlock.Symbol == roadmap.Symbol_PlayerAndBankerPairAndTie ||
			lastBlock.Symbol == roadmap.Symbol_PlayerAndPlayerPairAndTie ||
			lastBlock.Symbol == roadmap.Symbol_PlayerAndBothPairAndTie {
			return
		}
		switch lastBlock.Symbol {
		case roadmap.Symbol_Banker:
			lastBlock.Symbol = roadmap.Symbol_BankerAndTie
		case roadmap.Symbol_BankerAndBankerPair:
			lastBlock.Symbol = roadmap.Symbol_BankerAndBankerPairAndTie
		case roadmap.Symbol_BankerAndPlayerPair:
			lastBlock.Symbol = roadmap.Symbol_BankerAndPlayerPairAndTie
		case roadmap.Symbol_BankerAndBothPair:
			lastBlock.Symbol = roadmap.Symbol_BankerAndBothPairAndTie
		case roadmap.Symbol_Player:
			lastBlock.Symbol = roadmap.Symbol_PlayerAndTie
		case roadmap.Symbol_PlayerAndBankerPair:
			lastBlock.Symbol = roadmap.Symbol_PlayerAndBankerPairAndTie
		case roadmap.Symbol_PlayerAndPlayerPair:
			lastBlock.Symbol = roadmap.Symbol_PlayerAndPlayerPairAndTie
		case roadmap.Symbol_PlayerAndBothPair:
			lastBlock.Symbol = roadmap.Symbol_PlayerAndBothPairAndTie
		}
		return
	}

	lastColumnFirstBlock := lastColumn.Blocks[0]
	lastColumnLastBlock := lastColumn.Blocks[len(lastColumn.Blocks)-1]

	if symbol == roadmap.Symbol_Banker ||
		symbol == roadmap.Symbol_BankerAndBankerPair ||
		symbol == roadmap.Symbol_BankerAndPlayerPair ||
		symbol == roadmap.Symbol_BankerAndBothPair {
		r.ResultCounter.BigRoadCounts.BankerCount++
		if lastColumnLastBlock.Symbol == roadmap.Symbol_Tie ||
			lastColumnLastBlock.Symbol == roadmap.Symbol_TieAndBankerPair ||
			lastColumnLastBlock.Symbol == roadmap.Symbol_TieAndPlayerPair ||
			lastColumnLastBlock.Symbol == roadmap.Symbol_TieAndBothPair {
			switch symbol {
			case roadmap.Symbol_Banker:
				lastColumnLastBlock.Symbol = roadmap.Symbol_BankerAndTie
			case roadmap.Symbol_BankerAndBankerPair:
				lastColumnLastBlock.Symbol = roadmap.Symbol_BankerAndBankerPairAndTie
			case roadmap.Symbol_BankerAndPlayerPair:
				lastColumnLastBlock.Symbol = roadmap.Symbol_BankerAndPlayerPairAndTie
			case roadmap.Symbol_BankerAndBothPair:
				lastColumnLastBlock.Symbol = roadmap.Symbol_BankerAndBothPairAndTie
			}
			return
		}

		if lastColumnFirstBlock.Symbol == roadmap.Symbol_Banker ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndBankerPair ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndPlayerPair ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndBothPair ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndTie ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndBankerPairAndTie ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndPlayerPairAndTie ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_BankerAndBothPairAndTie {
			lastColumn.Blocks = append(lastColumn.Blocks, &roadmap.Block{
				Symbol:   symbol,
				TieCount: 0,
			})
			return
		} else {
			bigRoad.Columns = append(bigRoad.Columns, &roadmap.Column{
				Blocks: []*roadmap.Block{
					{
						Symbol:   symbol,
						TieCount: 0,
					},
				},
			})
			return
		}
	}

	if symbol == roadmap.Symbol_Player ||
		symbol == roadmap.Symbol_PlayerAndBankerPair ||
		symbol == roadmap.Symbol_PlayerAndPlayerPair ||
		symbol == roadmap.Symbol_PlayerAndBothPair {
		r.ResultCounter.BigRoadCounts.PlayerCount++
		if lastColumnLastBlock.Symbol == roadmap.Symbol_Tie ||
			lastColumnLastBlock.Symbol == roadmap.Symbol_TieAndBankerPair ||
			lastColumnLastBlock.Symbol == roadmap.Symbol_TieAndPlayerPair ||
			lastColumnLastBlock.Symbol == roadmap.Symbol_TieAndBothPair {
			switch symbol {
			case roadmap.Symbol_Player:
				lastColumnLastBlock.Symbol = roadmap.Symbol_PlayerAndTie
			case roadmap.Symbol_PlayerAndBankerPair:
				lastColumnLastBlock.Symbol = roadmap.Symbol_PlayerAndBankerPairAndTie
			case roadmap.Symbol_PlayerAndPlayerPair:
				lastColumnLastBlock.Symbol = roadmap.Symbol_PlayerAndPlayerPairAndTie
			case roadmap.Symbol_PlayerAndBothPair:
				lastColumnLastBlock.Symbol = roadmap.Symbol_PlayerAndBothPairAndTie
			}
			return
		}
		if lastColumnFirstBlock.Symbol == roadmap.Symbol_Player ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndBankerPair ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndPlayerPair ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndBothPair ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndTie ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndBankerPairAndTie ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndPlayerPairAndTie ||
			lastColumnFirstBlock.Symbol == roadmap.Symbol_PlayerAndBothPairAndTie {
			lastColumn.Blocks = append(lastColumn.Blocks, &roadmap.Block{
				Symbol:   symbol,
				TieCount: 0,
			})
			return
		} else {
			bigRoad.Columns = append(bigRoad.Columns, &roadmap.Column{
				Blocks: []*roadmap.Block{
					{
						Symbol:   symbol,
						TieCount: 0,
					},
				},
			})
			return
		}
	}
}

func (r *RoadmapManager) drawBigEyeRoad(block roadmap.Symbol) {
	if block == roadmap.Symbol_Tie ||
		block == roadmap.Symbol_TieAndBankerPair ||
		block == roadmap.Symbol_TieAndPlayerPair ||
		block == roadmap.Symbol_TieAndBothPair {
		return
	}

	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 2 {
		return
	}

	if len(bigRoad.Columns[1].Blocks) == 1 && len(bigRoad.Columns) == 2 {
		return
	}

	bigEyeRoad := r.Roadmaps.BigEyeRoad
	bigEyeRoadLatestColumn := bigEyeRoad.Columns[0]
	if len(bigEyeRoad.Columns) > 0 {
		bigEyeRoadLatestColumn = bigEyeRoad.Columns[len(bigEyeRoad.Columns)-1]
	}
	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	// 換列
	if len(bigRoadLatestColumn.Blocks) == 1 {
		// 比對前一列和前二列的結果位置是齊整，則於大眼路畫紅圈。
		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
			len(bigRoad.Columns[len(bigRoad.Columns)-3].Blocks) {
			r.ResultCounter.BigEyeRoadCounts.BankerCount++
			r.bigEyeRoadNewBlock(bigEyeRoadLatestColumn, bigEyeRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Banker,
				TieCount: 0,
			})
			return
		} else {
			// 以大路最新的結果，比對前一列與前二列的結果位置是不齊整，則於大眼路畫藍圈
			r.ResultCounter.BigEyeRoadCounts.PlayerCount++
			r.bigEyeRoadNewBlock(bigEyeRoadLatestColumn, bigEyeRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Player,
				TieCount: 0,
			})
			return
		}
	} else {
		// 向下
		diff := len(bigRoadLatestColumn.Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks)
		if diff == 1 {
			// 以大路最新的結果，水平方向跟前一列作對比，如前一列無結果，則於大眼路畫藍圈
			r.ResultCounter.BigEyeRoadCounts.PlayerCount++
			r.bigEyeRoadNewBlock(bigEyeRoadLatestColumn, bigEyeRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Player,
				TieCount: 0,
			})
			return
		}
		// 以大路最新的結果，水平方向跟前一列作對比，如前一列有結果，則於大眼路畫紅圈
		// 以大路最新的結果，水平方向跟前一列作對比，如前一列的前二行或以上都無結果，則於大眼路畫紅圈
		r.ResultCounter.BigEyeRoadCounts.BankerCount++
		r.bigEyeRoadNewBlock(bigEyeRoadLatestColumn, bigEyeRoad, &roadmap.Block{
			Symbol:   roadmap.Symbol_Banker,
			TieCount: 0,
		})
		return
	}
}

func (r *RoadmapManager) drawSmallEyeRoad(block roadmap.Symbol) {
	if block == roadmap.Symbol_Tie ||
		block == roadmap.Symbol_TieAndBankerPair ||
		block == roadmap.Symbol_TieAndPlayerPair ||
		block == roadmap.Symbol_TieAndBothPair {
		return
	}

	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 3 {
		return
	}

	if len(bigRoad.Columns[2].Blocks) == 1 && len(bigRoad.Columns) == 3 {
		return
	}

	smallRoad := r.Roadmaps.SmallRoad
	smallRoadLatestColumn := smallRoad.Columns[0]
	if len(smallRoad.Columns) > 0 {
		smallRoadLatestColumn = smallRoad.Columns[len(smallRoad.Columns)-1]
	}
	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	if len(bigRoadLatestColumn.Blocks) == 1 {
		// 換列
		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
			len(bigRoad.Columns[len(bigRoad.Columns)-4].Blocks) {
			// 以大路最新的結果，比對前一列與前三列結果位置是齊整，則於小路畫紅點。
			r.ResultCounter.SmallRoadCounts.BankerCount++
			r.smallRoadNewBlock(smallRoadLatestColumn, smallRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Banker,
				TieCount: 0,
			})
			return
		} else {
			// 以大路最新的結果，對比前一列與前三列的位置是不齊整，則於小路畫藍點。
			r.ResultCounter.SmallRoadCounts.PlayerCount++
			r.smallRoadNewBlock(smallRoadLatestColumn, smallRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Player,
				TieCount: 0,
			})
			return
		}
	} else {
		// 向下
		diff := len(bigRoad.Columns[len(bigRoad.Columns)-1].Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-3].Blocks)
		if diff == 1 {
			// 以大路最新的結果，對比前一列與前三列的位置是不齊整，則於小路畫藍點。
			r.ResultCounter.SmallRoadCounts.PlayerCount++
			r.smallRoadNewBlock(smallRoadLatestColumn, smallRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Player,
				TieCount: 0,
			})
			return
		} else {
			// 以大路最新的結果，水平方向跟前二列作對比，前二列有結果時，則於小路畫紅點。
			// 以大路最新的結果，水平方向跟前二列作對比，如前二列的前二行或以上都無結果，則於小路畫紅點。
			r.ResultCounter.SmallRoadCounts.BankerCount++
			r.smallRoadNewBlock(smallRoadLatestColumn, smallRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Banker,
				TieCount: 0,
			})
			return
		}
	}
}

func (r *RoadmapManager) drawCockroachRoad(block roadmap.Symbol) {
	if block == roadmap.Symbol_Tie ||
		block == roadmap.Symbol_TieAndBankerPair ||
		block == roadmap.Symbol_TieAndPlayerPair ||
		block == roadmap.Symbol_TieAndBothPair {
		return
	}

	bigRoad := r.Roadmaps.BigRoad
	if len(bigRoad.Columns) < 4 {
		return
	}

	if len(bigRoad.Columns[3].Blocks) == 1 && len(bigRoad.Columns) == 4 {
		return
	}

	cockroachRoad := r.Roadmaps.CockroachRoad
	cockroachRoadLatestColumn := cockroachRoad.Columns[0]
	if len(cockroachRoad.Columns) > 0 {
		cockroachRoadLatestColumn = cockroachRoad.Columns[len(cockroachRoad.Columns)-1]
	}
	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]

	if len(bigRoadLatestColumn.Blocks) == 1 {
		// 換列 對比前一列與前四列結果位置
		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
			len(bigRoad.Columns[len(bigRoad.Columns)-5].Blocks) {
			// 以大路最新的結果，對比前一列與前四列結果位置是齊整，則於小強路畫紅色斜線。
			r.ResultCounter.CockroachRoadCounts.BankerCount++
			r.cockroachRoadNewBlock(cockroachRoadLatestColumn, cockroachRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Banker,
				TieCount: 0,
			})
			return
		} else {
			// 以大路最新的結果，對比前一列與前四列結果位置是不齊整，則於大眼路畫藍色斜線。
			r.ResultCounter.CockroachRoadCounts.PlayerCount++
			r.cockroachRoadNewBlock(cockroachRoadLatestColumn, cockroachRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Player,
				TieCount: 0,
			})
			return
		}
	} else {
		// 向下
		diff := len(bigRoad.Columns[len(bigRoad.Columns)-1].Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-4].Blocks)
		if diff == 1 {
			// 以大路最新的結果，水平方向跟前三列作對比，如前三列無結果，則於小強路畫藍色斜線。
			r.ResultCounter.CockroachRoadCounts.PlayerCount++
			r.cockroachRoadNewBlock(cockroachRoadLatestColumn, cockroachRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Player,
				TieCount: 0,
			})
			return
		} else {
			// 以大路最新的結果，水平方向跟前三列作對比，如前三列的前二行或以上都無結果，則於小強路畫紅色斜線。
			// 以大路最新的結果，水平方向跟前三列作對比，前三列有結果時，則於小強路畫紅色斜線。
			r.ResultCounter.CockroachRoadCounts.BankerCount++
			r.cockroachRoadNewBlock(cockroachRoadLatestColumn, cockroachRoad, &roadmap.Block{
				Symbol:   roadmap.Symbol_Banker,
				TieCount: 0,
			})
			return
		}
	}
}

func (r *RoadmapManager) bigEyeRoadNewBlock(latestColumn *roadmap.Column, road *roadmap.BigEyeRoad, block *roadmap.Block) {
	if len(latestColumn.Blocks) == 0 {
		latestColumn.Blocks = append(latestColumn.Blocks, block)
		return
	}
	if latestColumn.Blocks[0] == block {
		latestColumn.Blocks = append(latestColumn.Blocks, block)
	} else {
		road.Columns = append(road.Columns,
			&roadmap.Column{
				Blocks: []*roadmap.Block{block},
			},
		)
	}
}

func (r *RoadmapManager) smallRoadNewBlock(latestColumn *roadmap.Column, road *roadmap.SmallRoad, block *roadmap.Block) {
	if len(latestColumn.Blocks) == 0 {
		latestColumn.Blocks = append(latestColumn.Blocks, block)
		return
	}
	if latestColumn.Blocks[0] == block {
		latestColumn.Blocks = append(latestColumn.Blocks, block)
	} else {
		road.Columns = append(road.Columns,
			&roadmap.Column{
				Blocks: []*roadmap.Block{block},
			},
		)
	}
}

func (r *RoadmapManager) cockroachRoadNewBlock(latestColumn *roadmap.Column, road *roadmap.CockroachRoad, block *roadmap.Block) {
	if len(latestColumn.Blocks) == 0 {
		latestColumn.Blocks = append(latestColumn.Blocks, block)
		return
	}
	if latestColumn.Blocks[0] == block {
		latestColumn.Blocks = append(latestColumn.Blocks, block)
	} else {
		road.Columns = append(road.Columns,
			&roadmap.Column{
				Blocks: []*roadmap.Block{block},
			},
		)
	}
}

//
//func (r *RoadmapManager) AskRoad(block roadmap.Block) *roadmap.AskRoadRecall {
//	r.mutex.RLock()
//	defer r.mutex.RUnlock()
//
//	recall := &roadmap.AskRoadRecall{
//		Header: &foundation.Header{
//			Uri: route.URI_AskRoadRecall,
//		},
//		BigEyeRoadNext:    nil,
//		SmallRoadNext:     nil,
//		CockroachRoadNext: nil,
//		AskRoadCall: &roadmap.AskRoadCall{
//			roadmap.Block: &block,
//		},
//	}
//
//	if len(r.roadmaps.BigRoad.Columns) < 2 {
//		return recall
//	}
//
//	var copyRoadmapManager = NewRoadmapManager()
//	copyBigroad := copyRoadmapManager.roadmaps.BigRoad
//	copyBigroad.Columns = make([]*roadmap.Column, len(r.roadmaps.BigRoad.Columns))
//	for i := range r.roadmaps.BigRoad.Columns {
//		copyBigroad.Columns[i] = &roadmap.Column{Blocks: []*roadmap.Block{}}
//		copyBigroad.Columns[i].Blocks = make([]*roadmap.Block, len(r.roadmaps.BigRoad.Columns[i].Blocks))
//		copy(copyBigroad.Columns[i].Blocks, r.roadmaps.BigRoad.Columns[i].Blocks)
//	}
//
//	copyRoadmapManager.drawBigRoad(block.Symbol)
//
//	recall.BigEyeRoadNext = &roadmap.Block{
//		Symbol: copyRoadmapManager.askBigEyeRoad(block.Symbol),
//		TieCount:       0,
//	}
//	recall.SmallRoadNext = &roadmap.Block{
//		Symbol: copyRoadmapManager.askSmallEyeRoad(block.Symbol),
//		TieCount:       0,
//	}
//	recall.CockroachRoadNext = &roadmap.Block{
//		Symbol: copyRoadmapManager.askCockroachRoad(block.Symbol),
//		TieCount:       0,
//	}
//
//	return recall
//}
//
//func (r *RoadmapManager) askBigEyeRoad(symbol roadmap.Symbol) roadmap.Symbol {
//	if symbol == roadmap.Symbol_Tie ||
//		symbol == roadmap.Symbol_TieAndBankerPair ||
//		symbol == roadmap.Symbol_TieAndPlayerPair ||
//		symbol == roadmap.Symbol_TieAndBothPair {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	bigRoad := r.roadmaps.BigRoad
//	if len(bigRoad.Columns) < 2 {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	if len(bigRoad.Columns[1].Blocks) == 1 && len(bigRoad.Columns) == 2 {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]
//
//	// 換列
//	if len(bigRoadLatestColumn.Blocks) == 1 {
//		// 比對前一列和前二列的結果位置是齊整，則於大眼路畫紅圈。
//		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
//			len(bigRoad.Columns[len(bigRoad.Columns)-3].Blocks) {
//			return roadmap.Symbol_Banker
//		} else {
//			// 以大路最新的結果，比對前一列與前二列的結果位置是不齊整，則於大眼路畫藍圈
//			return roadmap.Symbol_Player
//		}
//	} else {
//		// 向下
//		diff := len(bigRoadLatestColumn.Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks)
//		if diff == 1 {
//			// 以大路最新的結果，水平方向跟前一列作對比，如前一列無結果，則於大眼路畫藍圈
//			return roadmap.Symbol_Player
//		}
//		// 以大路最新的結果，水平方向跟前一列作對比，如前一列有結果，則於大眼路畫紅圈
//		// 以大路最新的結果，水平方向跟前一列作對比，如前一列的前二行或以上都無結果，則於大眼路畫紅圈
//		return roadmap.Symbol_Banker
//	}
//}
//
//func (r *RoadmapManager) askSmallEyeRoad(symbol roadmap.Symbol) roadmap.Symbol {
//	if symbol == roadmap.Symbol_Tie ||
//		symbol == roadmap.Symbol_TieAndBankerPair ||
//		symbol == roadmap.Symbol_TieAndPlayerPair ||
//		symbol == roadmap.Symbol_TieAndBothPair {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	bigRoad := r.roadmaps.BigRoad
//	if len(bigRoad.Columns) < 3 {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	if len(bigRoad.Columns[2].Blocks) == 1 && len(bigRoad.Columns) == 3 {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]
//
//	if len(bigRoadLatestColumn.Blocks) == 1 {
//		// 換列
//		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
//			len(bigRoad.Columns[len(bigRoad.Columns)-4].Blocks) {
//			// 以大路最新的結果，比對前一列與前三列結果位置是齊整，則於小路畫紅點。
//			return roadmap.Symbol_Banker
//		} else {
//			// 以大路最新的結果，對比前一列與前三列的位置是不齊整，則於小路畫藍點。
//			return roadmap.Symbol_Player
//		}
//	} else {
//		// 向下
//		diff := len(bigRoad.Columns[len(bigRoad.Columns)-1].Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-3].Blocks)
//		if diff == 1 {
//			// 以大路最新的結果，對比前一列與前三列的位置是不齊整，則於小路畫藍點。
//			return roadmap.Symbol_Player
//		} else {
//			// 以大路最新的結果，水平方向跟前二列作對比，前二列有結果時，則於小路畫紅點。
//			// 以大路最新的結果，水平方向跟前二列作對比，如前二列的前二行或以上都無結果，則於小路畫紅點。
//			return roadmap.Symbol_Banker
//		}
//	}
//
//}
//
//func (r *RoadmapManager) askCockroachRoad(block roadmap.Symbol) roadmap.Symbol {
//	if block == roadmap.Symbol_Tie ||
//		block == roadmap.Symbol_TieAndBankerPair ||
//		block == roadmap.Symbol_TieAndPlayerPair ||
//		block == roadmap.Symbol_TieAndBothPair {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	bigRoad := r.roadmaps.BigRoad
//	if len(bigRoad.Columns) < 4 {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	if len(bigRoad.Columns[3].Blocks) == 1 && len(bigRoad.Columns) == 4 {
//		return roadmap.Symbol_BlockDefault
//	}
//
//	bigRoadLatestColumn := bigRoad.Columns[len(bigRoad.Columns)-1]
//
//	if len(bigRoadLatestColumn.Blocks) == 1 {
//		// 換列 對比前一列與前四列結果位置
//		if len(bigRoad.Columns[len(bigRoad.Columns)-2].Blocks) ==
//			len(bigRoad.Columns[len(bigRoad.Columns)-5].Blocks) {
//			// 以大路最新的結果，對比前一列與前四列結果位置是齊整，則於小強路畫紅色斜線。
//			return roadmap.Symbol_Banker
//		} else {
//			// 以大路最新的結果，對比前一列與前四列結果位置是不齊整，則於大眼路畫藍色斜線。
//			return roadmap.Symbol_Player
//		}
//	} else {
//		// 向下
//		diff := len(bigRoad.Columns[len(bigRoad.Columns)-1].Blocks) - len(bigRoad.Columns[len(bigRoad.Columns)-4].Blocks)
//		if diff == 1 {
//			// 以大路最新的結果，水平方向跟前三列作對比，如前三列無結果，則於小強路畫藍色斜線。
//			return roadmap.Symbol_Player
//		} else {
//			// 以大路最新的結果，水平方向跟前三列作對比，如前三列的前二行或以上都無結果，則於小強路畫紅色斜線。
//			// 以大路最新的結果，水平方向跟前三列作對比，前三列有結果時，則於小強路畫紅色斜線。
//			return roadmap.Symbol_Banker
//		}
//	}
//}
//
//func (r *RoadmapManager) AskRoads() []*roadmap.AskRoadRecall {
//	blocks := []roadmap.Block{{
//		Symbol: roadmap.Symbol_Banker,
//	}, {
//		Symbol: roadmap.Symbol_Player,
//	}}
//	var recalls []*roadmap.AskRoadRecall
//	for i := range blocks {
//		recalls = append(recalls, r.AskRoad(blocks[i]))
//	}
//	return recalls
//}

func (r *RoadmapManager) Restore() {
	r.Roadmaps = &roadmap.Roadmap{
		BeadPlate: &roadmap.BeadPlate{
			Blocks: []*roadmap.Block{},
		},
		BigRoad: &roadmap.BigRoad{
			Columns: []*roadmap.Column{},
		},
		BigEyeRoad: &roadmap.BigEyeRoad{
			Columns: []*roadmap.Column{
				&roadmap.Column{
					Blocks: []*roadmap.Block{},
				},
			},
		},
		SmallRoad: &roadmap.SmallRoad{
			Columns: []*roadmap.Column{
				&roadmap.Column{
					Blocks: []*roadmap.Block{},
				},
			},
		},
		CockroachRoad: &roadmap.CockroachRoad{
			Columns: []*roadmap.Column{
				&roadmap.Column{
					Blocks: []*roadmap.Block{},
				},
			},
		},
	}

	r.Roadmaps.BeadPlate.Blocks = make([]*roadmap.Block, len(r.previousMaps.BeadPlate.Blocks))
	copy(r.Roadmaps.BeadPlate.Blocks, r.previousMaps.BeadPlate.Blocks)

	for i := range r.previousMaps.BigRoad.Columns {
		if len(r.Roadmaps.BigRoad.Columns)-1 < i {
			r.Roadmaps.BigRoad.Columns = append(r.Roadmaps.BigRoad.Columns, &roadmap.Column{})
		}
		r.Roadmaps.BigRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.previousMaps.BigRoad.Columns[i].Blocks))
		copy(r.Roadmaps.BigRoad.Columns[i].Blocks, r.previousMaps.BigRoad.Columns[i].Blocks)
	}

	for i := range r.previousMaps.BigEyeRoad.Columns {
		if len(r.Roadmaps.BigEyeRoad.Columns)-1 < i {
			r.Roadmaps.BigEyeRoad.Columns = append(r.Roadmaps.BigEyeRoad.Columns, &roadmap.Column{})
		}
		r.Roadmaps.BigEyeRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.previousMaps.BigEyeRoad.Columns[i].Blocks))
		copy(r.Roadmaps.BigEyeRoad.Columns[i].Blocks, r.previousMaps.BigEyeRoad.Columns[i].Blocks)
	}

	for i := range r.previousMaps.SmallRoad.Columns {
		if len(r.Roadmaps.SmallRoad.Columns)-1 < i {
			r.Roadmaps.SmallRoad.Columns = append(r.Roadmaps.SmallRoad.Columns, &roadmap.Column{})
		}
		r.Roadmaps.SmallRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.previousMaps.SmallRoad.Columns[i].Blocks))
		copy(r.Roadmaps.SmallRoad.Columns[i].Blocks, r.previousMaps.SmallRoad.Columns[i].Blocks)
	}

	for i := range r.previousMaps.CockroachRoad.Columns {
		if len(r.Roadmaps.CockroachRoad.Columns)-1 < i {
			r.Roadmaps.CockroachRoad.Columns = append(r.Roadmaps.CockroachRoad.Columns, &roadmap.Column{})
		}
		r.Roadmaps.CockroachRoad.Columns[i].Blocks = make([]*roadmap.Block, len(r.previousMaps.CockroachRoad.Columns[i].Blocks))
		copy(r.Roadmaps.CockroachRoad.Columns[i].Blocks, r.previousMaps.CockroachRoad.Columns[i].Blocks)
	}

}
