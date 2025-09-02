package main

import "fmt"

func CountMinMap(sd []string) bool {
	fh := make(map[string]string)
	fh["{"] = "}"
	fh["["] = "]"
	fh["("] = ")"
	var sls []string
	for i := 0; i < len(sd); i++ {
		if len(sls) > 0 && fh[sls[len(sls)-1]] == sd[i] {
			sls = sls[:len(sls)-1]
			fmt.Println("匹配到的数据:", sls)
		} else {
			sls = append(sls, sd[i])
			fmt.Println("压入后的数据", sls)
		}
	}
	if len(sls) == 0 {
		return true
	}
	return false
}

func main() {
	m := "{{}()[]}"
	var str []string
	for _, s := range m {
		str = append(str, string(s))
	}
	res := CountMinMap(str)
	fmt.Println(res)
}
