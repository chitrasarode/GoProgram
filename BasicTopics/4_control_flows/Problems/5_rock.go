/*
Rock, Paper, Scissors Game:

Implement a simple text-based Rock, Paper, Scissors game. Prompt the user to enter their choice
(Rock, Paper, or Scissors) and generate a random choice for the computer. Determine the winner based on the
rules: Rock crushes Scissors, Scissors cuts Paper, Paper covers Rock. Both having same choice is a tie.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userchoice string
	fmt.Println("\n-----Rock,Paper,or Scissors Game-----")
	rand.Seed(time.Now().UnixNano())
	choices := []string{"rock", "paper", "scissor"}
	randomIndex := rand.Intn(len(choices))
	randomChoice := choices[randomIndex]

	fmt.Print("Enter your choice : {rock,paper or scissor} : ")
	fmt.Scanln(&userchoice)
	if userchoice != "rock" && userchoice != "paper" && userchoice != "scissor" {
		fmt.Println("Please enter proper choice")
		return
	}
	if userchoice == randomChoice {
		fmt.Println("Both have same choice so its tie")
	} else if userchoice == "rock" && randomChoice == "scissor" {
		fmt.Println("Congratulation!...You are winner...Rock crushes scissors")
	} else if userchoice == "scissor" && randomChoice == "paper" {
		fmt.Println("Congratulation!...You are winner...Scissors cuts Paper")
	} else if userchoice == "paper" && randomChoice == "rock" {
		fmt.Println("Congratulation!...You are winner...Paper covers Rock")
	} else {
		fmt.Println("Oops!...You lost the game")
	}
	fmt.Println("Random choice is : ", randomChoice)

}
