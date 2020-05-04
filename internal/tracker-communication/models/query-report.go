package models


//Query fetch informations from vehicle ecu
type Query struct {
	CmdSeq          int
	RespCount       int
	RespIndex       int
	FailCount       int
	FailTagArray    []string
	SuccessCount    int
	SuccessTLVArray []TLV
}
