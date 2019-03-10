package easy

func ShortestToChar(S string, C byte) []int {
 
	var index,ret[]int 
	for i := range S {
		if S[i] == C {
			index = append(index, i)
		}
	}

	
	for i, j:=0,0; i < len(S); i++{
		if i <index[0]{
			ret = append(ret, index[j]-i)
		}
		
			j ++
			if j < len(index){
			}else {
				ret = append(ret, i - index[j])
			}
		}
	return ret
}