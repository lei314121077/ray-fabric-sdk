package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ray/fsdk"
	"strconv"
)


/*TODO
@name 1.新增记录接口：
@desc 我会给你一条以工单为主的信息（包括工单号，工单状态，订单号，供方id, 需方id）,你这边保存到库，返回状态给我
*/
func (o *Order) AddOrderHistoryApi(w http.ResponseWriter, r *http.Request) {
	workNo := r.FormValue("workNo")								// 工单号
	workStatus, _ := strconv.Atoi(r.FormValue("workStatus"))	// 工单状态
	orderNo := r.FormValue("orderNo")							// 订单号
	publishNo := r.FormValue("publishNo")						// 供方id
	replyNo := r.FormValue("replyNo")							// 需方id

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

	fmt.Printf("ok:%v", msg)

	js, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/*
@name 2.记录更新接口：
@desc 我会给你四个参数（工单号，用户id, 是供方还是需方，工单状态），你这边根据工单号和供方id/需方id去更新工单状态，返回状态给我
*/
func (o *Order) ModifyHistoryApi(w http.ResponseWriter, r *http.Request){

	workNo := r.FormValue("workNo")								// 工单号
	workStatus, _ := strconv.Atoi(r.FormValue("workStatus"))	// 工单状态
	orderNo := r.FormValue("orderNo")							// 订单号
	publishNo := r.FormValue("publishNo")						// 供方id
	replyNo := r.FormValue("replyNo")							// 需方id

	order := Order{
		workNo,
		workStatus,
		orderNo,
		publishNo,
		replyNo,
	}
	msg, err := o.ModifyOrderHistorySer(order, &fsdk.App)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:%v", msg)

	js, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}


/*
@name.记录查询接口：
@desc 我会给你一组工单号，用户id, 是供方还是需方这三个参数，你这边根据工单号和供方id/需方id去匹配 非这组工单号 的工单记录，并返回数据给我
*/
func (o *Order) QueryHistoryApi(w http.ResponseWriter, r *http.Request){
	workNo := r.FormValue("workNo")									// 工单号
	userId := r.FormValue("userId")									// 用户id
	isHow := r.FormValue("isHow")									// 是供方还是需方
	msg, err := o.OrderHistorySer(workNo, userId, isHow, &fsdk.App)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ok:%v", msg)

	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}





