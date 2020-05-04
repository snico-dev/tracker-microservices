package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("Should reverse a string", func(t *testing.T) {
		got := DescInGroup("3118")
		want := "1831"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
