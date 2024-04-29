package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

var Patterns = map[int]string{
	0: "無",
	1: "A",
	2: "B",
}

type AnalyzeManager struct {
	Pattern1    int
	Pattern2    int
	Predictions *Predictions `json:"predictions"`
}

type Predictions struct {
	BigRoad       *Prediction `json:"bigRoad"`
	BigEyeRoad    *Prediction `json:"bigEyeRoad"`
	SmallRoad     *Prediction `json:"smallRoad"`
	CockroachRoad *Prediction `json:"cockroachRoad"`
	TotalRoad     *Prediction `json:"totalRoad"`
}

type Prediction struct {
	Bet     int `json:"bet"`
	BetArea int `json:"betArea"`
}

func NewAnalyzeManager() *AnalyzeManager {
	return &AnalyzeManager{
		Pattern1: 0,
		Pattern2: 0,
		Predictions: &Predictions{
			TotalRoad: &Prediction{
				Bet:     0,
				BetArea: 0,
			},
			BigRoad: &Prediction{
				Bet:     0,
				BetArea: 0,
			},
			BigEyeRoad: &Prediction{
				Bet:     0,
				BetArea: 0,
			},
			SmallRoad: &Prediction{
				Bet:     0,
				BetArea: 0,
			},
			CockroachRoad: &Prediction{
				Bet:     0,
				BetArea: 0,
			},
		},
	}
}

func (analyzeManager *AnalyzeManager) Analyze(roadmap *roadmap.Roadmap) *roadmap.Roadmap {
	for _, c := range roadmap.BigRoad.Columns {
		c.Result = 0
		for _, b := range c.Blocks {
			b.Result = 0
		}
	}

	for _, c := range roadmap.BigEyeRoad.Columns {
		c.Result = 0
		for _, b := range c.Blocks {
			b.Result = 0
		}
	}

	for _, c := range roadmap.SmallRoad.Columns {
		c.Result = 0
		for _, b := range c.Blocks {
			b.Result = 0
		}
	}

	for _, c := range roadmap.CockroachRoad.Columns {
		c.Result = 0
		for _, b := range c.Blocks {
			b.Result = 0
		}
	}

	analyzeManager.Predictions.TotalRoad.Bet = 0
	analyzeManager.Predictions.TotalRoad.BetArea = 0
	analyzeManager.Predictions.BigEyeRoad.Bet = 0
	analyzeManager.Predictions.BigEyeRoad.BetArea = 0
	analyzeManager.Predictions.BigRoad.Bet = 0
	analyzeManager.Predictions.BigRoad.BetArea = 0
	analyzeManager.Predictions.SmallRoad.Bet = 0
	analyzeManager.Predictions.SmallRoad.BetArea = 0
	analyzeManager.Predictions.CockroachRoad.Bet = 0
	analyzeManager.Predictions.CockroachRoad.BetArea = 0

	analyzeManager.AnalyzeWithPattern(roadmap, analyzeManager.Pattern1)
	analyzeManager.AnalyzeWithPattern(roadmap, analyzeManager.Pattern2)

	// m.sumResultInTotalRoad(roadmap)

	return roadmap
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPattern(roadmap *roadmap.Roadmap, pattern int) {
	if pattern == 1 {
		analyzeManager.AnalyzeWithPatternA(roadmap)
		analyzeManager.sumResultInTotalRoad(roadmap)
	}
	if pattern == 2 {
		analyzeManager.AnalyzeWithPatternB(roadmap)
		analyzeManager.sumResultInTotalRoad(roadmap)
	}
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPatternA(roadmap *roadmap.Roadmap) {
	analyzeManager.PatternAInBigRoad(roadmap.BigRoad)
	analyzeManager.PatternAInBigEyeRoad(roadmap.BigEyeRoad)
	analyzeManager.PatternAInSmallRoad(roadmap.SmallRoad)
	analyzeManager.PatternAInCockroachRoad(roadmap.CockroachRoad)
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPatternB(roadmap *roadmap.Roadmap) {
	analyzeManager.PatternBInBigRoad(roadmap.BigRoad)
	analyzeManager.PatternBInBigEyeRoad(roadmap.BigEyeRoad)
	analyzeManager.PatternBInSmallRoad(roadmap.SmallRoad)
	analyzeManager.PatternBInCockroachRoad(roadmap.CockroachRoad)
}

func GetPatterns() map[int]string {
	return Patterns
}

func (analyzeManager *AnalyzeManager) sumResultInTotalRoad(r *roadmap.Roadmap) {
	if len(r.TotalRoad.Columns) == 0 {
		return
	}

	// sum prediction result
	analyzeManager.sumPredictions(analyzeManager.Predictions)
	analyzeManager.drawTotalRoad(r.TotalRoad)
}

func (analyzeManager *AnalyzeManager) sumPredictions(predictions *Predictions) {
	betArea1Total := 0
	betArea2Total := 0

	if predictions.BigRoad.Bet != 0 {
		if predictions.BigRoad.BetArea == 1 {
			betArea1Total += predictions.BigRoad.Bet
		} else if predictions.BigRoad.BetArea == 2 {
			betArea2Total += predictions.BigRoad.Bet
		}
	}

	if predictions.BigEyeRoad.Bet != 0 {
		if predictions.BigEyeRoad.BetArea == 1 {
			betArea1Total += predictions.BigEyeRoad.Bet
		} else if predictions.BigEyeRoad.BetArea == 2 {
			betArea2Total += predictions.BigEyeRoad.Bet
		}
	}

	if predictions.SmallRoad.Bet != 0 {
		if predictions.SmallRoad.BetArea == 1 {
			betArea1Total += predictions.SmallRoad.Bet
		} else if predictions.SmallRoad.BetArea == 2 {
			betArea2Total += predictions.SmallRoad.Bet
		}
	}

	if predictions.CockroachRoad.Bet != 0 {
		if predictions.CockroachRoad.BetArea == 1 {
			betArea1Total += predictions.CockroachRoad.Bet
		} else if predictions.CockroachRoad.BetArea == 2 {
			betArea2Total += predictions.CockroachRoad.Bet
		}
	}

	if betArea1Total > betArea2Total {
		betArea1Total -= betArea2Total
		analyzeManager.Predictions.TotalRoad.Bet = betArea1Total
		analyzeManager.Predictions.TotalRoad.BetArea = 1
	} else if betArea1Total < betArea2Total {
		betArea2Total -= betArea1Total
		analyzeManager.Predictions.TotalRoad.Bet = betArea2Total
		analyzeManager.Predictions.TotalRoad.BetArea = 2
	} else {
		analyzeManager.Predictions.TotalRoad.Bet = 0
		analyzeManager.Predictions.TotalRoad.BetArea = 0
	}
}

// draw Prediction result in Total Road, only draw Result to
func (analyzeManager *AnalyzeManager) drawTotalRoad(road *roadmap.BigRoad) {
	if analyzeManager.Predictions.TotalRoad.Bet == 0 {
		return
	}
	lastColumn := road.Columns[len(road.Columns)-1]
	lastBlock := lastColumn.Blocks[len(lastColumn.Blocks)-1]

	if lastBlock.Symbol == roadmap.Symbol_Banker {
		if analyzeManager.Predictions.TotalRoad.BetArea == 1 {
			lastColumn.Blocks = append(lastColumn.Blocks, &roadmap.Block{
				Symbol: roadmap.Symbol_OnlyResult,
				Result: int32(analyzeManager.Predictions.TotalRoad.Bet),
			})
			return
		} else if analyzeManager.Predictions.TotalRoad.BetArea == 2 {
			road.Columns = append(road.Columns, &roadmap.Column{
				Blocks: []*roadmap.Block{
					{
						Symbol: roadmap.Symbol_OnlyResult,
						Result: int32(analyzeManager.Predictions.TotalRoad.Bet),
					},
				},
			})
			return
		}
	}
	if lastBlock.Symbol == roadmap.Symbol_Player {
		if analyzeManager.Predictions.TotalRoad.BetArea == 1 {
			road.Columns = append(road.Columns, &roadmap.Column{
				Blocks: []*roadmap.Block{
					{
						Symbol: roadmap.Symbol_OnlyResult,
						Result: int32(analyzeManager.Predictions.TotalRoad.Bet),
					},
				},
			})
			return
		} else if analyzeManager.Predictions.TotalRoad.BetArea == 2 {
			lastColumn.Blocks = append(lastColumn.Blocks, &roadmap.Block{
				Symbol: roadmap.Symbol_OnlyResult,
				Result: int32(analyzeManager.Predictions.TotalRoad.Bet),
			})
			return
		}
	}
}
