package main

import (
	"fmt"
	"os"
)

func main() {

	var age1 int
    var age2 int
	var age3 int

	fmt.Scan(&age1, &age2, &age3)
	fmt.Fprintf(os.Stdout, "%d - %d - %d", age1, age2, age3)
}