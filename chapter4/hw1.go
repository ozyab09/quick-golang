//напишите программу, переводящую температуру из градусов Фаренгейта в градусы Цельсия. (C = (F - 32) * 5/9)
package main

import "fmt"

func main() {
    fmt.Println("Program of convering Fahrenheit degrees to Celsium")
    fmt.Print("Enter Farenheit degrees: ")
    var degreesFahrenheit float64
    fmt.Scanf("%f", &degreesFahrenheit) //присвоенние переменной введенного числа
    degreesCelsium := (degreesFahrenheit - 32) * 5 / 9
    fmt.Println("Celsium degrees: ", degreesCelsium)    
}

