package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score int
}

func (s *Student) AverageScore(students []Student) float64 {
	totalScore := 0
	for _, student := range students {
		totalScore += student.Score
	}
	return float64(totalScore) / float64(len(students))
}

func (s *Student) MinMaxScore(students []Student) (int, int) {
	minScore := students[0].Score
	maxScore := students[0].Score
	for _, student := range students {
		if student.Score < minScore {
			minScore = student.Score
		}
		if student.Score > maxScore {
			maxScore = student.Score
		}
	}
	return minScore, maxScore
}

func main() {
	var students []Student
	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("input %d Student's Name: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("input %d Students Score: ", i+1)
		fmt.Scan(&score)
		students = append(students, Student{Name: name, Score: score})
	}

	var student Student
	averageScore := student.AverageScore(students)
	minScore, maxScore := student.MinMaxScore(students)

	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Min Score of students: %d\n", minScore)
	fmt.Printf("Max score of students: %d\n", maxScore)
}
