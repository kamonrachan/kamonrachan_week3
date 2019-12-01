package main

import "fmt"

func main() {
	var e1, e2 float32
	fmt.Scan(&e1, &e2)
	e3 := e1 + e2
	fmt.Printf("%.30e", e3)
}

