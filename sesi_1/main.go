package main

import "fmt"

/*
Challenge Basic GO

1. Buatlah sebuah variabel dengan nama `myNum` dengan tipe data int32 dan assign nilai 50 ke dalam variabel tersebut
   Lalu, print variable tersebut ke dalam terminal

2. Buatlah sebuah variabel dengan nama `myNum2` dengan tipe data float32 dan assignnilai 51 ke dalam variabel tersebut
   Lalu, print variable tersebut ke dalam terminal

3. Buatlah sebuah variabel dengan nama `myNumStr` dengan tipe data string dan assignkarakter "50" ke dalam variabel tersebut
   Lalu, print variable tersebut ke dalam terminal

4. Buatlah variabel dengan dengan ketentuan berikut
  - x dan assign nilai int32 5
	- y dan assign nilai int32 105.

	Buatlah variabel baru dengan nama z, yang merupakan hasil penjumlahan dari x dan y. Lalu print ke dalam terminal

*/

func main() {
	var myNum int32 = 50
	fmt.Println("myNum: ", myNum)

	var myNum2 float32 = 51
	fmt.Println("myNum2: ", myNum2)

	var MyNumStr string = "50"
	fmt.Println("myNumStr: ", MyNumStr)

	x := 5
	y := 105
	z := x + y
	fmt.Printf("z = %v + %v = %v \n", x, y, z)
}
