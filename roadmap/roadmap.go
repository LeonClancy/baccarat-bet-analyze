package roadmap

type Block struct {
	Symbol   Symbol `json:"symbol"`
	TieCount int32  `json:"tieCount"`
	Result   int32  `json:"result"`
}

type Column struct {
	Blocks []*Block `json:"blocks"`
	Result int32    `json:"result"`
}

type BeadPlate struct {
	Blocks []*Block `json:"blocks"`
}

type BigRoad struct {
	Columns      []*Column `json:"columns"`
	LevelManager *LevelManager
	Result       int32
	TotalBet     int32
}

type BigEyeRoad struct {
	Columns      []*Column `json:"columns"`
	LevelManager *LevelManager
	Result       int32
	TotalBet     int32
}

type SmallRoad struct {
	Columns      []*Column `json:"columns"`
	LevelManager *LevelManager
	Result       int32
	TotalBet     int32
}

type CockroachRoad struct {
	Columns      []*Column `json:"columns"`
	LevelManager *LevelManager
	Result       int32
	TotalBet     int32
}

type AskRoadResult struct {
	BigEyeRoadNext    *Block
	SmallRoadNext     *Block
	CockroachRoadNext *Block
}

type Roadmap struct {
	TotalRoad     *BigRoad       `json:"totalRoad"`
	BeadPlate     *BeadPlate     `json:"beadPlate"`
	BigRoad       *BigRoad       `json:"bigRoad"`
	BigEyeRoad    *BigEyeRoad    `json:"bigEyeRoad"`
	SmallRoad     *SmallRoad     `json:"smallRoad"`
	CockroachRoad *CockroachRoad `json:"cockroachRoad"`
}
