package report

import (
	"encoding/hex"
	"strings"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/command"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/utils"
)

// RetriveLogin return login package response
func RetriveLogin(command []byte, dataConv timeconvert.TimeConvert, parse command.Parser) ([]byte, error) {

	startHex, err := parse.GetHead(command)
	versionHex, err := parse.GetVersion(command)
	deviceIDHex, err := parse.GetHexDeviceID(command)
	endHex, err := parse.GetEnd()
	reTryServerHex := parse.GetIPHex()
	portHex := parse.GetPort()
	packageLengthHex := "0x2900"
	utcHex := "0x" + dataConv.GetHexUTC()
	orderHex := "0x9001"

	length, err := hex.DecodeString(clearPrefix(packageLengthHex))
	head, err := hex.DecodeString(clearPrefix(startHex))
	version, err := hex.DecodeString(clearPrefix(versionHex))
	deviceID, err := hex.DecodeString(clearPrefix(deviceIDHex))
	order, err := hex.DecodeString(clearPrefix(orderHex))
	end, err := hex.DecodeString(clearPrefix(endHex))
	can1, err := hex.DecodeString(clearPrefix(reTryServerHex))
	can2, err := hex.DecodeString(clearPrefix(portHex))
	utc, err := hex.DecodeString(clearPrefix(utcHex))

	psrc := make([]byte, 0)
	psrc = append(psrc, head...)
	psrc = append(psrc, length...)
	psrc = append(psrc, version...)
	psrc = append(psrc, deviceID...)
	psrc = append(psrc, order...)

	psrc = append(psrc, can1...)
	psrc = append(psrc, can2...)
	psrc = append(psrc, utc...)

	finalByte := make([]byte, 0)
	crcBytes, err := getCrcBytes(psrc, true)
	finalByte = append(finalByte, psrc...)
	finalByte = append(finalByte, crcBytes...)
	finalByte = append(finalByte, end...)

	return finalByte, err
}

//QueryCommand ask to the terminal info in vehicle's ecu
func QueryCommand(command []byte, lengthHex string, sequenceHex string, queryNumberHex string, parametersHex string, parse command.Parser) ([]byte, error) {
	startHex, err := parse.GetHead(command)
	versionHex, err := parse.GetVersion(command)
	deviceIDHex, err := parse.GetHexDeviceID(command)
	orderHex := "0x2002"
	endHex, err := parse.GetEnd()

	head, err := hex.DecodeString(clearPrefix(startHex))
	length, err := hex.DecodeString(clearPrefix(lengthHex))
	version, err := hex.DecodeString(clearPrefix(versionHex))
	deviceID, err := hex.DecodeString(clearPrefix(deviceIDHex))
	order, err := hex.DecodeString(clearPrefix(orderHex))
	sequence, err := hex.DecodeString(clearPrefix(sequenceHex))
	queryNumber, err := hex.DecodeString(clearPrefix(queryNumberHex))
	parameters, err := hex.DecodeString(clearPrefix(parametersHex))
	end, err := hex.DecodeString(clearPrefix(endHex))

	psrc := make([]byte, 0)
	psrc = append(psrc, head...)
	psrc = append(psrc, length...)
	psrc = append(psrc, version...)
	psrc = append(psrc, deviceID...)
	psrc = append(psrc, order...)
	psrc = append(psrc, sequence...)
	psrc = append(psrc, queryNumber...)
	psrc = append(psrc, parameters...)

	finalByte := make([]byte, 0)
	crcBytes, err := getCrcBytes(psrc, false)
	finalByte = append(finalByte, psrc...)
	finalByte = append(finalByte, crcBytes...)
	finalByte = append(finalByte, end...)

	return finalByte, err
}

func clearPrefix(hex string) string {
	return strings.Replace(hex, "0x", "", 1)
}

func getCrcBytes(psrc []byte, desc bool) ([]byte, error) {
	crc, err := crc.Make(psrc, len(psrc))

	if err != nil {
		return nil, err
	}
	crcWithOutPrefix := clearPrefix(crc)

	if desc {
		return hex.DecodeString(utils.DescInGroup(crcWithOutPrefix))
	}

	return hex.DecodeString(crcWithOutPrefix)
}
