//Область видимости переменной

package main

import "fmt"

//переменная вынесена из функции main
var x string = "Hello World"

func main() {
    fmt.Println(x)
}
//переменная видна функции f
func f() {
    fmt.Println(x)
}

//переменная видна внутри блока {}
