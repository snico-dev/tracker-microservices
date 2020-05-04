package timeconvert

import (
	"testing"
)

func TestDate(t *testing.T) {
	t.Run("Should conv date to hex", func(t *testing.T) {
		unixDate := NewUnixDateStub(int64(1581175201))
		dateConvert := NewTimeConvert(unixDate)
		got := dateConvert.GetHexUTC()
		want := "A1D13E5E"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
