/*
Exercise 1.3: Experiment tom easure the difference in running time bet ween our  
potentially ineffiient versions and the one that uses strings.Join. 
(Section 1.6 illustrates part of the time pakage, and Setion 11.4 shows how 
to write benchmark tests for systematic performance evaluation.)
*/

package main

import (
	"fmt"
	"os"
)


func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, " ", arg)
	}
}