package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ibrahimker/hacktiv/latihan/interface/service"
)

type Employee struct {
	Name string `json:"name"`
}

var port = ":8080"
var db []*service.User
var userSvc = service.NewUserService(db)

func main() {
	http.HandleFunc("/user", getEmployrees)

	http.HandleFunc("/register", register)

	fmt.Println("Aplication is listening on port", port)
	http.ListenAndServe(port, nil)

}

func getEmployrees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resGet := userSvc.GetUser()

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(resGet)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]

	if r.Method == "POST" {
		userSvc.Register(&service.User{Nama: name})
		fmt.Fprint(w, name+" Berhasil Didaftarkan")
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}
