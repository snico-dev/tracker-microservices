package models

import "time"

//StateData data type
type StateData struct {
	LastAcconTime      time.Time
	UTCTime            time.Time
	TotalTripMileage   float64
	CurrentTripMileage float64
	TotalFuel          float64
	CurrentFuel        float64
}
