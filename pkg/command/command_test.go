package command

import (
	"strconv"
	"testing"
)

var (
	loginCommand = []byte{64, 64, 199, 0, 4, 50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0, 16, 1, 235, 163, 18, 94, 210, 164, 18, 94, 1, 188, 0, 0, 0, 0, 0, 0, 23, 1, 0, 0, 3, 0, 0, 0, 4, 0, 3, 49, 0, 87, 0, 0, 0, 0, 1, 6, 1, 20, 1, 52, 5, 184, 181, 17, 5, 112, 60, 7, 10}
)

func TestLoginInterprete(t *testing.T) {

	parser := NewParser()
	t.Run("Should return true when command is login", func(t *testing.T) {
		buffer := loginCommand
		got := parser.IsLogin(buffer)
		want := true
		if got != want {
			t.Errorf("got %q want %q", strconv.FormatBool(got), strconv.FormatBool(want))
		}
	})

	t.Run("Shouldnt return true when command isant login", func(t *testing.T) {
		buffer := []byte{48, 48, 49, 49, 48, 48, 48, 68, 56, 49, 53, 53, 49, 49, 52, 57, 54, 52, 48, 52, 54, 50, 70, 52, 48, 48, 48, 56, 48, 48, 50, 65, 48, 48, 52, 55, 48, 48, 52, 51, 48, 48, 53, 50, 48, 48, 51, 53, 48, 48, 51, 49, 48, 48, 51, 57, 48, 48, 51, 53, 48, 48, 50, 99, 48, 48, 50, 48, 48, 48, 54, 57, 48, 48, 54, 55, 48, 48, 54, 101, 48, 48, 54, 57, 48, 48, 55, 52, 48, 48, 54, 57, 48, 48, 54, 102, 48, 48, 54, 101, 48, 48, 50, 48, 48, 48, 54, 102, 48, 48, 54, 101, 48, 48, 50, 49, 26, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 64, 199, 0, 4, 50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0, 16, 1, 235, 163, 18, 94, 88, 164, 18, 94, 1, 188, 0, 0, 0, 0, 0, 0, 23, 1, 0, 0, 2, 0, 0, 0, 4, 0, 3, 61, 0, 87, 0, 0, 0, 0, 1, 6, 1, 20, 1, 52, 5, 184, 181, 17, 5, 112, 60, 7, 10}
		got := parser.IsLogin(buffer)
		want := false
		if got != want {
			t.Errorf("got %q want %q", strconv.FormatBool(got), strconv.FormatBool(want))
		}
	})
}

func TestExtractCommand(t *testing.T) {
	parser := NewParser()
	t.Run("Should extract device id", func(t *testing.T) {
		buffer := loginCommand
		got, err := parser.GetHexDeviceID(buffer)
		want := "0x3231334744503230313830323234323100000000"

		if got != want && err == nil {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should extract head", func(t *testing.T) {
		buffer := loginCommand
		got, err := parser.GetHead(buffer)
		want := "0x4040"

		if got != want && err == nil {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should extract version", func(t *testing.T) {
		buffer := loginCommand
		got, err := parser.GetVersion(buffer)
		want := "0x04"

		if got != want && err == nil {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should extract end", func(t *testing.T) {
		got, err := parser.GetEnd()
		want := "0x0d0a"

		if got != want && err == nil {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should get hex ip address", func(t *testing.T) {
		got := parser.GetIPHex()
		want := "0xFFFFFFFF"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
