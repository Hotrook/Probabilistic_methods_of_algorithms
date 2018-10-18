package sorting

type args struct {
	arr []int
}

var tests = []struct {
	name     string
	args     args
	expected []int
}{
	{
		"Regular RegularQsort",
		args{[]int{3, 2, 1}},
		[]int{1, 2, 3},
	},
	{
		"Already sorted",
		args{[]int{1, 2, 3}},
		[]int{1, 2, 3},
	},
	{
		"Regular RegularQsort",
		args{[]int{3, 4, 7, 5, 6, 9, 8, 2, 1}},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		"Error sequence",
		args{[]int{2, 3, 0, 1}},
		[]int{0, 1, 2, 3},
	},
}
