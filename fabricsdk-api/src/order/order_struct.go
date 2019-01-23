package order


type Order struct {
	WorkNo 		string	`json:"workNo,omitempty"`		// 包括工单号
	WorkStatus 	int		`json:"workStatus,omitempty"`	// 工单状态
	OrderNo 	string 	`json:"orderNo,omitempty"`		// 订单号
	PublishNo 	string 	`json:"publishNo,omitempty"`	// 供方id
	ReplyNo 	string 	`json:"replyNo,omitempty"`		// 需方id
}

