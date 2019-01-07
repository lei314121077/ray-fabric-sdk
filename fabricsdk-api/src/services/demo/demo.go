package demo

import (
	"encoding/json"
	"net/http"
)

func (app *Application) DemoApi(w http.ResponseWriter, r *http.Request){


	res := Demo{"hello", "123456"}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
