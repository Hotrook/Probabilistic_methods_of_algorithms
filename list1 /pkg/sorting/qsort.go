package sorting

type Qsort struct {
	counter int
}

func (qsort *Qsort) count() {
	qsort.counter++
}

func (qsort *Qsort) init() {
	qsort.counter = 0
}

func (qsort *Qsort) sort(arr []int) {
	qsort.sortArr(arr, 0, len(arr)-1)
}

func (qsort *Qsort) sortArr(arr []int, start, stop int) {
	if start >= stop {
		return
	}

	pivot := arr[start]
	index := start

	for i := start + 1; i <= stop; i++ {
		if arr[i] < pivot {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}

	arr[start], arr[index] = arr[index], arr[start]

	qsort.sortArr(arr, start, index-1)
	qsort.sortArr(arr, index+1, stop)
}
