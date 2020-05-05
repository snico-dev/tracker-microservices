package command

import (
	"errors"
	"strings"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/convert"
)

//Parser interface of command parser
type Parser interface {
	IsAlarmReport(command []byte) bool
	IsGpsReport(command []byte) bool
	IsQueryReport(command []byte) bool
	IsLogin(command []byte) bool
	IsPidReport(command []byte) bool
	GetHexDeviceID(command []byte) (string, error)
	GetHead(command []byte) (string, error)
	GetVersion(command []byte) (string, error)
	GetEnd() (string, error)
	GetIPHex() string
	GetPort() string
}

type parser struct {
}

//NewParser constructor
func NewParser() Parser {
	return &parser{}
}

//IsQueryReport return true if is query report
func (c parser) IsQueryReport(command []byte) bool {
	if len(command) == 0 {
		return false
	}

	loginByte := command[25:27]
	hexReportNumber, err := convert.FromByteToHex(loginByte)

	if (strings.ToUpper(hexReportNumber) == "0xA002") && err == nil {
		return true
	}

	return false
}

//IsAlarmReport return if command is alarm
func (c parser) IsPidReport(command []byte) bool {
	if len(command) == 0 {
		return false
	}

	loginByte := command[25:27]
	hexReportNumber, err := convert.FromByteToHex(loginByte)

	if (hexReportNumber == "0x4002" || hexReportNumber == "0x1002") && err == nil {
		return true
	}

	return false
}

//IsAlarmReport return if command is alarm
func (c parser) IsAlarmReport(command []byte) bool {
	if len(command) == 0 {
		return false
	}

	loginByte := command[25:27]
	hexReportNumber, err := convert.FromByteToHex(loginByte)

	if hexReportNumber == "0x4007" && err == nil {
		return true
	}

	return false
}

//IsGpsReport return if command is gps
func (c parser) IsGpsReport(command []byte) bool {
	if len(command) == 0 {
		return false
	}

	loginByte := command[25:27]
	hexReportNumber, err := convert.FromByteToHex(loginByte)

	if hexReportNumber == "0x4001" && err == nil {
		return true
	}

	return false
}

//IsLogin return if command is login
func (c parser) IsLogin(command []byte) bool {
	if len(command) == 0 {
		return false
	}

	loginByte := command[25:27]
	hexReportNumber, err := convert.FromByteToHex(loginByte)

	if hexReportNumber == "0x1001" && err == nil {
		return true
	}

	return false
}

// GetHexDeviceID return hex device id convert.from byte
func (c parser) GetHexDeviceID(command []byte) (string, error) {
	return extractByteRangeAndReturnHex(command, 5, 25, "Identificador do device não encontrado")
}

// GetHead return start info
func (c parser) GetHead(command []byte) (string, error) {
	return extractByteRangeAndReturnHex(command, 0, 2, "Posição start não encontrada")
}

// GetVersion return version
func (c parser) GetVersion(command []byte) (string, error) {
	return extractByteRangeAndReturnHex(command, 4, 5, "Versão não encontrada")
}

// GetEnd return end
func (c parser) GetEnd() (string, error) {
	return "0x0d0a", nil
}

// GetIPHex return hex ip
func (c parser) GetIPHex() string {
	return "0xFFFFFFFF"
}

//GetPort return command port
func (c parser) GetPort() string {
	return "0x0000"
}

func extractByteRangeAndReturnHex(command []byte, startAt int, endAt int, message string) (string, error) {
	if len(command) == 0 || endAt > len(command) {
		return "", errors.New("buffer não pode ser vazio ou ser acessado por um index maior")
	}

	dataByte := command[startAt:endAt]

	data, err := convert.FromByteToHex(dataByte)

	if err != nil {
		return "", errors.New(message)
	}

	return data, nil
}
