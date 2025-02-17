package interfaces

import "testing"

func TestWeight(t *testing.T) {
	checkWeight := func(t testing.TB, animal Animal, want float64) {
		t.Helper()
		got := animal.Weight()
		if got != want {
			t.Errorf("got %.2f, weight %.2f", got, want)
		}
	}

	t.Run("cow", func(t *testing.T) {
		cow := Cow{30}
		want := 30 * 9.8
		checkWeight(t, &cow, want)
		checkWeight(t, &cow, CalculateWeight(&cow))
	})

	t.Run("whale", func(t *testing.T) {
		whale := Whale{120, 52}
		want := 120*9.8 - 52
		checkWeight(t, &whale, want)
	})
}
