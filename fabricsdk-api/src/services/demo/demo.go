package demo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ray/fsdkapi"
)

func (app *fsdkapi.Application) DemoApi(w http.ResponseWriter, r *http.Request){


	res := Demo{"hello", "123456"}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := fsdkapi.ServiceSetup.DemoSer(u.name, u.passwd)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:", msg)


	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
