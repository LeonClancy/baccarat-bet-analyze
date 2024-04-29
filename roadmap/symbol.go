package roadmap

type Symbol int32

const (
	Symbol_BlockDefault              Symbol = 0
	Symbol_Banker                    Symbol = 1
	Symbol_Player                    Symbol = 2
	Symbol_Tie                       Symbol = 3
	Symbol_BankerAndBankerPair       Symbol = 4
	Symbol_BankerAndPlayerPair       Symbol = 5
	Symbol_BankerAndBothPair         Symbol = 6
	Symbol_PlayerAndBankerPair       Symbol = 7
	Symbol_PlayerAndPlayerPair       Symbol = 8
	Symbol_PlayerAndBothPair         Symbol = 9
	Symbol_TieAndBankerPair          Symbol = 10
	Symbol_TieAndPlayerPair          Symbol = 11
	Symbol_TieAndBothPair            Symbol = 12
	Symbol_BankerAndTie              Symbol = 13
	Symbol_BankerAndBankerPairAndTie Symbol = 14
	Symbol_BankerAndPlayerPairAndTie Symbol = 15
	Symbol_BankerAndBothPairAndTie   Symbol = 16
	Symbol_PlayerAndTie              Symbol = 17
	Symbol_PlayerAndBankerPairAndTie Symbol = 18
	Symbol_PlayerAndPlayerPairAndTie Symbol = 19
	Symbol_PlayerAndBothPairAndTie   Symbol = 20
	Symbol_OnlyResult                Symbol = 21
	Symbol_OnlyResultAndNewLine      Symbol = 22
)
