package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Masukkan Kalimat : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	count := make(map[string]int)
	for _, v := range input {
		count[string(v)]++
	}

	for _, v := range input {
		fmt.Println(string(v))
	}

	fmt.Println(count)
}
