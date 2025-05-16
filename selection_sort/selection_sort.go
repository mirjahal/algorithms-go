package selectionsort

func selectionSort(numbers []int) []int {
	orderedNumbers := []int{}

	for range numbers {
		indexOfSmallestItem := findSmallest(numbers)
		orderedNumbers = append(orderedNumbers, numbers[indexOfSmallestItem])
		numbers = removeItem(numbers, indexOfSmallestItem)
	}

	return orderedNumbers
}

func findSmallest(numbers []int) int {
	smallest := numbers[0]
	indexOfSmallestItem := 0

	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	// https://go.dev/tour/moretypes/16
	for index, value := range numbers {
		if numbers[index] < smallest {
			smallest = value
			indexOfSmallestItem = index
		}
	}

	return indexOfSmallestItem
}

func removeItem(numbers []int, indexToRemove int) []int {
	return append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
}
