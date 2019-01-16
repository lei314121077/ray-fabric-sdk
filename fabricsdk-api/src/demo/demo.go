package demo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ray/fsdk"
)

var (
	app = &fsdk.Application{}
)

func (d DemoController) DemoApi(w http.ResponseWriter, r *http.Request){

	res := Demo{"hello", "123456"}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := d.DemoSer(res.name, res.passwd, app)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:", msg)


	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
