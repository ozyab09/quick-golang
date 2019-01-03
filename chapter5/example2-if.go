//Оператор if аналогичен оператору for в том, что он выполняет блок в зависимости от условия. 

package main

import "fmt"
func main() {
    for i := 1; i <= 10; i++ {
       if i % 2 == 0 {
            fmt.Println(i, "even")
        // } else if i % 3 == 0 {
        //  divisible by 3   
        } else {
            fmt.Println(i, "odd")
        }
    }
}