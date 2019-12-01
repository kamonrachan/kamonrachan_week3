package main

import "fmt"

func main() {
	fmt.Print("input :")
	var name string
	var age int
	var height float32
	var weigth float32
	n, err := fmt.Scan(&name, &age, &height, &weigth)
	fmt.Println(name, age, height, weigth)
	fmt.Println(`number of argument`, n)
	fmt.Println(`error `, err)
}