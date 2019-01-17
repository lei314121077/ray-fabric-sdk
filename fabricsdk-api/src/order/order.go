package order

import (
	"encoding/json"
	"fmt"
	"ray/fsdk"
	"net/http"
	"strconv"
)


/*TODO
@name 1.新增记录接口：
@desc 我会给你一条以工单为主的信息（包括工单号，工单状态，订单号，供方id, 需方id）,你这边保存到库，返回状态给我
*/

func (o *Order) AddOrderHistoryApi(w http.ResponseWriter, r *http.Request) {
	workNo := r.FormValue("workNo")
	workStatus, _ := strconv.Atoi(r.FormValue("workStatus"))	// 工单状态
	orderNo := r.FormValue("orderNo")		// 订单号
	publishNo := r.FormValue("publishNo")	// 供方id
	replyNo := r.FormValue("replyNo")		// 需方id

	order := Order{
		workNo,
		workStatus,
		orderNo,
		publishNo,
		replyNo,
	}

	msg, err := o.AddOrderHistorySer(order, &fsdk.App)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:", msg)

	js, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

