package linearsearch

func numberExistsInArray(numberToCheck int, numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if (numbers[i] == numberToCheck) {
			return true;
		}
	}

	return false;
}