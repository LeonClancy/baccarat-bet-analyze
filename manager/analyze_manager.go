package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

var Patterns = map[int]string{
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
	for _, c := range roadmap.TotalRoad.Columns {
		c.Result = 0
		for _, b := range c.Blocks {
			b.Result = 0
		}
	}

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

	for i := 0; i < len(r.BigRoad.Columns); i++ {
		r.TotalRoad.Columns[i].Result += r.BigRoad.Columns[i].Result
		for j := range r.BigRoad.Columns[i].Blocks {
			r.TotalRoad.Columns[i].Blocks[j].Result += r.BigRoad.Columns[i].Blocks[j].Result
		}
	}

	// for i := 0 ; i < len(r.BigEyeRoad.Columns) ; i++ {
	// 	r.TotalRoad.Columns[i].Result += r.BigEyeRoad.Columns[i].Result
	// 	// 如果統計路沒有該行，則要新增到該行
	// 	if len(r.TotalRoad.Columns) <= i {
	// 		r.TotalRoad.Columns = append(r.TotalRoad.Columns, &roadmap.Column{})
	// 	}
	// 	for j := range r.BigEyeRoad.Columns[i].Blocks {
	// 		// 如果統計路沒有該區塊，則要新增到該行該格的 result
	// 		if len(r.TotalRoad.Columns[i].Blocks) <= j {
	// 			r.TotalRoad.Columns[i].Blocks = append(r.TotalRoad.Columns[i].Blocks, &roadmap.Block{})
	// 		}
	// 		r.TotalRoad.Columns[i].Blocks[j].Result += r.BigEyeRoad.Columns[i].Blocks[j].Result
	// 	}
	// }

}
