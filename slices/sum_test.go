package slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("expected %d but got %d given, %v", expected, sum, numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	}

	t.Run("Sum the tails of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{5, 6, 7, 8})
		want := []int{14, 21}
		checkSums(t, got, want)
	})

	t.Run("Safely sum the tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}
		checkSums(t, got, want)
	})
}
