package main

import (
	"fmt"
	"time"
)

func main() {
	inputTime := "2022-01-01 12:34:56"
	fmt.Printf("inputTime type : %T \n", inputTime)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", inputTime)
	fmt.Printf("ParsedTime Type : %T \n", parsedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed Time: ", parsedTime)

	currentTime := time.Now()
	formattedTime := currentTime.Weekday().String()[:3]
	fmt.Printf("Formatted Time: %s\n current Time :%s", formattedTime, currentTime)

	zone, offset := currentTime.Zone()
	fmt.Println("Current Zone : ", zone, "\nOffset : ", offset)

	location, err := time.LoadLocation("America/New_York")
	currentTime1 := time.Now().In(location)
	fmt.Println("current time in New York : ", currentTime1)
	fmt.Println(currentTime1.Zone())

	location1, err := time.LoadLocation("Europe/London")
	currentTime2 := time.Now().In(location1)
	fmt.Println("current time in London : ", currentTime2)
	fmt.Println(currentTime2.Zone())

	location2, err := time.LoadLocation("Asia/Tehran")
	currentTime3 := time.Now().In(location2)
	fmt.Println("current time in Tokyo : ", currentTime3)
	fmt.Println(currentTime3.Zone())

}
