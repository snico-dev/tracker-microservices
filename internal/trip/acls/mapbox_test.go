package acls

import "testing"

func TestMapBox(t *testing.T) {

	t.Run("Should get address", func(t *testing.T) {
		mapboxACL := NewMapBoxACL()

		got, err := mapboxACL.GetAddressName(-23.710731666666668, -46.84625833333333)
		want := "Rua Itajobí Itapecerica Da Serra - São Paulo, 06850, Brazil"

		if err != nil || got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
