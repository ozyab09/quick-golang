//Чему равна длина среза, созданного таким способом: make([]int, 3, 9)

package main

import "fmt"

func main() {
    slice := make([]int, 3, 9)
    fmt.Println(len(slice)) 
}