/*

     *
	* *
   * * *
	* *
	 *
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Star pattern program")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println()
		}
	}

}
