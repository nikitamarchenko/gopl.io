/*
Exercise 5.19: Use panic and recover to write a function that contains no 
return statement yet returns a non-zero value.
*/

package main

import "fmt"

func f() (r string) {
	
	defer func ()  {
		if err := recover(); err != nil {
			r = "val"
		}
	}()

	panic(0)

}

func main() {

	fmt.Println(f())

}