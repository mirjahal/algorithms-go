package selectionsort

import "testing"

func TestArrayLenAfterSelectionSortSuccessfully(t *testing.T) {
	numbers := [5]int{1, 9, 7, 5, 3}

	orderedNumbers := selectionSort(numbers[:])

	if len(orderedNumbers) != 5 {
		t.Errorf(`The array len returned was %v, but the len 5 was expected`, len(orderedNumbers))
	}
}

func TestSortAfterSelectionSortExecutedSuccessfully(t *testing.T) {
	numbers := [5]int{1, 9, 7, 5, 3}

	orderedNumbers := selectionSort(numbers[:])

	if orderedNumbers[0] != 1 {
		t.Errorf(`The value returned by index 0 was %v, but the value 1 was expected`, orderedNumbers[0])
	}

	if orderedNumbers[1] != 3 {
		t.Errorf(`The value returned by index 1 was %v, but the value 3 was expected`, orderedNumbers[1])
	}

	if orderedNumbers[2] != 5 {
		t.Errorf(`The value returned by index 2 was %v, but the value 5 was expected`, orderedNumbers[2])
	}

	if orderedNumbers[3] != 7 {
		t.Errorf(`The value returned by index 3 was %v, but the value 7 was expected`, orderedNumbers[3])
	}

	if orderedNumbers[4] != 9 {
		t.Errorf(`The value returned by index 4 was %v, but the value 9 was expected`, orderedNumbers[4])
	}
}
