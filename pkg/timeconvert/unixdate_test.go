package timeconvert

import (
	"testing"
	"time"
)

func TestUnixDate(t *testing.T) {

	t.Run("Should get date from seconds", func(t *testing.T) {
		unixdate := NewUnixDate()
		got := unixdate.GetDateFromSeconds(1584322719)
		want := time.Date(2020, time.March, 16, 1, 38, 39, 0, GetLocation())

		if got.Equal(want) == false {
			t.Errorf("data not equals")
		}
	})
}
