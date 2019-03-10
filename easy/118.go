package easy

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
*/
func Generate(numRows int) [][]int {

	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0], ret[i][i] = 1, 1
		if i > 1 {
			for j := 1; j < i; j++ {
				ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}

	}

	return ret
}
