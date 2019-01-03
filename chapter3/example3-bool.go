package main

import "fmt"

func main() {
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}

/* go run bool.go
результат согласно таблице истинности:
true
false
true
true
false
*/

