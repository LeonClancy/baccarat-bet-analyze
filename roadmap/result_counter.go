package roadmap

type Result struct {
	TieCount    int
	PlayerCount int
	BankerCount int
}

type RoadmapsResultCount struct {
	BigRoadCounts       *Result
	BigEyeRoadCounts    *Result
	SmallRoadCounts     *Result
	CockroachRoadCounts *Result
}
