package main

import (
	"fmt"
	"gopl/ch2/cfall/meterconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No Arguments were provided.\n")
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while parsing %s to float. Err: %v", arg, err)
			os.Exit(1)
		}
		f := meterconv.Feet(t)
		m := meterconv.Meter(t)

		fmt.Printf("%f Meter(s) = %f Feet(s).\n", t, meterconv.CMtF(m))
		fmt.Printf("%f Feet(s) = %f Meter(s).\n", t,meterconv.CFtM(f))
	}
}
