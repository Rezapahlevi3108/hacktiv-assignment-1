package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var biodataList = map[int]Biodata{
	1: {"Budi", "Jakarta", "Pekerjaan Budi", "Alasan Budi"},
	2: {"Tono", "Depok", "Pekerjaan Tono", "Alasan Tono"},
	3: {"Ahmad", "Bandung", "Pekerjaan Ahmad", "Alasan Ahmad"},
	4: {"Adam", "Surabaya", "Pekerjaan Adam", "Alasan Adam"},
	5: {"Andi", "Bekasi", "Pekerjaan Andi", "Alasan Andi"},
}

func printData(absen int) {
	biodata, found := biodataList[absen]

	if !found {
		fmt.Println("Data teman dengan absen no.", absen, "tidak dapat ditemukan!")
		return
	}

	fmt.Println("Data teman dengan absen no.", absen, "adalah:")
	fmt.Println("Nama\t\t:", biodata.nama)
	fmt.Println("Alamat\t\t:", biodata.alamat)
	fmt.Println("Pekerjaan\t:", biodata.pekerjaan)
	fmt.Println("Alasan Masuk\t:", biodata.alasan)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor_absen>")
		return
	}

	absenStr := os.Args[1]
	absen, err := strconv.Atoi(absenStr)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}

	printData(absen)
}
