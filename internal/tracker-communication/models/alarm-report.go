package models

//AlarmReport type
type AlarmReport struct {
	AlarmSeq   int
	AlarmCount int
	AlarmArray []AlarmData
	StateData  StateData
	GpsData    GpsData
}
