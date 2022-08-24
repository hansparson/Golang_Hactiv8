package main

import "fmt"

func main() {
	salam()

}

func salam() {
	name := []string{
		"Kevin Hugo",
		"Kadek Bintang Anjasmara ",
		"Guntur Satrya Saputro",
		"Achmad Fathoni",
		"Edwin Setya Noegroho",
		"Jaka Prima Maulana",
		"Stevanus Dewana",
		"Teguh Ridho Afdilla",
		"Rizki Ramadhan",
		"Billy Tambunan"}

	// fmt.Println("My Slice 1:", name)

	for i, s := range name {
		fmt.Println(i+1, s)
	}
}
