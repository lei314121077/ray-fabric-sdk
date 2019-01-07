package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func (app *Application) UserApi(w http.ResponseWriter, r *http.Request) {
	loginName := r.FormValue("loginName")
	password := r.FormValue("password")

	res := Result{loginName, password}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Println(loginName, password)
}

