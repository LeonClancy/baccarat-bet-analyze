package roadmap

type Block struct {
	Symbol   Symbol `json:"symbol"`
	TieCount int32  `json:"tieCount"`
}

type Column struct {
	Blocks []*Block `json:"blocks"`
}

type BeadPlate struct {
	Blocks []*Block `json:"blocks"`
}

type BigRoad struct {
	Columns []*Column `json:"columns"`
}

type BigEyeRoad struct {
	Columns []*Column `json:"columns"`
}

type SmallRoad struct {
	Columns []*Column `json:"columns"`
}

type CockroachRoad struct {
	Columns []*Column `json:"columns"`
}

type Roadmap struct {
	BeadPlate     *BeadPlate     `json:"bead_plate"`
	BigRoad       *BigRoad       `json:"big_road"`
	BigEyeRoad    *BigEyeRoad    `json:"big_eye_road"`
	SmallRoad     *SmallRoad     `json:"small_road"`
	CockroachRoad *CockroachRoad `json:"cockroach_road"`
}
