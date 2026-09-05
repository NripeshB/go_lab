package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing empty slices in tail", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{}
		got := SumAllTails(nums1, nums2)
		want := []int{0, 0}
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d , given %v, %v", got, want, nums1, nums2)
		}
	})

	t.Run("Test sum of tails", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{1, 3}
		got := SumAllTails(nums1, nums2)
		want := []int{5, 3}
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d , given %v, %v", got, want, nums1, nums2)
		}
	})
	t.Run("Sum All given arrays", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{1, 3}
		got := SumAll(nums1, nums2)
		want := []int{6, 4}
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d , given %v, %v", got, want, nums1, nums2)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
