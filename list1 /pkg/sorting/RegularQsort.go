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

func (qsort *RegularQsort) sortArr(arr []int, left, right int) {
	if right <= left {
		return
	}

	p := arr[right]
	i := left - 1
	j := right

	for {
		for {
			i++
			if !qsort.compareLess(arr[i], p) {
				break
			}
		}
		for {
			j--
			if j == -1 || !qsort.compareLess(p, arr[j]) {
				break
			}
		}
		if j > i {
			qsort.swap(&arr[i], &arr[j])
		}
		if !(j > i) {
			break
		}
	}
	qsort.swap(&arr[i], &arr[right])

	qsort.sortArr(arr, left, i-1)
	qsort.sortArr(arr, i+1, right)
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
