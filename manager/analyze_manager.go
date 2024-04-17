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

func (analyzeManger *AnalyzeManager) Analyze(roadmap *roadmap.Roadmap) *roadmap.Roadmap {
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

	analyzeManger.Predictions.BigEyeRoad.Bet = 0
	analyzeManger.Predictions.BigEyeRoad.BetArea = 0
	analyzeManger.Predictions.BigRoad.Bet = 0
	analyzeManger.Predictions.BigRoad.BetArea = 0
	analyzeManger.Predictions.SmallRoad.Bet = 0
	analyzeManger.Predictions.SmallRoad.BetArea = 0
	analyzeManger.Predictions.CockroachRoad.Bet = 0
	analyzeManger.Predictions.CockroachRoad.BetArea = 0

	analyzeManger.AnalyzeWithPattern(roadmap, analyzeManger.Pattern1)
	analyzeManger.AnalyzeWithPattern(roadmap, analyzeManger.Pattern2)

	// m.sumResultInTotalRoad(roadmap)

	return roadmap
}

func (analyzeManger *AnalyzeManager) AnalyzeWithPattern(roadmap *roadmap.Roadmap, pattern int) {
	if pattern == 1 {
		analyzeManger.AnalyzeWithPatternA(roadmap)
		analyzeManger.sumResultInTotalRoad(roadmap)
	}
	if pattern == 2 {
		analyzeManger.AnalyzeWithPatternB(roadmap)
		analyzeManger.sumResultInTotalRoad(roadmap)
	}
}

func (analyzeManger *AnalyzeManager) AnalyzeWithPatternA(roadmap *roadmap.Roadmap) {
	analyzeManger.PatternAInBigRoad(roadmap.BigRoad)
	//analyzeManger.PatternAInBigEyeRoad(roadmap.BigEyeRoad)
	// m.PatternAInSmallRoad(roadmap.SmallRoad)
	// m.PatternAInCockroachRoad(roadmap.CockroachRoad)
}

func (analyzeManger *AnalyzeManager) AnalyzeWithPatternB(roadmap *roadmap.Roadmap) {
	analyzeManger.PatternBInBigRoad(roadmap.BigRoad)
	//analyzeManger.PatternBInBigEyeRoad(roadmap.BigEyeRoad)
}

func GetPatterns() map[int]string {
	return Patterns
}

func (analyzeManger *AnalyzeManager) sumResultInTotalRoad(r *roadmap.Roadmap) {
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
