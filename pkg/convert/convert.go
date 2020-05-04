package convert

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//FromHexToASCII return ascii parsed from hex
func FromHexToASCII(hexstr string) (string, error) {
	if len(hexstr) == 0 || hexstr == "" {
		return "", errors.New("argument can't be null or empty")
	}

	hexstr = strings.Replace(hexstr, "0x", "", 1)

	bs, err := hex.DecodeString(hexstr)

	if err != nil {
		return "", errors.New("can´t parse command")
	}

	newhexstr := string(bs)
	newhexstr = strings.ReplaceAll(newhexstr, "\x00", "")

	return newhexstr, nil
}

//FromByteHexToASCII from hex bytes to ascii
func FromByteHexToASCII(bytes []byte) (string, error) {
	hex, error := FromByteToHex(bytes)

	if error != nil {
		return "", error
	}

	return FromHexToASCII(hex)
}

func fromByteToHex(bytes []byte, rv bool) (string, error) {
	if len(bytes) == 0 {
		return "", errors.New("argument can't be null or empty")
	}

	if rv == true {
		bytes = reverse(bytes)
	}

	return "0x" + hex.EncodeToString(bytes), nil
}

//FromByteToHex return ascii parsed from hex
func FromByteToHex(bytes []byte) (string, error) {
	return fromByteToHex(bytes, false)
}

//FromByteHexToInt convert byte hex to int
func FromByteHexToInt(bytes []byte) (int, error) {
	hex, err := fromByteToHex(bytes, false)
	if err != nil {
		return -1, err
	}

	return FromHexToInt(hex)
}

//FromByteHexReverseToInt convert byte hex to int
func FromByteHexReverseToInt(bytes []byte) (int, error) {
	hex, err := fromByteToHex(bytes, true)
	if err != nil {
		return -1, err
	}

	return FromHexToInt(hex)
}

//FromByteHexToFloat64 convert byte hex to float64
func FromByteHexToFloat64(bytes []byte) (float64, error) {
	hex, err := fromByteToHex(bytes, false)
	if err != nil {
		return -1, err
	}

	return FromHexToFloat64(hex)
}

//FromByteHexReverseToFloat64 convert byte hex to float64
func FromByteHexReverseToFloat64(bytes []byte) (float64, error) {
	hex, err := fromByteToHex(bytes, true)
	if err != nil {
		return -1, err
	}

	return FromHexToFloat64(hex)
}

//FromByteToReverseHex return hex parsed reversed
func FromByteToReverseHex(bytes []byte) (string, error) {
	return fromByteToHex(bytes, true)
}

func reverse(bt []byte) []byte {
	for i, j := 0, len(bt)-1; i < j; i, j = i+1, j-1 {
		bt[i], bt[j] = bt[j], bt[i]
	}
	return bt
}

//FromByteToHexWithOutPrefix return hex
func FromByteToHexWithOutPrefix(bytes []byte) (string, error) {
	if len(bytes) == 0 {
		return "", errors.New("argument can't be null or empty")
	}

	return hex.EncodeToString(bytes), nil
}

// FromDecimalToHex return hex number
func FromDecimalToHex(n int64) string {
	return fmt.Sprintf("%02x", n)
}

// FromLongToHex return hex number
func FromLongToHex(n uint64) string {
	return fmt.Sprintf("%02x", n)
}

// FromUnit16ToByte return hex number
func FromUnit16ToByte(n uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, n)
	return b
}

// FromIntToByte return hex number
func FromIntToByte(n int) byte {
	return byte(n)
}

//FromHexToInt convert from hexdecimal to uint32
func FromHexToInt(value string) (int, error) {
	newvalue := strings.Replace(value, "0x", "", 1)
	n, err := strconv.ParseUint(newvalue, 16, 32)

	if err != nil {
		return 0, errors.New("error when try parse value")
	}

	return int(n), nil
}

//FromHexToInt64 convert from hexdecimal to uint16
func FromHexToInt64(value string) (int64, error) {
	newvalue := strings.Replace(value, "0x", "", 1)
	number, err := strconv.ParseInt(newvalue, 16, 64)

	if err != nil {
		return 0, errors.New("error when try parse value")
	}

	return number, nil
}

//FromHexToFloat64 convert from hexdecimal to float64
func FromHexToFloat64(value string) (float64, error) {
	newvalue := strings.Replace(value, "0x", "", 1)
	number, err := strconv.ParseInt(newvalue, 16, 32)

	if err != nil {
		return 0, errors.New("error when try parse value")
	}

	return float64(number), nil
}

//FromHexToBitArray from hex to bit array
func FromHexToBitArray(hex string) ([]bool, error) {
	number, err := FromHexToInt64(hex)

	if err != nil {
		return nil, err
	}

	strnumber := strconv.FormatInt(number, 2)

	bits := []bool{}
	for _, r := range strnumber {
		c := string(r)
		val, _ := strconv.ParseUint(c, 10, 16)
		bits = append(bits, !(uint16(val) == 0))
	}

	return bits, nil
}
