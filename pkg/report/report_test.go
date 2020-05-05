package report

import (
	"strings"
	"testing"

	"github.com/NicolasDeveloper/tracker-udp-server/apps/shared/convert"
	"github.com/NicolasDeveloper/tracker-udp-server/apps/shared/timeconvert"
)

func TestResponseCommand(t *testing.T) {
	var buffer = []byte{64, 64, 127, 0, 4, 49, 48, 48, 49, 49, 49, 50, 53, 50, 57, 57, 56, 55, 0, 0, 0, 0, 0, 0, 0, 16, 1, 193, 240, 105, 82, 253, 240, 105, 82, 156, 145, 17, 0, 0, 0, 0, 0, 105, 131, 0, 0, 12, 0, 0, 0, 0, 0, 3, 100, 1, 1, 76, 0, 3, 0, 1, 25, 10, 13, 4, 18, 26, 20, 128, 214, 4, 136, 197, 114, 24, 0, 0, 0, 0, 175, 73, 68, 68, 95, 50, 49, 54, 71, 48, 50, 95, 83, 32, 86, 49, 46, 50, 46, 49, 0, 73, 68, 68, 95, 50, 49, 54, 71, 48, 50, 95, 72, 32, 86, 49, 46, 50, 46, 49, 0, 0, 0, 223, 100, 13, 10}
	// var buffer = []byte{64, 64, 208, 0, 4, 50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0, 16, 1, 121, 100, 64, 94, 118, 102, 64, 94, 26, 242, 0, 0, 167, 4, 0, 0, 26, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 87, 0, 0, 0, 0, 0, 73, 68, 68, 95, 50, 49, 51, 87, 48, 49, 95, 83, 32, 86, 50, 46, 50, 46, 48, 0, 73, 68, 68, 95, 50, 49, 51, 87, 48, 49, 95, 72, 32, 86, 50, 46, 50, 46, 48, 0, 50, 0, 1, 24, 2, 24, 1, 26, 1, 27, 1, 30, 1, 31, 2, 31, 3, 31, 4, 31, 5, 31, 6, 31, 7, 31, 1, 33, 2, 33, 1, 16, 2, 16, 3, 16, 4, 16, 5, 16, 6, 16, 7, 16, 8, 16, 9, 16, 10}
	t.Run("Should retrive login command", func(t *testing.T) {
		parseStub := NewParseStub("0x4040", "0x3130303131313235323939383700000000000000", "0x04", "0xFFFFFFFF", "0x0000", "0x0D0A")
		loginCommand, err := RetriveLogin(buffer, timeconvert.NewTimeConvert(timeconvert.NewUnixDateStub(1383718593)), parseStub)
		// loginCommand, err := RetriveLogin(buffer, timeconvert.NewTimeConvert(timeconvert.NewUnixDate()), command.NewParser())
		got, err := convert.FromByteToHexWithOutPrefix(loginCommand)

		if strings.ToUpper(got) != "404029000431303031313132353239393837000000000000009001FFFFFFFF0000C1DE7952DDA50D0A" && err == nil {
			t.Error(strings.ToUpper(got))
		}
	})
}
