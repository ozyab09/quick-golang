/*
Напишите программу, которая находит самый наименьший элемент в этом списке:
x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
*/

package main

import "fmt"

func main() {
    
    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    
    minElement := 0
    for _, value := range x {
        if minElement > value {
            minElement = value
        } else if minElement == 0 {
            minElement = value
        }
    }

    fmt.Println(minElement)

}