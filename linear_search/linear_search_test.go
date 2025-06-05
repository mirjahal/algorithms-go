package linearsearch

import "testing"

func TestLinearSearchFindNumberSuccessfully(t *testing.T) {
	numbers := []int{1, 3, 5, 7, 9}
	number := 7

	if !numberExistsInArray(number, numbers) {
		t.Errorf(`The number does not exists in array numbers, but it should exist`)
	}
}

func TestLinearSearchNumberNotFound(t *testing.T) {
	numbers := []int{1, 3, 5, 7, 9}
	number := 41

	if numberExistsInArray(number, numbers) {
		t.Errorf(`The number exists in the numbers array, but it shouldn't exist`)
	}
}

func TestLinearSearchNumberNotFoundEmptyArray(t *testing.T) {
	numbers := []int{}
	number := 1

	if numberExistsInArray(number, numbers) {
		t.Errorf(`The number exists in the numbers array, but it shouldn't exist`)
	}
}