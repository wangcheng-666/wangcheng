package main

import (
	"fmt"
)

func Getstr(str []string) string {
	var s string
	for i := 0; i < len(str[0]); i++ {
		for j := 1; j < len(str); j++ {
			if str[0][i] == str[j][i] {
				if len(str)-1 == j {
					s += string(str[0][i])
				}
			} else {
				return s
			}
		}

	}
	return s
}

func main() {
	str := make([]string, 3)
	str[0] = "silpe"
	str[1] = "silpd"
	str[2] = "silas"
	fmt.Println(Getstr(str))
}
