package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1 + 1) // используя целочисленые значения
	fmt.Println("1 + 1 =", 1.0 + 1.0) // используя числа с плавающей точкой

}

//go run int.go

//мы используем .0, чтобы сказать Go, что это число с плавающей точкой, а не целое
