package sorting

type DoublePivotQsort struct {
	Qsort
	swapCounter       int
	comparisonCounter int
}

func (qsort *DoublePivotQsort) sort(arr []int) {
	qsort.init()
	qsort.sortArr(arr, 0, len(arr)-1)
}

func (qsort *DoublePivotQsort) sortArr(arr []int, left, right int) {
	if qsort.compare(right, left) {
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
			if qsort.compare(j, i) {
				done = true
				break
			}
			if qsort.compare(arr[i], p) {
				arr[i1] = arr[i]
				i1++
				arr[i] = arr[i1]
			}
			i++
		}
		if done {
			break
		}

		j--
		for qsort.compare(p, arr[j]) {
			if qsort.compare(q, arr[j]) {
				arr[j1] = arr[j]
				j1--
				arr[j] = arr[j1]
			}
			if qsort.compare(j, i) {
				done = true
				break
			}
			j--
		}
		if done {
			break
		}
		arr[i1] = arr[j]
		arr[j1] = arr[i]
		i++
		j--
		arr[i] = arr[i1]
		arr[j] = arr[j1]
	}
	arr[i1] = p
	arr[j1] = q
	qsort.sortArr(arr, left, i1-1)
	qsort.sortArr(arr, i1+1, j1-1)
	qsort.sortArr(arr, j1+1, right)
}
