package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func main() {
	samples := []string{
		"Example Word", "As Soon As Possible", "Liquid-crystal display", "Thank George It's Friday!", "Portable Network Graphics",
	}
	for _, sample := range samples {
		aggro(sample)
	}

	fmt.Printf("\n===============================================\n\n")

	scrabbler("exampleword")
}

/*
Buat sebuah program yang mengubah istilah panjang seperti Portable Network Graphics menjadi akronimnya yaitu PNG
Dengan catatan tanda hubung adalah pemisah kata (seperti spasi); semua tanda baca lainnya dapat dihapus dari input.

Contoh:
As Soon As Possible- ASAP
Liquid-crystal display- LCD
Thank George It's Friday!- TGIF

Buatlah program yang dapat menerima input slice of string,
dan akan menampilkan hasil akronim nya dengan format`Example Word - EW`
Buatlah program agar berjalan secara asynchronous untuk setiap elemen input string nya.
*/
func aggro(words string) {
	if len(words) == 0 {
		return
	}

	pattern := "[\\s-]"
	regex := regexp.MustCompile(pattern)
	splitResult := regex.Split(words, -1)

	type Result struct {
		idx  int
		char string
	}

	job := func(idx int, word string, ch chan Result) {
		char := string(word[0])
		result := Result{
			idx:  idx,
			char: strings.ToUpper(char),
		}
		ch <- result
	}

	ch := make(chan Result)
	defer close(ch)
	for i, word := range splitResult {
		go job(i, word, ch)
	}

	letters := make([]string, len(splitResult))
	for i := 0; i < len(splitResult); i++ {
		res := <-ch
		letters[res.idx] = res.char
	}

	fmt.Println(words, " - ", strings.Join(letters, ""))
}

/*
Buatlah sebuah program yang dapat menerima input berupa slice of string.
Program tersebut akan menghitung score scrabble dari tiap string nya,lalu program tersebut akan menampilkan hasilnya dengan format
`exampleword | scrabble score 30`
berikut adalah point untuk tiap huruf

Letter                                                       Value
A, E, I, O, U, L, N, R, S, T                          				1
D, G                                                          2
B, C, M, P                                                    3
F, H, V, W, Y                                                 4
K                                                             5
J, X                                                          8
Q, Z                                                          10

Kerjakan program dengan asynchronous untuk setiap elemen string nya sehingga program akan berjalan denganlebihcepat
*/
var mapper = map[string]int{
	"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
	"D": 2, "G": 2,
	"B": 3, "C": 3, "M": 3, "P": 3,
	"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
	"K": 5,
	"J": 8, "X": 8,
	"Q": 10, "Z": 10,
}

func scrabbler(word string) {
	var wg sync.WaitGroup

	count := 0

	job := func(r rune) {
		char := strings.ToUpper(string(r))
		if val, ok := mapper[char]; ok {
			count += val
		}
		wg.Done()
	}

	for _, r := range word {
		wg.Add(1)
		go job(r)
	}

	wg.Wait()

	fmt.Println(word, " | scrabble score ", count)
}
