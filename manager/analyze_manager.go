package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

var Patterns = map[int]string{
	1: "A",
	2: "B",
}

type AnalyzeManager struct {
	Pattern1 int
	Pattern2 int
	Predictions *Predictions `json:"predictions"`
}

type Predictions struct {
	BigRoad *Prediction `json:"bigRoad"`
	BigEyeRoad *Prediction `json:"bigEyeRoad"`
	SmallRoad *Prediction `json:"smallRoad"`
	CockroachRoad *Prediction `json:"cockroachRoad"`
}

type Prediction struct {
	Bet int `json:"bet"`
	BetArea int `json:"betArea"`
}

func NewAnalyzeManager() *AnalyzeManager {
	return &AnalyzeManager{
		Pattern1: 0,
		Pattern2: 0,
		Predictions: &Predictions{
			BigRoad: &Prediction{
				Bet: 0,
				BetArea: 0,
			},
			BigEyeRoad: &Prediction{
				Bet: 0,
				BetArea: 0,
			},
			SmallRoad: &Prediction{
				Bet: 0,
				BetArea: 0,
			},
			CockroachRoad: &Prediction{
				Bet: 0,
				BetArea: 0,
			},
		},
	}
}

func (m *AnalyzeManager) Analyze(roadmap *roadmap.Roadmap) *roadmap.Roadmap {
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

	m.Predictions.BigEyeRoad.Bet = 0
	m.Predictions.BigEyeRoad.BetArea = 0
	m.Predictions.BigRoad.Bet = 0
	m.Predictions.BigRoad.BetArea = 0
	m.Predictions.SmallRoad.Bet = 0
	m.Predictions.SmallRoad.BetArea = 0
	m.Predictions.CockroachRoad.Bet = 0
	m.Predictions.CockroachRoad.BetArea = 0

	m.AnalyzeWithPattern(roadmap, m.Pattern1)
	m.AnalyzeWithPattern(roadmap, m.Pattern2)
	
	// m.sumResultInTotalRoad(roadmap)

	return roadmap
}

func (m *AnalyzeManager) AnalyzeWithPattern(roadmap *roadmap.Roadmap, pattern int) {
	if pattern == 1 {
		m.AnalyzeWithPatternA(roadmap)
	}
	if pattern == 2 {
		m.AnalyzeWithPatternB(roadmap)
	}
}

func (m *AnalyzeManager) AnalyzeWithPatternA(roadmap *roadmap.Roadmap) {
	m.PatternAInBigRoad(roadmap.BigRoad)
}

func (m *AnalyzeManager) AnalyzeWithPatternB(roadmap *roadmap.Roadmap) {

}

func GetPatterns() map[int]string {
	return Patterns
}

func (m *AnalyzeManager) sumResultInTotalRoad(r *roadmap.Roadmap) {
	if len(r.TotalRoad.Columns) == 0 {
		return
	}

	for i := range r.BigRoad.Columns {
		r.TotalRoad.Columns[i].Result += r.BigRoad.Columns[i].Result
		for j := range r.BigRoad.Columns[i].Blocks {
			if r.TotalRoad.Columns[i].Blocks[j] == nil {
				r.TotalRoad.Columns[i].Blocks[j] = &roadmap.Block{
					Result: 0,
				}
			}
			r.TotalRoad.Columns[i].Blocks[j].Result += r.BigRoad.Columns[i].Blocks[j].Result
		}
	}

	for i := range r.SmallRoad.Columns {
		r.TotalRoad.Columns[i].Result += r.SmallRoad.Columns[i].Result
		for j := range r.SmallRoad.Columns[i].Blocks {
			if len(r.TotalRoad.Columns[i].Blocks) < len(r.SmallRoad.Columns[i].Blocks) {
				r.TotalRoad.Columns[i].Blocks = append(r.TotalRoad.Columns[i].Blocks, &roadmap.Block{
					Result: 0,
				})
			}
			if r.TotalRoad.Columns[i].Blocks[j] == nil {
				r.TotalRoad.Columns[i].Blocks[j] = &roadmap.Block{
					Result: 0,
				}
			}
			r.TotalRoad.Columns[i].Blocks[j].Result += r.SmallRoad.Columns[i].Blocks[j].Result
		}
	}

	for i := range r.BigEyeRoad.Columns {
		if len(r.BigEyeRoad.Columns) > len(r.TotalRoad.Columns) {
			r.TotalRoad.Columns = append(r.TotalRoad.Columns, &roadmap.Column{
				Blocks: make([]*roadmap.Block, len(r.BigEyeRoad.Columns[i].Blocks)),
			})
		}
		r.TotalRoad.Columns[i].Result += r.BigEyeRoad.Columns[i].Result
		for j := range r.BigEyeRoad.Columns[i].Blocks {
			if len(r.TotalRoad.Columns[i].Blocks) < len(r.BigEyeRoad.Columns[i].Blocks) {
				r.TotalRoad.Columns[i].Blocks = append(r.TotalRoad.Columns[i].Blocks, &roadmap.Block{
					Result: 0,
				})
			}
			if r.TotalRoad.Columns[i].Blocks[j] == nil {
				r.TotalRoad.Columns[i].Blocks[j] = &roadmap.Block{
					Result: 0,
				}
			}
			r.TotalRoad.Columns[i].Blocks[j].Result += r.BigEyeRoad.Columns[i].Blocks[j].Result
		}
	}

	for i := range r.CockroachRoad.Columns {
		if len(r.CockroachRoad.Columns) > len(r.TotalRoad.Columns) {
			r.TotalRoad.Columns = append(r.TotalRoad.Columns, &roadmap.Column{
				Blocks: make([]*roadmap.Block, len(r.CockroachRoad.Columns[i].Blocks)),
			})
		}
		r.TotalRoad.Columns[i].Result += r.CockroachRoad.Columns[i].Result
		for j := range r.CockroachRoad.Columns[i].Blocks {
			if len(r.TotalRoad.Columns[i].Blocks) < len(r.CockroachRoad.Columns[i].Blocks) {
				r.TotalRoad.Columns[i].Blocks = append(r.TotalRoad.Columns[i].Blocks, &roadmap.Block{
					Result: 0,
				})
			}
			if r.TotalRoad.Columns[i].Blocks[j] == nil {
				r.TotalRoad.Columns[i].Blocks[j] = &roadmap.Block{
					Result: 0,
				}
			}
			r.TotalRoad.Columns[i].Blocks[j].Result += r.CockroachRoad.Columns[i].Blocks[j].Result
		}
	}
}