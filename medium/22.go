package medium

func GenerateParenthesis(n int) []string {

	if n == 0 {
		return []string{""}
	}
	var ret []string
	for n > 0 {
		var s string
		for i := 0; i < n; i++ {
			index := n
			for i := 0; i < index; i++ {
				s += "("
			}
			for i := 0; i < index; i++ {
				s += ")"
			}
		}
		ret = append(ret, s)
		n--
	}

	return ret
}
