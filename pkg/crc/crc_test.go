package crc

import (
	"testing"
)

func TestCrc(t *testing.T) {
	t.Run("Should make crc", func(t *testing.T) {
		buffer := []byte{64, 64, 41, 0, 4, 49, 48, 48, 49, 49, 49, 50, 53, 50, 57, 57, 56, 55, 0, 0, 0, 0, 0, 0, 0, 144, 1, 255, 255, 255, 255, 0, 0, 193, 222, 121, 82}
		got, err := Make(buffer, len(buffer))
		want := "0xa5dd"

		if got != want || err != nil {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
