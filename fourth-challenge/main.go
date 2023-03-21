package main

import (
	"fmt"
	"os"
)

type students struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

const job = "IT Dev"

func biodata(listNama []string) []students {
	address := []string{"Jakarta", "Bandung", "Karawang", "Purwakarta", "Yogyakarta"}
	reason := []string{"Milik Beni", "Milik Budi", "Milik Caca", "Milik Dede", "Milik Danang"}
	generate := make([]students, 0)
	var s students
	for key, value := range listNama {
		s.nama = value
		s.alamat = address[key]
		s.pekerjaan = job
		s.alasan = reason[key]
		generate = append(generate, s)
	}
	return generate
}

func main() {
	nama := []string{"Beni", "Budi", "Caca", "Dede", "Danang"}
	output := biodata(nama)
	var arg string = os.Args[1]
	var inama int

	for i, v := range nama {
		if arg == v {
			inama = i
		}
	}

	for i, v := range output {
		if inama == i {
			fmt.Println("ID : ", i)
			fmt.Println("Nama : ", v.nama)
			fmt.Println("Alamat : ", v.alamat)
			fmt.Println("Pekerjaan : ", v.pekerjaan)
			fmt.Println("Alasan : ", v.alasan)
		}
	}
	//fmt.Println(output)
}
