package main

import "fmt"

func main() {
    // fmt.Println("Hello world")
	looping_angka()
}

func looping_angka(){
	for i:=1; i<11; i++{
		if i%2 == 0 {
			fmt.Print(i, " adalah Bilangan Genap \n")
		}else{
			fmt.Print(i, " adalah Bilangan Ganjil \n")
		}
	}
}
