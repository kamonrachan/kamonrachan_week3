package main

import(
	"fmt"
	"os"
)

func main() {
	var e int

	fmt.Scanf("%#X\n", e)
	fmt.Fprintf(os.Stdout, "%#X\n", e)
}