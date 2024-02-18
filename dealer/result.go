package dealer

type Result int32

const (
	Result_Default    Result = 0
	Result_Banker     Result = 1
	Result_Player     Result = 2
	Result_Tie        Result = 3
	Result_BankerPair Result = 4
	Result_PlayerPair Result = 5
)
