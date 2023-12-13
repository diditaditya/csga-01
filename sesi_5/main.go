package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"sync"
)

const (
	RECTANGLE = "rectangle"
	CIRCLE    = "circle"
	TRIANGLE  = "triangle"
)

type Shape struct {
	ShapeType string // Circle/Rectangle
	Length    int
	Area      float32
}

func main() {
	calculateSumOfSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	input := []Shape{
		{ShapeType: RECTANGLE, Length: 5},
		{ShapeType: CIRCLE, Length: 3},
		{ShapeType: TRIANGLE, Length: 5},
		{ShapeType: RECTANGLE, Length: 15},
		{ShapeType: CIRCLE, Length: 5},
	}
	calculateArea(input)

	countWord()
}

/*
Challenge: Sum of Square Calculation

The objective of this assignment is to create a program that calculates the sum of squares of numbers concurrently using Goroutines.
● Implement a function calculateSumOfSquares(numbers []int) int that takes a slice of integers as input and calculates the sum of squares of these numbers.
● The function should divide the task of calculating the sum of squares among multiple Goroutines. Each Goroutine should handle a subset of the numbers.
● Ensure that the main application waits for all Goroutines to complete their tasks and then returns the final result.
*/
func calculateSumOfSquares(numbers []int) int {
	var wg sync.WaitGroup

	total := 0
	job := func(nums []int) {
		for _, num := range nums {
			total += num ^ 2
		}
		wg.Done()
	}

	batchSize := 3
	current := []int{}
	for i := 0; i < len(numbers); i++ {
		if len(current) < batchSize {
			current = append(current, numbers[i])
		}

		if len(current) == batchSize || i == len(numbers)-1 {
			wg.Add(1)
			go job(current)
			current = []int{}
		}
	}

	wg.Wait()
	fmt.Println("total: ", total)
	return total
}

/*
Challenge: Area Calculation

Buatlah sebuah program untuk melakukan kalkulasi dari bentuk 2-dimensi (persegi, lingkaran, dan segitiga).
Buatlah 3 buah channel masing-masing untuk persegi, lingkaran, dan segitiga.
Buatlah sebuah goroutine, untuk memasukan tiap bentuk 2-dimensi ke dalam channel nya masing-masing
Lalu ambilah value dari tiap-tiap channel dan hitunglah luas-an dari tiap bentuk 2 dimensi.
lalu print data dari tiap bentuk 2-dimensi.
Gunakan for dan select untuk mengolah data dari channel yang ada.

*/

func calculateArea(shapes []Shape) {
	recChan := make(chan int)
	triChan := make(chan int)
	cirChan := make(chan int)

	defer close(recChan)
	defer close(triChan)
	defer close(cirChan)

	go func() {
		for _, shape := range shapes {
			switch shape.ShapeType {
			case RECTANGLE:
				recChan <- shape.Length
			case CIRCLE:
				cirChan <- shape.Length
			case TRIANGLE:
				triChan <- shape.Length
			}
		}
	}()

	for i := 0; i < len(shapes); i++ {
		select {
		case recLength := <-recChan:
			area := math.Pow(float64(recLength), 2)
			fmt.Printf("Rectangle, length: %d, area: %v\n", recLength, area)
		case cirLength := <-cirChan:
			area := 3.14 * 0.25 * math.Pow(float64(cirLength), 2)
			fmt.Printf("Circle, length: %d, area: %f\n", cirLength, area)
		case triLength := <-triChan:
			area := 0.5 * float32(triLength*triLength)
			fmt.Printf("Triangle, length: %d, area: %f\n", triLength, area)
		}
	}
}

/*
Challenge

Objective: The objective of this assignment is to create a program that performs a concurrent word count on multiple text files.
Files1.txt : This is a sample text file for word count.
Files2.txt : The quick brown fox jumps over the lazy dog.
Files3.txt : Goroutines and channels are used for concurrency in Golang.

● Implement a function countWords(filename string) int that takes the name of a text file as input and returns the number of words in that file.
● Create a slice of filenames representing multiple text files to be processed.
● Use Goroutines to concurrently process each file and count the number of words.
● The main application should wait for all Goroutines to complete their tasks and then display the total word count for all files.
*/

func countWord() {
	files := []string{"Files1.txt", "Files2.txt", "Files3.txt"}

	var wg sync.WaitGroup

	count := 0
	readFile := func(fileName string) {
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		sentence := string(content)
		currentCount := len(strings.Split(sentence, " "))
		fmt.Println(sentence, currentCount)
		count += currentCount
		wg.Done()
	}

	for _, fileName := range files {
		wg.Add(1)
		go readFile(fileName)
	}

	wg.Wait()
	fmt.Println("word count: ", count)
}
