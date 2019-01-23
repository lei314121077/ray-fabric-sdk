package demo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ray/fsdk"
)

var (
	app = &fsdk.Application{}
)

func (d DemoController) DemoApi(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)
	res := Demo{param["name"], param["passwd"]}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg, err := d.DemoSer(res.Name, res.Passwd, app)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("ok:", msg)

	// 返回JSON格式编码
	//json.NewEncoder(w).Encode(msg)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
