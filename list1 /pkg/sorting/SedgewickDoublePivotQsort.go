package sorting

type SedgewickDoublePivotQsort struct {
	RegularQsort
}

func (qsort *SedgewickDoublePivotQsort) Sort(arr []int) (int, int) {
	qsort.init()
	qsort.sortArr(arr, 0, len(arr)-1)
	return qsort.swapCounter, qsort.comparisonCounter
}

func (qsort *SedgewickDoublePivotQsort) sortArr(arr []int, left, right int) {
	if right <= left {
		return
	}

	p, q := arr[left], arr[right]
	i1, j1 := left, right
	i, j := left, right

	if qsort.compare(q, p) {
		qsort.swap(&q, &p)
	}
	done := false
	for {
		i++
		for qsort.compare(arr[i], q) {
			if j <= i {
				done = true
				break
			}
			if qsort.compareLess(arr[i], p) {
				arr[i1] = arr[i]
				qsort.swapCount()
				i1++
				arr[i] = arr[i1]
				qsort.swapCount()
			}
			i++
		}
		if done {
			break
		}

		j--
		for qsort.compare(p, arr[j]) {
			if qsort.compareLess(q, arr[j]) {
				arr[j1] = arr[j]
				qsort.swapCount()
				j1--
				arr[j] = arr[j1]
				qsort.swapCount()
			}
			if j <= i {
				done = true
				break
			}
			j--
		}
		if done {
			break
		}
		arr[i1] = arr[j]
		qsort.swapCount()
		arr[j1] = arr[i]
		qsort.swapCount()

		i1++
		j1--
		arr[i] = arr[i1]
		qsort.swapCount()
		arr[j] = arr[j1]
		qsort.swapCount()

	}

	arr[i1] = p
	arr[j1] = q

	qsort.sortArr(arr, left, i1-1)
	qsort.sortArr(arr, i1+1, j1-1)
	qsort.sortArr(arr, j1+1, right)
}
