package binarysearch

import "testing"

func TestBinarySearchFindIndexSuccessfully(t *testing.T) {
	numbers := [5]int{1, 3, 5, 7, 9}
	guess := 7
	expected := 3

	index := binarySearch(numbers[:], guess)

	if index != expected {
		t.Errorf(`The index returned was %v, but %v was expected`, index, expected)
	}
}

func TestBinarySearchIndexNotFoundInArray(t *testing.T) {
	numbers := [9]int{1, 3, 5, 7, 9, 12, 15, 17, 19}
	guess := 25
	expected := -1

	index := binarySearch(numbers[:], guess)

	if index != expected {
		t.Errorf(`The index returned was %v, but %v was expected`, index, expected)
	}
}

func TestBinarySearchIndexNotFoundInEmptyArray(t *testing.T) {
	numbers := [9]int{}
	guess := 25
	expected := -1

	index := binarySearch(numbers[:], guess)

	if index != expected {
		t.Errorf(`The index returned was %v, but %v was expected`, index, expected)
	}
}
