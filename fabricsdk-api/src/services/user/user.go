package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ray/fsdkapi"
)

//TODO
func (app *fsdkapi.Application) UserApi(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	u := Result{loginName, password}

	js, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := fsdkapi.ServiceSetup.UserDemoSer(u.name, u.passwd)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:", msg)


	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Println(loginName, password)
}

