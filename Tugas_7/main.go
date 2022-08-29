package main

import (
	"fmt"

	"github.com/ibrahimker/hacktiv/latihan/interface/service"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)

	res := userSvc.Register(&service.User{Nama: "Naruto"})
	fmt.Println(res)

	res = userSvc.Register(&service.User{Nama: "Sasuke"})
	fmt.Println(res)

	res = userSvc.Register(&service.User{Nama: "Joko"})
	fmt.Println(res)

	resGet := userSvc.GetUser()

	fmt.Println("-------------- HASIL GET USER -------------")

	for i, name := range resGet {
		fmt.Println(i+1, name.Nama)
	}
}
