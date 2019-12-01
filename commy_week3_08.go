package main

import "fmt"

func main() {
	fmt.Print("Input Your number:")
	var input float64
	var input2 float64
	fmt.Scanf("%f", &input)
	fmt.Scanf("%f", &input2)
	output := input / 2
	output1 := input2 + 2
	fmt.Println(output)
	fmt.Println(output1)
}
