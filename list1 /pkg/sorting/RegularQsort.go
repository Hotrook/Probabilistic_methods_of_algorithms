package sorting

type RegularQsort struct {
	swapCounter       int
	comparisonCounter int
}

func (qsort *RegularQsort) Sort(arr []int) (int, int) {
	qsort.init()
	qsort.sortArr(arr, 0, len(arr)-1)
	return qsort.swapCounter, qsort.comparisonCounter
}

func (qsort *RegularQsort) sortArr(arr []int, start, stop int) {
	if qsort.compare(stop, start) {
		return
	}

	pivot := arr[start]
	pivotIndex := start

	for i := start + 1; i <= stop; i++ {
		if qsort.compare(arr[i], pivot) {
			pivotIndex++
			qsort.swap(&arr[pivotIndex], &arr[i])
		}
	}

	qsort.swap(&arr[start], &arr[pivotIndex])

	qsort.sortArr(arr, start, pivotIndex-1)
	qsort.sortArr(arr, pivotIndex+1, stop)
}

func (qsort *RegularQsort) swap(a, b *int) {
	qsort.swapCount()
	*a, *b = *b, *a
}

func (qsort *RegularQsort) init() {
	qsort.swapCounter = 0
	qsort.comparisonCounter = 0
}

func (qsort *RegularQsort) compare(a, b int) bool {
	qsort.comparisonCounter++
	return a <= b
}

func (qsort *RegularQsort) compareLess(a, b int) bool {
	qsort.comparisonCounter++
	return a < b
}

func (qsort *RegularQsort) swapCount() {
	qsort.swapCounter++
}
