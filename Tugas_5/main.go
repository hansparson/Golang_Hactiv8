package main

import "fmt"

type Employee struct {
	name string
}

func main() {

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
	employees := []*Employee{}

	for i, name := range list_nama {
		fmt.Printf("loop ke %d, nama: %s\n", i, name)
		employees = append(employees, &Employee{name: name})
	}
	fmt.Printf("Isi array employees: %+v\n", employees)
	for _, e := range employees {
		fmt.Println(e.name)
	}
}
