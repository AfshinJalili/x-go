package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"x-go/gopl/ch2/tempconv"
)

func main() {
	from := flag.String("from", "", "the measurement unit you provide value in")
	to := flag.String("to", "", "the measurement unit you want to the value to be converted to")
	value := flag.Float64("val", math.NaN(), "value required")
	flag.Parse()

	if *from == "" || *to == "" {
		fmt.Printf("Empty from/to args\n")
		os.Exit(1)
	}

	fromVal := strings.ToLower(*from)
	toVal := strings.ToLower(*to)

	if !slices.Contains(tempconv.CommonUnits, fromVal) {
		fmt.Printf("%s is not a supported unit\n", fromVal)
		os.Exit(1)
	}
	if !slices.Contains(tempconv.CommonUnits, toVal) {
		fmt.Printf("%s is not a supported unit\n", toVal)
		os.Exit(1)
	}
	if math.IsNaN(*value) {
		fmt.Println("Please provide a value")
		os.Exit(1)
	}

	result, ok := tempconv.Convert(fromVal, toVal, *value)
	if !ok {
		fmt.Printf("cannot convert %s to %s (different unit types)\n", fromVal, toVal)
		os.Exit(1)
	}
	fmt.Printf("%g %s = %g %s\n", *value, fromVal, result, toVal)
}
