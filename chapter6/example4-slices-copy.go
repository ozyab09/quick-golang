// copy

package main

import "fmt"

func main() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}

// slice1 будет содержать [1,2,3], а slice2 — [1,2]
// Содержимое slice1 копируется в slice2, но поскольку в slice2 есть место 
// только для двух элементов, то только два первых элемента slice1 будут скопированы.