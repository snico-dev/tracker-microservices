package services

import (
	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/buffer"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/convert"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
)

//IReportType interface
type IReportType interface {
	GetAlarmReport(bufferpack []byte) (models.AlarmReport, error)
	GetGpsReport(bufferpack []byte) (models.GpsReport, error)
	GetFirstMessagePack(bufferpack []byte) (models.Message, error)
	GetSecondMessagePack(bufferpack []byte, model models.Message) (models.Message, error)
	GetQuery(bufferpack []byte) (models.Query, error)
}

type reportType struct {
	buff     buffer.IBuffer
	unixDate timeconvert.UnixDate
	dataType IDataType
}

//NewReport contructor
func NewReport() IReportType {
	buff := buffer.NewBuffer()
	unixDate := timeconvert.NewUnixDate()
	return &reportType{
		buff:     buff,
		unixDate: unixDate,
		dataType: NewDataType(buff, unixDate),
	}
}

func (t *reportType) GetAlarmReport(bufferpack []byte) (models.AlarmReport, error) {
	alarmSeq, error := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 4))
	stateData, error := t.dataType.GetStateData(bufferpack)
	gpsData, error := t.dataType.GetGpsData(bufferpack)

	alarmCount, error := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 1))

	data := models.AlarmReport{
		AlarmSeq:   alarmSeq,
		StateData:  stateData,
		GpsData:    gpsData,
		AlarmCount: alarmCount,
		AlarmArray: []models.AlarmData{},
	}

	for i := 0; i < alarmCount; i++ {
		alarmData, error := t.dataType.GetAlarmData(bufferpack)

		if error == nil {
			data.AlarmArray = append(data.AlarmArray, alarmData)
		}
	}

	return data, error
}

func (t *reportType) GetGpsReport(bufferpack []byte) (models.GpsReport, error) {
	flag, error := convert.FromByteToHex(t.buff.Slice(bufferpack, 1))
	statData, error := t.dataType.GetStateData(bufferpack)
	gpsData, error := t.dataType.GetGpsData(bufferpack)
	rpmData, error := t.dataType.GetRpmData(bufferpack)

	return models.GpsReport{
		Flag:      flag,
		StateData: statData,
		GpsData:   gpsData,
		RpmData:   rpmData,
	}, error
}

func (t *reportType) GetFirstMessagePack(bufferpack []byte) (models.Message, error) {
	head, error := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 2))
	length, error := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 2))
	version, error := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 1))
	deviceID, error := convert.FromByteHexToASCII(t.buff.Slice(bufferpack, 20))
	protocolID, error := convert.FromByteToHex(t.buff.Slice(bufferpack, 2))

	return models.Message{
		Head:       head,
		Length:     length,
		Version:    version,
		DeviceID:   deviceID,
		ProtocolID: protocolID,
	}, error
}

func (t *reportType) GetSecondMessagePack(bufferpack []byte, model models.Message) (models.Message, error) {
	crc, error := convert.FromByteToHex(t.buff.Slice(bufferpack, 2))
	tail, _ := convert.FromByteToHex(t.buff.Slice(bufferpack, 2))

	model.CRC = crc
	model.Tail = tail
	return model, error
}

func (t *reportType) GetQuery(bufferpack []byte) (models.Query, error) {
	cmdSeq, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 2))
	respCount, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 1))
	respIndex, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 1))
	failCount, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 1))
	failTagArray := []string{}

	for i := 0; i < failCount; i++ {
		failTag, err := convert.FromByteToReverseHex(t.buff.Slice(bufferpack, 2))
		if err == nil {
			failTagArray = append(failTagArray, failTag)
		}
	}

	successCount, err := convert.FromByteHexReverseToInt(t.buff.Slice(bufferpack, 1))
	successTLVArray := []models.TLV{}

	for i := 0; i < successCount; i++ {
		tlv, err := t.dataType.GetTLVData(bufferpack)
		if err == nil {
			successTLVArray = append(successTLVArray, tlv)
		}
	}

	return models.Query{
		CmdSeq:          cmdSeq,
		RespCount:       respCount,
		FailCount:       failCount,
		FailTagArray:    failTagArray,
		RespIndex:       respIndex,
		SuccessCount:    successCount,
		SuccessTLVArray: successTLVArray,
	}, err
}
