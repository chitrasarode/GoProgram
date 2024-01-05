/*
Build a program that reads a text file, replaces occurrences of a specified word with another word, and then writes the modified content

	back to the file.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.OpenFile("File3.txt", os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Error in opening file : ", err)
	}
	reader := bufio.NewReader(file)
	var modifiedData string
	for {
		data, err := reader.ReadString('\n')
		replacedData := strings.ReplaceAll(data, "rajashriChatrikar", "chitraSarode")
		modifiedData += replacedData
		if err != nil {
			break
		}
	}
	fmt.Println(modifiedData)
	fmt.Println("After modified data")

	// file, err = os.Create("File3.txt")
	// if err != nil {
	// 	fmt.Printf("error while creating file : ", err)
	// }
	// writer := bufio.NewWriter(file)
	// _, err = writer.WriteString(modifiedData)
	// if err != nil {
	// 	fmt.Printf("error while writing file : ", err)
	// } else {
	// 	fmt.Println("Data written successfully")
	// }
	// // Flush the buffer to ensure data is written to the file
	// err = writer.Flush()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file size: ", fileInfo.Size())

	_, err = file.Seek(0, 0)
	newSize := int64(90)
	err = file.Truncate(newSize)
	fmt.Println("file size: ", fileInfo.Size())

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(modifiedData)
	if err != nil {
		fmt.Println("error writing in file")
		return
	} else {
		fmt.Println("Data written successfully")
	}
	// Flush the buffer to ensure data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}

}
