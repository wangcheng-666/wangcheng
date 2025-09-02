package main

import "fmt"

func setMethed(sum []int) {
	for index, _ := range sum {
		sum[index] *= 2
	}
}

func main() {
	fh := []int{1, 2, 3, 66, 5}
	setMethed(fh)
	fmt.Println(fh)
}
