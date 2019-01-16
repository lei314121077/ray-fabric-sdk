package user

import (
	"encoding/json"
	"fmt"
	"ray/fsdk"
	"net/http"
)

var (
	app = &fsdk.Application{}
)
//TODO
func (u *User) UserApi(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	user := Result{loginName, password}

	js, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := u.UserDemoSer(user.name, user.passwd, app)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:", msg)


	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Println(loginName, password)
}

