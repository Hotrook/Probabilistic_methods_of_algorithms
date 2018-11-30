package permutation

func CalculateFixedPoints(permutation []int) int {
	counter := 0

	for i := 0; i < len(permutation); i++ {
		if permutation[i] == i {
			counter++
		}
	}

	return counter
}

func CalculateCycles(permutation []int) int {
	cyclesNumber := 0
	size := len(permutation)
	visited := make([]bool, size)

	for i := 0; i < len(permutation); i++ {
		if !visited[i] {
			position := i
			for !visited[position] {
				visited[position] = true
				position = permutation[position]
			}
			cyclesNumber++
		}
	}

	return cyclesNumber
}

func CalculateRecords(permutation []int) int {
	recordsNumber := 0
	biggestSeenNumber := -1

	for i := 0; i < len(permutation); i++ {
		if permutation[i] > biggestSeenNumber {
			recordsNumber++
			biggestSeenNumber = permutation[i]
		}
	}

	return recordsNumber
}
