package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	for i, v := range os.Args[1:] {
		fmt.Println(i, v)
	}

	fmt.Printf("For range: %.2f\n", time.Since(start).Seconds())

	start = time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Printf("strings.Join: %.2f\n", time.Since(start).Seconds())
}
