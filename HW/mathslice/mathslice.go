package mathslice

func SumSlice(s []int) int {
	res := 0
	for _, val := range s {
		res += val
	}
	return res
}

func MapSlice(s []int, op func(int) int) {
	for i, val := range s {
		s[i] = op(val)
	}
}

func FoldSlice(s []int, op func(int, int) int, init int) int {
	if len(s) == 0 {
		return 0
	}
	res := init
	for _, val := range s {
		res = op(res, val)
	}
	return res
}
