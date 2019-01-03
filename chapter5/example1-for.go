//for даёт возможность повторять список инструкций (блок) определённое количество раз

package main

import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
	}
	//второй вариант записи
	for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
