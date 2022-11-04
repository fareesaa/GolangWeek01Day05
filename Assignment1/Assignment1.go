package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"assignment1.5/model"
)

func main() {
	var staff []model.Staff

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\n")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("\tMENU")
		fmt.Println(strings.Repeat("-", 20))
		fmt.Println("1. Input Data Staff")
		fmt.Println("2. Lihat Daftar Staff")
		fmt.Println("3. EXIT")
		fmt.Printf("Masukan pilihan: ")

		scanner.Scan()
		userInput, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if userInput == 3 {
			fmt.Println("----------------------------")
			fmt.Println("------- EXIT PROGRAM -------")
			fmt.Println("----------------------------")
			break
		}
		switch userInput {
		case 1:
			var id int
			var nama string
			var jabatan string

			fmt.Print("Enter id: ")
			fmt.Scan(&id)

			fmt.Print("Enter name: ")
			fmt.Scan(&nama)

			fmt.Print("Masukkan Jabatan Karyawan: ")
			fmt.Scan(&jabatan)

			staffNew := model.Staff{Jabatan: jabatan, Worker: model.Worker{IdKaryawan: id, NamaKaryawan: nama}}
			staffNew.InputWorker(&staff)
		case 2:
			fmt.Println("-------------------------------------------------")
			fmt.Println("------- Menampilkan Data (Durasi 3 detik) -------")
			fmt.Println("-------------------------------------------------")
			go tampilStaff(staff, 3)
		}
	}
}

func tampilStaff(staff []model.Staff, timeInSecs int64) {
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
