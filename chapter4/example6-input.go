//пример программы

package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input) //присвоенние переменной введенного числа
    output := input * 2
    fmt.Println(output)    
}
