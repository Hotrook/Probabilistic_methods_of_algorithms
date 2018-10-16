package sorting

type Qsort struct {
	swapCounter       int
	comparisonCounter int
}

func (qsort *Qsort) sort(arr []int) {
	qsort.init()
	qsort.sortArr(arr, 0, len(arr)-1)
}

func (qsort *Qsort) sortArr(arr []int, start, stop int) {
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

func (qsort *Qsort) swap(a, b *int) {
	qsort.swapCounter++
	*a, *b = *b, *a
}

func (qsort *Qsort) init() {
	qsort.swapCounter = 0
	qsort.comparisonCounter = 0
}

func (qsort *Qsort) compare(a, b int) bool {
	qsort.comparisonCounter++
	return a <= b
}
