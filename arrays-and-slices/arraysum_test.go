package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := ArraySum(numbers)

		if got != want {
			t.Errorf("got %d, want %d with input %v", got, want, numbers)
		}
	})

	t.Run("collection of numbers of any length", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		expected := 6
		got := ArraySum(numbers)

		if got != expected {
			t.Errorf("got %d, want %d, with input %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4, 5}
	slice3 := []int{1, 4}

	// expect to return slice of sum of each slice
	want := []int{3, 12, 5}
	got := SumAll(slice1, slice2, slice3)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v with input: %v, %v, %v", got, want, slice1, slice2, slice3)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("some of regular slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{2, 9}

		want := []int{5, 9}
		got := SumAllTails(slice1, slice2)

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{1, 2, 3}

		want := []int{0, 5}
		got := SumAllTails(slice1, slice2)

		checkSums(t, got, want)
	})
}
