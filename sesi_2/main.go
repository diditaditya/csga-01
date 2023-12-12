package main

import (
	"fmt"
	"strings"
)

func listPeople() {

	people := []map[string]any{
		{"name": "Hank", "Age": 50, "Job": "Polisi"},
		{"name": "Heisenberg", "Age": 52, "Job": "Ilmuwan"},
		{"name": "Skyler", "Age": 48, "Job": "Akuntan"},
	}

	for _, person := range people {
		fmt.Printf("Hi perkenalkan, Nama saya %v, umur saya %v, dan saya bekerja sebagai %v\n", person["name"], person["Age"], person["Job"])
	}
}

func arrangeStars(row int) {
	for i := 0; i < row; i++ {
		fmt.Println("*")
	}
}

func nestStars(row int) {
	for i := 0; i < row; i++ {
		fmt.Println(strings.Repeat("*", i+1))
	}
}

func main() {
	fmt.Println("==== No. 1 ====")
	listPeople()
	fmt.Println("==== No. 2 ====")
	arrangeStars(5)
	fmt.Println("==== No. 3 ====")
	nestStars(5)
}
