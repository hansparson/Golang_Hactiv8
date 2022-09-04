package main

import (
	"fmt"
	"sync"

	"github.com/ibrahimker/hacktiv/latihan/interface/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)

	////LIST NAMA yang Mendaftar

	arr := []string{"Naruto", "Sasuke", "Sakura", "Goku", "Doraemon", "Ultraman"}

	for _, name := range arr {
		res := userSvc.Register(&service.User{Nama: name})
		fmt.Println(res)
	}

	resGet := userSvc.GetUser()

	fmt.Println("-------------- HASIL GET USER -------------")

	// for i, name := range resGet {
	// 	fmt.Println(i+1, name.Nama)
	// }

	var wg sync.WaitGroup
	wg.Add(len(resGet))
	for _, v := range resGet {
		go cetakNama(&wg, v.Nama)
	}
	wg.Wait()

}

func cetakNama(wg *sync.WaitGroup, name string) {
	fmt.Println("Nama : ", name)
	wg.Done()
}
