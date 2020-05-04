package models

//GpsReport data type
type GpsReport struct {
	Flag      string
	StateData StateData
	GpsData   GpsData
	RpmData   RpmData
}
