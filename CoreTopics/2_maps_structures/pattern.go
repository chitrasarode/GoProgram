/*
*
**
***
****
*****
****
***
**
*

 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Star pattern program")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 4; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
