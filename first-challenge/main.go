package main

import "fmt"

func main() {
	var i = 21
	var k float64 = 123.456

	fmt.Printf("Menampilkan Nilai i : %v \n", i)
	fmt.Printf("Menampilkan Tipe Data dari Variable i : %T \n", i)
	fmt.Printf("Menampilkan Tanda : %% \n")
	fmt.Printf("Menampilkan Nilai Boolean j : %t \n", true)
	fmt.Printf("Menampilkan Nilai Boolean j : %t \n", false)
	fmt.Printf("Menampilkan Unicode Russia : %d \n", '\u024F')
	fmt.Printf("Menampilkan Nilai Base 10 : %d \n", i)
	fmt.Printf("Menampilkan Nilai Base 8 : %o \n", i)
	fmt.Printf("Menampilkan Nilai Base 16 : %x \n", 15)
	fmt.Printf("Menampilkan Nilai Base 16 : %X \n", 15)
	fmt.Printf("Menampilkan Unicode Karakter Я : %U \n", '\u024F')
	fmt.Printf("Menampilkan Float : %g \n", k)
	fmt.Printf("Menampilkan Float : %E \n", k)
}
