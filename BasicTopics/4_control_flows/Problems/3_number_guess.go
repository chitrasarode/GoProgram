/*
Guess the Number:

Generate a random number between 1 and 100. Prompt the user to guess the number. Provide feedback to the user
if their guess is too high or too low. Continue prompting until the user guesses the correct number.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("-----Guessing Number Attemps-----")
	var guessNo int
	var attempts int

	//to generate random number
	rand.Seed(time.Now().UnixNano())

	//to generate number between 1 to 100
	randomNo := rand.Intn(100) + 1

	//for loop to prompt until the user guesses the correct number.
	for {
		attempts++
		fmt.Print("Enter number to guess the random number : ")
		fmt.Scanln(&guessNo)

		if guessNo == randomNo {
			fmt.Println("Your guess is correct")
			break
		} else if guessNo < randomNo {
			fmt.Println("Guess number is lower than random number")
		} else if guessNo > randomNo {
			fmt.Println("Guess number is higher than random number")
		}
	}
	fmt.Println("your attempts : ", attempts)
	fmt.Println("Generated random number : ", randomNo)
}
