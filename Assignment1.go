package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Worker struct {
	IdKaryawan   int
	NamaKaryawan string
}

type Staff struct {
	Jabatan string
	Worker
}

func InputWorker(staff *[]Staff) {

	var id int
	var nama string
	var jabatan string

	fmt.Print("Enter id: ")
	fmt.Scan(&id)

	fmt.Print("Enter name: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan Jabatan Karyawan: ")
	fmt.Scan(&jabatan)

	staffNew := Staff{}
	staffNew.IdKaryawan = id
	staffNew.NamaKaryawan = nama
	staffNew.Jabatan = jabatan

	*staff = append(*staff, staffNew)
}

func tampilStaff(staff []Staff, timeInSecs int64) {
	fmt.Println("")
	fmt.Println(strings.Repeat("_", 33))
	fmt.Println("| Id | |Nama\t |Jabatan\t|")
	fmt.Println(strings.Repeat("_", 33))

	for _, el := range staff {
		fmt.Printf("| %d  |", el.IdKaryawan)

		fmt.Printf(" |%s\t", el.NamaKaryawan)

		fmt.Printf(" |%s\t|", el.Jabatan)

		fmt.Println("")

		time.Sleep(time.Duration(timeInSecs) * 1e9)

	}
}

func main() {
	var staff []Staff

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\n")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("\tMENU")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("1. New Transaction")
		fmt.Println("2. Transaction Report")
		fmt.Println("3. List Product")
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
			InputWorker(&staff)
		case 2:
			fmt.Println("-------------------------------------------------")
			fmt.Println("------- Menampilkan Data (Durasi 3 detik) -------")
			fmt.Println("-------------------------------------------------")
			go tampilStaff(staff, 3)
		}
	}

}
