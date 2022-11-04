package model

import (
	"fmt"
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

// func (staff Staff) InputWorker(staff1 *[]Staff) Staff {
// 	var staffPush Staff
// 	staffPush.IdKaryawan = staff.IdKaryawan
// 	staffPush.NamaKaryawan = staff.NamaKaryawan
// 	staffPush.Jabatan = staff.Jabatan

// 	return staffPush
// }

func (staff Staff) InputWorker(staff1 *[]Staff) {
	var staffPush Staff
	staffPush.IdKaryawan = staff.IdKaryawan
	staffPush.NamaKaryawan = staff.NamaKaryawan
	staffPush.Jabatan = staff.Jabatan
	*staff1 = append(*staff1, staffPush)
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
