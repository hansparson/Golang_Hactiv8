package service

type User struct {
	Nama string
}

type UserService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserIface {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {
	return u.db
}

// func register(w http.ResponseWriter, r *http.Request) {
// 	body, _ := ioutil.ReadAll(r.Body) // check for errors
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal) // check for errors
// 	name := keyVal["name"]
// 	fmt.Println(name)
// 	fmt.Fprint(w, name)
// }
