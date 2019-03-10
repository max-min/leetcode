package easy

import (
	"fmt"
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {

	ret := make(map[string]int)
	for _, cpdomain := range cpdomains {

		strs := strings.Split(cpdomain, " ")
		if len(strs) == 2 {
			val, _ := strconv.Atoi(strs[0])

			cpd := strings.Split(strs[1], ".")
			lens := len(cpd)
			for i := 0; i < lens; i++ {
				var k string
				for j := i; j < lens; j++ {
					k += cpd[j]
					if j < lens-1 {
						k += "."
					}

				}
				fmt.Println(k)
				ret[k] = ret[k] + val
			}

		}
	}

	var strs []string
	for k, v := range ret {

		strs = append(strs, strconv.Itoa(v)+" "+k)
	}

	return strs
}
