//конкатенация строк
package main

import "fmt"

func main() {
    var x string
    x = "first"
    fmt.Println(x)
	x += "second" //аналогично x = x + "second" 
	fmt.Println(x) //вывод x после конткатенации строк
	var y string
	fmt.Println(x == y) //результат сравнения
	z := "Hello World" //присвоение значения переменной при создании // тип переменной определнится автоматически
	fmt.Println(z)
}

//переменная изменит свое значение
