package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
}

/* модификация hello.go из chapter2
go run hello.go
результат выполнения:
11
101
Hello World
Разбор:
11 - Пробел тоже считается символом, поэтому длина строки 11 символов, а не 10
101 - Строки “индексируются” начиная с 0, а не с 1.
[1] даст второй элемент, а не первый. Мы видим 101 вместо e - это происходит из-за того, что символ представляется байтом (байт — это целое число)
Hello World - конкатенация строк, полагаясь на типы аргументов. Если по обе стороны от + находятся строки, компилятор предположит, что вы имели в виду конкатенацию, а не сложение.
*/
