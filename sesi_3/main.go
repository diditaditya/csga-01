package main

import (
	"fmt"
	"strings"
)

func main() {
	alayed := alayGen("hello", "dunia", "apa", "kabarnya", "xoxo")
	fmt.Println(alayed)

	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
	fibbed := fibonacci(5)
	fmt.Println(fibbed)
}

/*
Buatlah sebuah variadic function yang dapat menerima argumen string yang tak terbatas jumlahnya.
Function tersebut akan me-return sebuah kalimat string, yang merupakan gabungan dari tiap kata yang ada di argumen.
Contoh:```AlayGen("hello", "world", "zzz") // akan menghasilkan "hello world zzz"```
selain mengabungkan katanya, kita perlu membuat tiap kata menjadi alay, dengan mengubah
```
a -> 4
e -> 3
i -> !
l -> 1
n -> N
s -> 5
x -> *
```
untuk mempermudah pengerjaan gunakan konsep modular function
*/

func alayGen(words ...string) string {
	alay := []string{}

	for _, word := range words {
		converted := ""
		for _, char := range word {
			stringified := string(char)
			switch stringified {
			case "a":
				converted += "4"
			case "e":
				converted += "3"
			case "i":
				converted += "!"
			case "l":
				converted += "1"
			case "n":
				converted += "N"
			case "s":
				converted += "5"
			case "x":
				converted += "*"
			default:
				converted += stringified
			}
		}
		alay = append(alay, converted)
	}

	return strings.Join(alay, " ")
}

/*
Buatlah sebuah function yang menerima argument sebuah `n` int, dan akan me-return bilangan Fibonacci ke-`n`
Untuk mempermudah berikut pengertian dan contoh Fibonacci
Source : https://en.wikipedia.org/wiki/Fibonacci_sequence

0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
*/

func fibonacci(n int) int {
	if n < 1 {
		fmt.Println("Cannot be less than 1")
		return 0
	}

	values := []int{0, 1}

	if n-1 < len(values) {
		return values[n-1]
	}

	currentIndex := 1
	for currentIndex < n {
		currentVal := values[currentIndex] + values[currentIndex-1]
		values = append(values, currentVal)
		currentIndex++
	}

	return values[len(values)-1]
}
