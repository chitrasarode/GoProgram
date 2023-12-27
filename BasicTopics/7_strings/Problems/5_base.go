/*
Implement a simple Base64 encoder and decoder. The encoder should take a string and encode it into Base64,
 while the decoder should reverse the process. Leverage the base64 package for encoding and decoding.
*/

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("-----Base 64 encoding-----")
	str := "hello,world!"
	strByte := []byte(str) //convert string to byte slice
	fmt.Println("Original String is ", str)

	//EncodeToString function takes byte slice
	encodedString := base64.StdEncoding.EncodeToString(strByte) //encode slice byte in base 64
	fmt.Println("Base 64 Encoded string : ", encodedString)

	//DecodeString returns decoded string and error
	decodedString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println("Error in dcoding string", err)
	}
	fmt.Println("Base 64 Decoded string in byte slice : ", decodedString)

	//convert this byte slice to string
	originalStr := string(strByte)
	fmt.Println("String : ", originalStr)

}
