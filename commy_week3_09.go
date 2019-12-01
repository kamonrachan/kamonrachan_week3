package main

import "fmt"

func main() {
	fmt.Printf("%t\n", 1>1)
	fmt.Printf("%t\n", 1>2)
	fmt.Printf("%t\n", 5<1)
	fmt.Printf("%t\n", 1<5)
	fmt.Printf("%t\n", 110>100)
	fmt.Printf("%t\n", 1521>=20)

}