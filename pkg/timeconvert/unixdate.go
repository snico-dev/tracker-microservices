package timeconvert

import (
	"time"
)

//UnixDate interface
type UnixDate interface {
	GetNowSeconds() int64
	GetDateFromSeconds(seconds int) time.Time
}

type unixDate struct{}

//NewUnixDate constructor
func NewUnixDate() UnixDate {
	return &unixDate{}
}

func (u unixDate) GetDateFromSeconds(seconds int) time.Time {
	location := GetLocation()
	return time.Date(1970, time.January, 1, 0, 0, 0, 0, location).Add(time.Second * time.Duration(seconds))
}

func (u unixDate) GetNowSeconds() int64 {
	location := GetLocation()
	timeNow := time.Now().In(location).UnixNano()
	return timeNow / int64(time.Second)
}

//GetLocation get current location
func GetLocation() *time.Location {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	return location
}
