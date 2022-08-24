package main

import "fmt"

func main() {

	munculin_nama := func() {
		a := "Kevin Hugo"
		b := "Kadek Bintang Anjasmara"
		c := "Guntur Satrya Saputro"
		d := "Achmad Fathoni"
		e := "Edwin Setya Noegroho"
		f := "Jaka Prima Maulana"
		g := "Stevanus Dewana"
		h := "Teguh Ridho Afdilla"
		i := "Rizki Ramadhan"
		j := "Billy Tambunan"

		list_nama := []string{a, b, c, d, e, f, g, h, i, j}
		nama := []*string{}

		for i := range list_nama {
			nama = append(nama, &list_nama[i])
		}
		fmt.Println(nama)

		for i := range list_nama {
			fmt.Println(nama[i], *nama[i])
		}
	}

	munculin_nama()
}
