package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

var Patterns = map[int]string{
	0: "無",
	1: "A",
	2: "B",
	3: "A1",
}

type AnalyzeManager struct {
	Pattern1       *Pattern
	Pattern2       *Pattern
	Predictions    *Predictions `json:"predictions"`
	AskRoadResults *AskRoadResults
}

type Pattern struct {
	Prediction  *Predictions
	PatternType int
}

type AskRoadResults struct {
	BankerAskRoadResult *roadmap.AskRoadResult
	PlayerAskRoadResult *roadmap.AskRoadResult
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
		Pattern1: &Pattern{
			PatternType: 0,
			Prediction: &Predictions{
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
		},
		Pattern2: &Pattern{
			PatternType: 0,
			Prediction: &Predictions{
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
		},
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
		AskRoadResults: &AskRoadResults{
			BankerAskRoadResult: &roadmap.AskRoadResult{},
			PlayerAskRoadResult: &roadmap.AskRoadResult{},
		},
	}
}

func (analyzeManager *AnalyzeManager) Analyze(roadmap *roadmap.Roadmap, bankerAskRoadResult roadmap.AskRoadResult, playerAskRoadResult roadmap.AskRoadResult) *roadmap.Roadmap {
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

	analyzeManager.initPredictions()

	analyzeManager.AskRoadResults.BankerAskRoadResult = &bankerAskRoadResult
	analyzeManager.AskRoadResults.PlayerAskRoadResult = &playerAskRoadResult

	analyzeManager.AnalyzeWithPattern(roadmap, analyzeManager.Pattern1)
	analyzeManager.AnalyzeWithPattern(roadmap, analyzeManager.Pattern2)

	analyzeManager.sumResultInTotalRoad(roadmap)

	return roadmap
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPattern(roadmap *roadmap.Roadmap, pattern *Pattern) {
	if pattern.PatternType == 1 {
		analyzeManager.AnalyzeWithPatternA(roadmap, pattern)
	}
	if pattern.PatternType == 2 {
		analyzeManager.AnalyzeWithPatternB(roadmap, pattern)
	}
	if pattern.PatternType == 3 {
		analyzeManager.AnalyzeWithPatternA1(roadmap, pattern)
	}
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPatternA(roadmap *roadmap.Roadmap, pattern *Pattern) {
	analyzeManager.PatternAInBigRoad(roadmap.BigRoad, pattern.Prediction.BigRoad)
	analyzeManager.PatternAInBigEyeRoad(roadmap.BigEyeRoad, pattern.Prediction.BigEyeRoad)
	analyzeManager.PatternAInSmallRoad(roadmap.SmallRoad, pattern.Prediction.SmallRoad)
	analyzeManager.PatternAInCockroachRoad(roadmap.CockroachRoad, pattern.Prediction.CockroachRoad)
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPatternB(roadmap *roadmap.Roadmap, pattern *Pattern) {
	analyzeManager.PatternBInBigRoad(roadmap.BigRoad, pattern.Prediction.BigRoad)
	analyzeManager.PatternBInBigEyeRoad(roadmap.BigEyeRoad, pattern.Prediction.BigEyeRoad)
	analyzeManager.PatternBInSmallRoad(roadmap.SmallRoad, pattern.Prediction.SmallRoad)
	analyzeManager.PatternBInCockroachRoad(roadmap.CockroachRoad, pattern.Prediction.CockroachRoad)
}

func (analyzeManager *AnalyzeManager) AnalyzeWithPatternA1(roadmap *roadmap.Roadmap, pattern *Pattern) {
	analyzeManager.PatternA1BigRoad(roadmap.BigRoad, pattern.Prediction.BigRoad)
	analyzeManager.PatternA1BigEyeRoad(roadmap.BigEyeRoad, pattern.Prediction.BigEyeRoad)
	analyzeManager.PatternA1SmallRoad(roadmap.SmallRoad, pattern.Prediction.SmallRoad)
	analyzeManager.PatternA1CockroachRoad(roadmap.CockroachRoad, pattern.Prediction.CockroachRoad)
}

func GetPatterns() map[int]string {
	return Patterns
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
						Symbol: roadmap.Symbol_OnlyResultAndNewLine,
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
						Symbol: roadmap.Symbol_OnlyResultAndNewLine,
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

func (analyzeManager *AnalyzeManager) sumResults(r *roadmap.Roadmap) {
	r.TotalRoad.Result = 0
	r.TotalRoad.TotalBet = 0

	for i := range r.TotalRoad.Columns {
		for j := range r.TotalRoad.Columns[i].Blocks {
			block := r.TotalRoad.Columns[i].Blocks[j]
			if block.Symbol != roadmap.Symbol_OnlyResult && block.Symbol != roadmap.Symbol_OnlyResultAndNewLine {
				r.TotalRoad.Result += block.Result
				if block.Result < 0 {
					r.TotalRoad.TotalBet -= block.Result
				} else {
					r.TotalRoad.TotalBet += block.Result
				}
			}
		}
	}

	r.BigRoad.Result = 0
	r.BigRoad.TotalBet = 0
	for i := range r.BigRoad.Columns {
		for j := range r.BigRoad.Columns[i].Blocks {
			r.BigRoad.Result += r.BigRoad.Columns[i].Blocks[j].Result
			if r.BigRoad.Columns[i].Blocks[j].Result < 0 {
				r.BigRoad.TotalBet -= r.BigRoad.Columns[i].Blocks[j].Result
			} else {
				r.BigRoad.TotalBet += r.BigRoad.Columns[i].Blocks[j].Result
			}
		}
	}

	r.BigEyeRoad.Result = 0
	r.BigEyeRoad.TotalBet = 0
	for i := range r.BigEyeRoad.Columns {
		for j := range r.BigEyeRoad.Columns[i].Blocks {
			r.BigEyeRoad.Result += r.BigEyeRoad.Columns[i].Blocks[j].Result
			if r.BigEyeRoad.Columns[i].Blocks[j].Result < 0 {
				r.BigEyeRoad.TotalBet -= r.BigEyeRoad.Columns[i].Blocks[j].Result
			} else {
				r.BigEyeRoad.TotalBet += r.BigEyeRoad.Columns[i].Blocks[j].Result
			}
		}
	}

	r.SmallRoad.Result = 0
	r.SmallRoad.TotalBet = 0
	for i := range r.SmallRoad.Columns {
		for j := range r.SmallRoad.Columns[i].Blocks {
			r.SmallRoad.Result += r.SmallRoad.Columns[i].Blocks[j].Result
			if r.SmallRoad.Columns[i].Blocks[j].Result < 0 {
				r.SmallRoad.TotalBet -= r.SmallRoad.Columns[i].Blocks[j].Result
			} else {
				r.SmallRoad.TotalBet += r.SmallRoad.Columns[i].Blocks[j].Result
			}
		}
	}

	r.CockroachRoad.Result = 0
	r.CockroachRoad.TotalBet = 0
	for i := range r.CockroachRoad.Columns {
		for j := range r.CockroachRoad.Columns[i].Blocks {
			r.CockroachRoad.Result += r.CockroachRoad.Columns[i].Blocks[j].Result
			if r.CockroachRoad.Columns[i].Blocks[j].Result < 0 {
				r.CockroachRoad.TotalBet -= r.CockroachRoad.Columns[i].Blocks[j].Result
			} else {
				r.CockroachRoad.TotalBet += r.CockroachRoad.Columns[i].Blocks[j].Result
			}
		}
	}
}

func (analyzeManager *AnalyzeManager) initPredictions() {
	analyzeManager.Predictions.BigRoad.Bet = 0
	analyzeManager.Predictions.BigRoad.BetArea = 0
	analyzeManager.Predictions.BigEyeRoad.Bet = 0
	analyzeManager.Predictions.BigEyeRoad.BetArea = 0
	analyzeManager.Predictions.SmallRoad.Bet = 0
	analyzeManager.Predictions.SmallRoad.BetArea = 0
	analyzeManager.Predictions.CockroachRoad.Bet = 0
	analyzeManager.Predictions.CockroachRoad.BetArea = 0
	analyzeManager.Predictions.TotalRoad.Bet = 0
	analyzeManager.Predictions.TotalRoad.BetArea = 0

	analyzeManager.Pattern1.Prediction.BigRoad.Bet = 0
	analyzeManager.Pattern1.Prediction.BigRoad.BetArea = 0
	analyzeManager.Pattern1.Prediction.BigEyeRoad.Bet = 0
	analyzeManager.Pattern1.Prediction.BigEyeRoad.BetArea = 0
	analyzeManager.Pattern1.Prediction.SmallRoad.Bet = 0
	analyzeManager.Pattern1.Prediction.SmallRoad.BetArea = 0
	analyzeManager.Pattern1.Prediction.CockroachRoad.Bet = 0
	analyzeManager.Pattern1.Prediction.CockroachRoad.BetArea = 0
	analyzeManager.Pattern1.Prediction.TotalRoad.Bet = 0
	analyzeManager.Pattern1.Prediction.TotalRoad.BetArea = 0

	analyzeManager.Pattern2.Prediction.BigRoad.Bet = 0
	analyzeManager.Pattern2.Prediction.BigRoad.BetArea = 0
	analyzeManager.Pattern2.Prediction.BigEyeRoad.Bet = 0
	analyzeManager.Pattern2.Prediction.BigEyeRoad.BetArea = 0
	analyzeManager.Pattern2.Prediction.SmallRoad.Bet = 0
	analyzeManager.Pattern2.Prediction.SmallRoad.BetArea = 0
	analyzeManager.Pattern2.Prediction.CockroachRoad.Bet = 0
	analyzeManager.Pattern2.Prediction.CockroachRoad.BetArea = 0
	analyzeManager.Pattern2.Prediction.TotalRoad.Bet = 0
	analyzeManager.Pattern2.Prediction.TotalRoad.BetArea = 0
}
