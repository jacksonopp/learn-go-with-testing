package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectNumber := func(t testing.TB, got, want int, given []int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, given)
		}
	}
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectNumber(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	assertAllTails := func(got, want []int, t testing.TB) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("it works", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertAllTails(got, want, t)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertAllTails(got, want, t)
	})

}
