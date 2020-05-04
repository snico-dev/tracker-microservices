package timeconvert

type timeConvertStub struct {
	TimeConvert
	hexUTC string
}

//NewTimeConvertStub contructor
func NewTimeConvertStub(hexUTC string) TimeConvert {
	return &timeConvertStub{
		hexUTC: hexUTC,
	}
}

//GetHexUTC return moock date in hex formater
func (t timeConvertStub) GetHexUTC() string {
	return t.hexUTC
}
