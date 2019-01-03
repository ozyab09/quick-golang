// Напишите программу, которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz», а для кратных как трём, так и пяти — «FizzBuzz»

package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
       if i % 3 == 0 {
           if i % 5 == 0 {
            fmt.Println("FizzBuzz")
            } else {
                fmt.Println("Fizz")
            }
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
	}
}
