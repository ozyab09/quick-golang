//Как назвать переменную
package main

import "fmt"

func main() {
	x := "Max"
	fmt.Println("My dog's name is", x)
	//но лучше давать осмысленные наименования переменным
	dogsName := "Max" 	//lower CamelCase (или camelBack)
	fmt.Println("My dog's name is", dogsName)
}
