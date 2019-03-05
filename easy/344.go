package easy

func ReverseString(s []byte) {

	slen := len(s)

	for i := 0; i < slen/2; i++ {
		s[i], s[slen-1-i] = s[slen-1-i], s[i]
	}

}
