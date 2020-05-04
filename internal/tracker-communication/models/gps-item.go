package models

import "time"

//GpsItem data type
type GpsItem struct {
	DateTime  time.Time
	Latitude  float64
	Longitude float64
	Speed     float64
	Direction float64
	Flag      string
}
