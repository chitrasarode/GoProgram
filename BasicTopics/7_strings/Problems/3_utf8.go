/* Implement a function that counts the number of characters (runes) in a UTF-8 encoded string. Use the utf8
package to handle Unicode characters correctly.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var count int
	utf8String := "こんにちは，世界！"

	fmt.Println("UTF8 string is : ", utf8String)
	fmt.Println("length of utf8 string is : ", len(utf8String))

	count = utf8.RuneCountInString(utf8String) //used to count number of characters
	fmt.Println("Number of character in utf8 string is : ", count)

	var count1 int
	for len(utf8String) > 0 {
		_, size := utf8.DecodeRuneInString(utf8String) //returns decoded firstrune and size
		utf8String = utf8String[size:]                 //essentially it will remove first chracter
		count1++
	}
	fmt.Println("Number of character in utf8 string is : ", count1)

	utf8String1 := "こんにちは，世界！"
	cnt := 0
	for _, char := range utf8String1 {
		if char > 127 {
			cnt++
		}
	}
	fmt.Println("Number of character in utf8 string is : ", cnt)

}
