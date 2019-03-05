package easy

func index(idx int) int {
	if idx == 0 || idx == 1 {
		return idx
	}
	if idx == 2 {
		return 1
	}

	return index(idx-1) + index(idx-2)
}
func Fib(N int) int {

	if N == 0 || N == 1 {
		return N
	}
	if N == 2 {
		return 1
	}

	return Fib(N-1) + Fib(N-2)
}
