package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (boilingF - 32) * 5 / 9
	fmt.Printf("Boiling Point: %g^F or %g^C\n", f, c)
}
