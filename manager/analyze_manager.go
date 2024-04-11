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
	
	if m.Pattern1 == 1 {
		for _, c := range roadmap.BigRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 1
				c.Result += b.Result
			}
		}

		m.Predictions.BigRoad.Bet = 1
		m.Predictions.BigRoad.BetArea = 1

		for _, c := range roadmap.BigEyeRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 1
				c.Result += b.Result
			}
		}

		m.Predictions.BigEyeRoad.Bet = 1
		m.Predictions.BigEyeRoad.BetArea = 1

		for _, c := range roadmap.SmallRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 1
				c.Result += b.Result
			}
		}

		m.Predictions.SmallRoad.Bet = 1
		m.Predictions.SmallRoad.BetArea = 1

		for _, c := range roadmap.CockroachRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 1
				c.Result += b.Result
			}
		}

		m.Predictions.CockroachRoad.Bet = 1
		m.Predictions.CockroachRoad.BetArea = 1
	}

	if m.Pattern2 == 2 {
		for _, c := range roadmap.BigRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 2
				c.Result += b.Result
			}
		}

		m.Predictions.BigRoad.Bet = 2
		m.Predictions.BigRoad.BetArea = 2

		for _, c := range roadmap.BigEyeRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 2
				c.Result += b.Result
			}
		}

		m.Predictions.BigEyeRoad.Bet = 2
		m.Predictions.BigEyeRoad.BetArea = 2

		for _, c := range roadmap.SmallRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 2
				c.Result += b.Result
			}
		}

		m.Predictions.SmallRoad.Bet = 2
		m.Predictions.SmallRoad.BetArea = 2

		for _, c := range roadmap.CockroachRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 2
				c.Result += b.Result
			}
		}

		m.Predictions.CockroachRoad.Bet = 2
		m.Predictions.CockroachRoad.BetArea = 2
	}
	
	return roadmap
}

func GetPatterns() map[int]string {
	return Patterns
}
