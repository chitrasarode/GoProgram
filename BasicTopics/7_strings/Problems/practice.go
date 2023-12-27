package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ASCII values for some characters
	fmt.Printf("ASCII value of 'A': %d\n", 'A') // Output: ASCII value of 'A': 65
	fmt.Printf("ASCII value of 'a': %d\n", 'a') // Output: ASCII value of 'a': 97
	fmt.Printf("ASCII value of '0': %d\n", '0') // Output: ASCII value of '0': 48

	// Using ASCII values to represent characters
	char := byte(65)                                                  // ASCII value for 'A'
	fmt.Printf("Character represented by ASCII value 65: %c\n", char) // Output: Character represented by ASCII value 65: A

	var runeVar rune = '世'                     // Unicode code point for the character '世'
	fmt.Printf("Rune variable: %c\n", runeVar) // Output: Rune variable: 世

	str := "Hello, 世界"
	fmt.Printf("String: %s\n", str)

	// Print each byte in the string
	for i, b := range []byte(str) {
		fmt.Printf("Byte %d: %X\n", i, b)
	}

	// Print each rune and its UTF-8 encoded bytes
	for i, r := range str {
		fmt.Printf("Rune %d: %c, Bytes: % X\n", i, r, []byte(string(r)))
	}

	// Count the number of runes in the string
	count := utf8.RuneCountInString(str)
	fmt.Printf("Number of runes: %d\n", count)

	firstRune := 'H'
	fmt.Printf("Type = %t", firstRune)

	str1 := "Hello, 世界"
	for _, char := range str1 {
		fmt.Printf("%c\t%t", char, char)

	}

}

/*
In Go, ASCII refers to the American Standard Code for Information Interchange, which is a character encoding
 standard that assigns numeric values to characters. The ASCII character set consists of 128 characters,
 including control characters(such as carriage return and line feed) and printable characters
  (such as letters, digits, and symbols).

 In Go, characters are represented using the byte data type, and the ASCII values of characters can be obtained
 using numericliterals.

Unicode defines a vast range of code points, including characters from various scripts, symbols, emojis, and
more. The Unicode standard allows for the representation of a wide variety of characters from different languages
and cultures around the world.In Unicode, a code point is a numerical value that uniquely represents a character.
Each character in the Unicode standard is assigned a unique code point, and these code points are usually
expressed in hexadecimal format.

For example, the code point for the Latin capital letter "A" is U+0041, where "U+" denotes a Unicode code point
and "0041" is the hexadecimal representation of the code point.

*/
