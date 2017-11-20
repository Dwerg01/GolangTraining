/*package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, " -- BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}*/

package main

import "fmt"

func main() {

	var i int
	for i = 0; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 && i != 0 {
			fmt.Println(i, "-- FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "-- Fizz")
		} else if i%3 == 0 {
			fmt.Println(i, "-- Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
