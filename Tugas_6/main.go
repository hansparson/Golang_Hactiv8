package main

import (
	"fmt"
	"os"
	"strconv"
)

type company struct {
	penjelasan string
	employees  []DataPerson
}

type DataPerson struct {
	name    string
	address string
	job     string
	Reason  string
}

func main() {

	arg := os.Args[1]
	fmt.Println(arg)
	i, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
	data_orang(i)

}

func data_orang(peserta int) {
	data_1 := DataPerson{"Kevin Hugo", "Meruya", "Guru", "Senang"}
	data_2 := DataPerson{"Kadek Bintang Anjasmara", "Kebon Jeruk", "Arsitek", "Penasaran"}
	data_3 := DataPerson{"Guntur Satrya Saputro", "Tanjung Duren", "Insinyur Mesin", "Ingin Tau"}
	data_4 := DataPerson{"Achmad Fathoni", "Kebon Sirih", "Perawat", "Menambah Pengalaman"}
	data_5 := DataPerson{"Edwin Setya Noegroho", "Pasar Senen", "Dokter", "Menambah Pengetahuan"}
	data_6 := DataPerson{"Jaka Prima Maulana", "Senayan", "Pemadam Kebakaran", "Menambah Ilmu"}
	data_7 := DataPerson{"Stevanus Dewana", "Grand Indonesia", "Kondektur", "Jam Kosong"}
	data_8 := DataPerson{"Teguh Ridho Afdilla", "Monumen Nasional", "Pilot", "Belajar Hal Baru"}
	data_9 := DataPerson{"Rizki Ramadhan", "Pluit", "Masinis", "Cari Teman Baru"}
	data_10 := DataPerson{"Billy Tambunan", "Condet", "Penulis", "Tidak ada Alasan"}

	employees := []DataPerson{data_1, data_2, data_3, data_4, data_5, data_6, data_7, data_8, data_9, data_10}

	fmt.Printf("%s Bertempat Tinggal Di %s dan bekerja sebagai %s, pendapat ikut Kelas Golang : %s", employees[peserta].name, employees[peserta].address, employees[peserta].job, employees[peserta].Reason)
}
