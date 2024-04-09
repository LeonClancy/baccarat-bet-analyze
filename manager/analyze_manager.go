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
}

func NewAnalyzeManager() *AnalyzeManager {
	return &AnalyzeManager{}
}

func (m *AnalyzeManager) Analyze(roadmap *roadmap.Roadmap) *roadmap.Roadmap {
	for _, c := range roadmap.BigRoad.Columns {
		c.Result = 0
		for _, b := range c.Blocks {
			b.Result = 0		
		}
	}
	
	if m.Pattern1 == 1 {
		for _, c := range roadmap.BigRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 1
				c.Result += b.Result
			}
		}
	}

	if m.Pattern2 == 2 {
		for _, c := range roadmap.BigRoad.Columns {
			for _, b := range c.Blocks {
				b.Result += 2
				c.Result += b.Result
			}
		}
	}
	
	return roadmap	
}
