package services

import (
	"math"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/buffer"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/convert"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
)

//IDataType interface
type IDataType interface {
	GetGpsItem(bufferpack []byte) (models.GpsItem, error)
	GetGpsData(bufferpack []byte) (models.GpsData, error)
	GetStateData(bufferpack []byte) (models.StateData, error)
	GetRpmItem(bufferpack []byte) (models.RpmItem, error)
	GetRpmData(bufferpack []byte) (models.RpmData, error)
	GetAlarmData(bufferpack []byte) (models.AlarmData, error)
	GetTLVData(bufferpack []byte) (models.TLV, error)
}

type dataType struct {
	buff     buffer.IBuffer
	unixDate timeconvert.UnixDate
}

//NewDataType create instance
func NewDataType(buff buffer.IBuffer, unixDate timeconvert.UnixDate) IDataType {
	return &dataType{
		buff:     buff,
		unixDate: unixDate,
	}
}

//GetGpsItem from bytes do data struct
func (t *dataType) GetGpsItem(bufferpack []byte) (models.GpsItem, error) {

	day, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))
	month, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))
	year, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))
	hour, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))
	minute, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))
	second, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))
	lat, error := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 4))
	long, error := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 4))
	speed, error := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 2))
	direction, error := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 2))
	flag, error := convert.FromByteToHex(t.buff.Slice(bufferpack, 1))
	lowFlag, error := convert.FromHexToBitArray(flag)

	if len(lowFlag) < 2 {
		lowFlag = append(lowFlag, false)
	}

	data := models.GpsItem{
		DateTime:  time.Date(timeconvert.GetFullYear(year), time.Month(month), day, hour, minute, second, 0, timeconvert.GetLocation()),
		Latitude:  normalizeCoordenate(lat, lowFlag[len(lowFlag)-1]),
		Longitude: normalizeCoordenate(long, lowFlag[len(lowFlag)-2]),
		Direction: normalizeDirection(direction),
		Speed:     convertCmsToKm(speed),
		Flag:      flag,
	}

	return data, error
}

func (t *dataType) GetStateData(bufferpack []byte) (models.StateData, error) {
	lastAcconTime, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 4))
	utcTime, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 4))
	totalTripMileage, err := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 4))
	currentTripMileage, err := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 4))
	totalFuel, err := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 4))
	currentFuel, err := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 2))
	t.buff.Slice(bufferpack, 4)
	t.buff.Slice(bufferpack, 8)

	return models.StateData{
		LastAcconTime:      t.unixDate.GetDateFromSeconds(lastAcconTime),
		UTCTime:            t.unixDate.GetDateFromSeconds(utcTime),
		TotalTripMileage:   totalTripMileage,
		CurrentTripMileage: currentTripMileage,
		TotalFuel:          getInLiter(totalFuel),
		CurrentFuel:        getInLiter(currentFuel),
	}, err
}

func getInLiter(consumption float64) float64 {
	return consumption * 0.01
}

func (t *dataType) GetRpmItem(bufferpack []byte) (models.RpmItem, error) {
	rpm, error := convert.FromByteHexReverseToFloat64(t.buff.Slice(bufferpack, 2))
	return models.RpmItem{Rpm: rpm}, error
}

func (t *dataType) GetRpmData(bufferpack []byte) (models.RpmData, error) {
	count, error := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))

	data := models.RpmData{
		Count: count,
		Rpm:   []models.RpmItem{},
	}

	for i := 0; i < count; i++ {
		item, err := t.GetRpmItem(bufferpack)

		if err == nil {
			data.Rpm = append(data.Rpm, item)
		}
	}

	return data, error
}

func (t *dataType) GetAlarmData(bufferpack []byte) (models.AlarmData, error) {
	newAlarmFlag, error := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 1))
	alarmType, error := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 1))
	alarmDescription, error := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 2))
	alarmThreshold, error := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 2))

	return models.AlarmData{
		AlarmThreshold:   alarmThreshold,
		NewAlarmFlag:     newAlarmFlag,
		AlarmType:        alarmType,
		AlarmDescription: alarmDescription,
	}, error
}

func (t *dataType) GetGpsData(bufferpack []byte) (models.GpsData, error) {
	count, err := convert.FromByteHexToInt(t.buff.Slice(bufferpack, 1))

	data := models.GpsData{
		Count:    count,
		GpsItems: []models.GpsItem{},
	}

	for i := 0; i < count; i++ {
		gpsItem, _ := t.GetGpsItem(bufferpack)
		data.GpsItems = append(data.GpsItems, gpsItem)
	}

	return data, err
}

func (t *dataType) GetTLVData(bufferpack []byte) (models.TLV, error) {
	tag, err := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 2))
	length, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 2))
	value := t.buff.Slice(bufferpack, length)

	tlv := models.TLV{
		Tag:        tag,
		Length:     length,
		ValueArray: value,
	}

	return tlv, err
}

func normalizeCoordenate(coordenate float64, isNegative bool) float64 {
	newcoordenate := coordenate / float64(3600000)
	if !isNegative {
		newcoordenate = newcoordenate * -1
	}
	return newcoordenate
}

func normalizeDirection(direction float64) float64 {
	return direction / 10
}

func convertCmsToKm(value float64) float64 {
	return math.Floor((value/27.778)*100) / 100
}
