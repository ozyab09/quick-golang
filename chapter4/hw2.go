//Напишите программу для перевода футов в метры (1 фут = 0.3048 метр).

package main

import "fmt"

func main() {
    fmt.Println("Program of convering feet degrees to meters")
    fmt.Print("Enter value in feets: ")
    var valueFeets float64
    fmt.Scanf("%f", &valueFeets) //присвоенние переменной введенного числа
    valueMeters := valueFeets  * 0.3048
    fmt.Println("Value in meters: ", valueMeters)    
}

