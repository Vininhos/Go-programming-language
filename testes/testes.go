package main

import "fmt"

func main() {
	var v int = 1
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}
