package buffer

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("Should slice buffer", func(t *testing.T) {
		bf := NewBuffer()

		bufferpack := []byte{0, 1, 2, 3}
		got := bf.Slice(bufferpack, 2)

		want := []byte{0, 1}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should get buffer slice size", func(t *testing.T) {
		got := NewBuffer()

		bufferpack := []byte{0, 1, 2, 3}
		got.Slice(bufferpack, 3)

		want := 3

		if got.Size() == want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Should slice buffer two times", func(t *testing.T) {
		bf := NewBuffer()

		bufferpack := []byte{0, 1, 2, 3}
		bf.Slice(bufferpack, 2)
		got := bf.Slice(bufferpack, 2)

		want := []byte{2, 3}

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
