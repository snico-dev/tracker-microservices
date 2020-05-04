package timeconvert

import (
	"strings"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/convert"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/utils"
)

//TimeConvert interface
type TimeConvert interface {
	GetHexUTC() string
}

type timeConvert struct {
	unixDate UnixDate
}

//NewTimeConvert contructor
func NewTimeConvert(unixDate UnixDate) TimeConvert {
	return &timeConvert{
		unixDate: unixDate,
	}
}

//GetHexUTC return date in hex formater
func (t timeConvert) GetHexUTC() string {
	hexDate := convert.FromDecimalToHex(t.unixDate.GetNowSeconds())
	return utils.DescInGroup(strings.ToUpper(hexDate))
}

//GetFullYear get full from short year
func GetFullYear(shortYear int) int {
	longYear := 0
	if shortYear >= 60 { // this should be the number where you think it stops to be 20xx (like 15 for 2015; for every number after that it will be 19xx)
		longYear = shortYear + 1900
	} else {
		longYear = shortYear + 2000
	}
	return longYear
}
