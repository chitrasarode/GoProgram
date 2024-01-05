// Online Quiz System:

// Design a program to manage an online quiz system. Create a struct to represent a quiz with fields like quiz ID,
// questions, and correct answers. Use a map to store quizzes, and implement functions to add a new quiz, allow
// users to attempt a quiz, and calculate their scores.

package main

import (
	"fmt"
)

type quiz struct {
	quizId    int
	questions string
	answers   string
	quiz1     map[int]quiz
}

func main() {
	fmt.Println("*****Online Quiz System*****")
	var q = new(quiz)
	q.quiz1 = make(map[int]quiz)
	q.addQuiz(1, "What is 10 + 10", "20")
	q.addQuiz(2, "What is 10 * 10", "100")
	q.viewQuiz()
}

func (q *quiz) addQuiz(id int, question, answer string) {
	newQuiz := quiz{
		quizId:    id,
		questions: question,
		answers:   answer,
	}
	q.quiz1[id] = newQuiz
}

func (q *quiz) viewQuiz() {
	for _, val := range q.quiz1 {
		fmt.Println("Quiz Id =", val.quizId, " Question :", val.questions, " Answer :", val.answers)
	}

}
