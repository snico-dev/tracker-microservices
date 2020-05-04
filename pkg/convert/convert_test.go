package convert

import (
	"reflect"
	"testing"
)

func TestParseHexToAscii(t *testing.T) {
	t.Run("Should parse hex to ascii", func(t *testing.T) {
		strhex := "0x3231334744503230313830323234323100000000"
		got, err := FromHexToASCII(strhex)
		want := "213GDP2018022421"

		if got != want && err != nil {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should thrown error when pass empty hex to parse", func(t *testing.T) {
		got, err := FromHexToASCII("")

		if err == nil || len(got) > 0 {
			t.Errorf("got success when should be error")
		}
	})

}

func TestParseByteToHex(t *testing.T) {

	t.Run("Should parse bytes to hex", func(t *testing.T) {
		buffer := []byte{50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0}

		got, err := FromByteToHex(buffer)
		want := "0x3231334744503230313830323234323100000000"

		if got != want && err != nil {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should thrown error when pass empty bytes to parse", func(t *testing.T) {
		got, err := FromByteToHex(nil)

		if err == nil || len(got) > 0 {
			t.Errorf("got success when should be error")
		}
	})
}

func TestDecimalToHex(t *testing.T) {
	got := FromDecimalToHex(15)
	want := "0f"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHexToBitArray(t *testing.T) {
	got, err := FromHexToBitArray("0x168")
	want := []bool{true, false, true, true, false, true, false, false, false}

	if err != nil || reflect.DeepEqual(got, want) == false {
		t.Error("object not equals")
	}
}
