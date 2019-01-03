// Массив — это нумерованная последовательность элементов одного типа с фиксированной длиной

package main

import "fmt"

func main() {
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83
//короткая запись массива: x := [5]float64{ 98, 93, 77, 82, 83 }
/* Длинная запись массива:
x := [5]float64{
    98,
    93,
    77,
    82,
    83, //здесь запятая обязательна!
}
*/

    var total float64 = 0
//    for i := 0; i < 5; i++ {
//    for i := 0; i < len(x); i++ {
    for _, value := range x {
//     total += x[i]
    total += value
    }
//    fmt.Println(total / 5)
    fmt.Println(total / float64(len(x)))    // len(x) и total имеют разный тип
}
