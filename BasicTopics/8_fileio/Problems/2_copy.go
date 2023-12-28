/*
Create a program that copies the content of one file to another. Allow the user to specify the source and destination file paths.
*/
package main

import (
	"fmt"
	"io"
	"os"
)

var err error

func main() {

	var file1, file2 string

	fmt.Println("Enter Source file name")
	fmt.Scan(&file1)

	fmt.Println("Enter destination file name")
	fmt.Scan(&file2)

	sourceFile, err := os.Open(file1)
	defer sourceFile.Close()

	if err != nil {
		fmt.Println("Error in opening file : ", err)
		return
	}

	destinationFile, err := os.Create(file2)
	defer destinationFile.Close()
	if err != nil {
		fmt.Println("Error while creating file : ", err)
	}

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error while copying file : ", err)
	}
	fmt.Println("File copied successfully")
}
