package manager

import (
	"github.com/LeonClancy/baccarat-bet-analyze/roadmap"
)

func (analyzeManager *AnalyzeManager) sumResultInTotalRoad(r *roadmap.Roadmap) {
	if len(r.TotalRoad.Columns) == 0 {
		return
	}

	if analyzeManager.Pattern1.PatternType == 4 && analyzeManager.Pattern2.PatternType == 6 {
		if len(r.TotalRoad.Columns) > 0 {
			if len(r.TotalRoad.Columns[len(r.TotalRoad.Columns)-1].Blocks) >= 4 {
				analyzeManager.initPredictions()
			}
		}
	}


	// sum prediction result
	analyzeManager.sumBigRaodPredicitons(analyzeManager.Pattern1.Prediction.BigRoad, analyzeManager.Pattern2.Prediction.BigRoad)
	analyzeManager.sumBigEyeRaodPredicitons(analyzeManager.Pattern1.Prediction.BigEyeRoad, analyzeManager.Pattern2.Prediction.BigEyeRoad)
	analyzeManager.sumSmallRaodPredicitons(analyzeManager.Pattern1.Prediction.SmallRoad, analyzeManager.Pattern2.Prediction.SmallRoad)
	analyzeManager.sumCockroachRaodPredicitons(analyzeManager.Pattern1.Prediction.CockroachRoad, analyzeManager.Pattern2.Prediction.CockroachRoad)
	
	analyzeManager.sumPredictions(analyzeManager.Predictions)

	if analyzeManager.Pattern1.PatternType == 5 && analyzeManager.Pattern2.PatternType == 7 {
		if len(r.TotalRoad.Columns) > 0 {
			if len(r.TotalRoad.Columns[len(r.TotalRoad.Columns)-1].Blocks) >= 4 {
				analyzeManager.reversePredictions()
			}
		}
	}

	analyzeManager.sumResults(r)
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

func (analyzeManager *AnalyzeManager) sumBigRaodPredicitons(p1, p2 *Prediction) {
	if p1.BetArea == p2.BetArea {
		analyzeManager.Predictions.BigRoad.BetArea = p1.BetArea
		analyzeManager.Predictions.BigRoad.Bet = p1.Bet + p2.Bet
	}

	if p1.BetArea != p2.BetArea {
		if p1.Bet > p2.Bet {
			analyzeManager.Predictions.BigRoad.BetArea = p1.BetArea
			analyzeManager.Predictions.BigRoad.Bet = p1.Bet - p2.Bet
		} else if p1.Bet < p2.Bet {
			analyzeManager.Predictions.BigRoad.BetArea = p2.BetArea
			analyzeManager.Predictions.BigRoad.Bet = p2.Bet - p1.Bet
		}
	}
}

func (analyzeManager *AnalyzeManager) sumBigEyeRaodPredicitons(p1, p2 *Prediction) {
	if p1.BetArea == p2.BetArea {
		analyzeManager.Predictions.BigEyeRoad.BetArea = p1.BetArea
		analyzeManager.Predictions.BigEyeRoad.Bet = p1.Bet + p2.Bet
	}

	if p1.BetArea != p2.BetArea {
		if p1.Bet > p2.Bet {
			analyzeManager.Predictions.BigEyeRoad.BetArea = p1.BetArea
			analyzeManager.Predictions.BigEyeRoad.Bet = p1.Bet - p2.Bet
		} else if p1.Bet < p2.Bet {
			analyzeManager.Predictions.BigEyeRoad.BetArea = p2.BetArea
			analyzeManager.Predictions.BigEyeRoad.Bet = p2.Bet - p1.Bet
		}
	}
}

func (analyzeManager *AnalyzeManager) sumSmallRaodPredicitons(p1, p2 *Prediction) {
	if p1.BetArea == p2.BetArea {
		analyzeManager.Predictions.SmallRoad.BetArea = p1.BetArea
		analyzeManager.Predictions.SmallRoad.Bet = p1.Bet + p2.Bet
	}

	if p1.BetArea != p2.BetArea {
		if p1.Bet > p2.Bet {
			analyzeManager.Predictions.SmallRoad.BetArea = p1.BetArea
			analyzeManager.Predictions.SmallRoad.Bet = p1.Bet - p2.Bet
		} else if p1.Bet < p2.Bet {
			analyzeManager.Predictions.SmallRoad.BetArea = p2.BetArea
			analyzeManager.Predictions.SmallRoad.Bet = p2.Bet - p1.Bet
		}
	}
}

func (analyzeManager *AnalyzeManager) sumCockroachRaodPredicitons(p1, p2 *Prediction) {
	if p1.BetArea == p2.BetArea {
		analyzeManager.Predictions.CockroachRoad.BetArea = p1.BetArea
		analyzeManager.Predictions.CockroachRoad.Bet = p1.Bet + p2.Bet
	}

	if p1.BetArea != p2.BetArea {
		if p1.Bet > p2.Bet {
			analyzeManager.Predictions.CockroachRoad.BetArea = p1.BetArea
			analyzeManager.Predictions.CockroachRoad.Bet = p1.Bet - p2.Bet
		} else if p1.Bet < p2.Bet {
			analyzeManager.Predictions.CockroachRoad.BetArea = p2.BetArea
			analyzeManager.Predictions.CockroachRoad.Bet = p2.Bet - p1.Bet
		}
	}
}
