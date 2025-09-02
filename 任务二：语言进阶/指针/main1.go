package main

import "fmt"

func setMethed(p *int) {
	*p += 10
}

func main() {
	num := 5
	fmt.Println(num)
	setMethed(&num)
	fmt.Println(num)
}
