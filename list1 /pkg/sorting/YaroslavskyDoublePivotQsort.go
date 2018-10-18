package sorting

type YaroslavskyDoublePivotQsort struct {
	RegularQsort
}

func (qsort *YaroslavskyDoublePivotQsort) Sort(arr []int) (int, int) {
	qsort.init()
	qsort.sortArr(arr, 0, len(arr)-1)
	return qsort.swapCounter, qsort.comparisonCounter
}

func (qsort *YaroslavskyDoublePivotQsort) sortArr(arr []int, left, right int) {
	if right <= left {
		return
	}

	p, q := arr[left], arr[right]
	l, g := left+1, right-1
	k := l

	if qsort.compare(q, p) {
		qsort.swap(&q, &p)
		qsort.swap(&arr[left], &arr[right])
	}

	for k <= g {
		if qsort.compareLess(arr[k], p) {
			qsort.swap(&arr[k], &arr[l])
			l++
		} else if qsort.compareLess(q, arr[k]) {
			for qsort.compareLess(q, arr[g]) && k < g {
				g--
			}
			qsort.swap(&arr[k], &arr[g])
			g--
			if qsort.compareLess(arr[k], p) {
				qsort.swap(&arr[k], &arr[l])
				l++
			}

		}
		k++
	}

	l--
	g++
	qsort.swap(&arr[left], &arr[l])
	qsort.swap(&arr[right], &arr[g])

	qsort.sortArr(arr, left, l-1)
	qsort.sortArr(arr, l+1, g-1)
	qsort.sortArr(arr, g+1, right)
}
