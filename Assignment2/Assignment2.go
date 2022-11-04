package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Mahasiswa struct {
	IdMhs int
	Nama  string
	Nilai []int
}

func main() {
	var mahasiswa []Mahasiswa
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\n")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("\tMENU")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("1. Input Data Mahasiswa")
		fmt.Println("2. List Data Mahasiswa")
		fmt.Println("3. Convert Data Mahasiswa(.csv)")
		fmt.Println("4. Exit")
		fmt.Printf("Masukan pilihan: ")

		scanner.Scan()
		userInput, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if userInput == 4 {
			fmt.Println("----------------------------")
			fmt.Println("------- EXIT PROGRAM -------")
			fmt.Println("----------------------------")
			break
		}
		switch userInput {
		case 1:
			InputMhs(&mahasiswa)
		case 2:
			// fmt.Println("-------------------------------------------------")
			// fmt.Println("------- Menampilkan Data (Durasi 3 detik) -------")
			// fmt.Println("-------------------------------------------------")
			// go tampilStaff(staff, 3)
			sort.Slice(mahasiswa, func(i, j int) bool {
				return mahasiswa[i].IdMhs < mahasiswa[j].IdMhs
			})
			tampilMhs(mahasiswa)
		case 3:
			var namaFile string
			fmt.Print("Masukkan Nama File: ")

			fmt.Scan(&namaFile)
			go tampilMhs(mahasiswa)
			go fileWriter(mahasiswa, namaFile)
		}

	}
}

func InputMhs(mahasiswa *[]Mahasiswa) {
	var id int
	var nama string
	var ing int
	var fis int
	var alg int

	fmt.Print("Enter id: ")
	fmt.Scan(&id)

	fmt.Print("Enter name: ")
	fmt.Scan(&nama)
	fmt.Print("Enter Poin of English: ")
	fmt.Scan(&ing)
	fmt.Print("Enter Poin of Physics: ")
	fmt.Scan(&fis)
	fmt.Print("Enter Poin of Algorithm: ")
	fmt.Scan(&alg)

	nilai := []int{ing, fis, alg}

	mhsNew := Mahasiswa{IdMhs: id, Nama: nama, Nilai: nilai}

	*mahasiswa = append(*mahasiswa, mhsNew)
}

func fileWriter(mahasiswa []Mahasiswa, namaFile string) {
	file, err := os.Create(namaFile + ".csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	var data [][]string
	data = append(data, []string{"id", "nama", "b.inggris", "fisika", "algoritma"})
	for _, record := range mahasiswa {
		row := []string{strconv.Itoa(int(record.IdMhs)), record.Nama, strconv.Itoa(int(record.Nilai[0])), strconv.Itoa(int(record.Nilai[1])), strconv.Itoa(int(record.Nilai[2]))}
		data = append(data, row)
	}
	w.WriteAll(data)
}

func tampilMhs(mahasiswa []Mahasiswa) {
	fmt.Println("")
	fmt.Println("| Id | |Nama\t |English\t |Physics\t |Algorithm\t |")
	for _, el := range mahasiswa {
		fmt.Printf("| %d  | |%s\t |%d\t\t |%d\t\t |%d\t\t |", el.IdMhs, el.Nama, el.Nilai[0], el.Nilai[1], el.Nilai[2])

		fmt.Println("")

		time.Sleep(time.Duration(1) * 1e9)
	}
}
