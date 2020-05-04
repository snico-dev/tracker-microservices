package timeconvert

import "time"

type unixDateStub struct {
	UnixDate,
	seconds int64
}

func (u unixDateStub) GetNowSeconds() int64 { return u.seconds }
func (u unixDateStub) GetDateFromSeconds(seconds int) time.Time {
	return time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local)
}

//NewUnixDateStub contructor stub
func NewUnixDateStub(seconds int64) UnixDate {
	return &unixDateStub{
		seconds: seconds,
	}
}
