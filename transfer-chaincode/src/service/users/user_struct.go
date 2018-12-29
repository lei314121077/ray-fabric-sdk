package users

/*
TODO
需求发布者
*/
type PublishData struct{
	Owner      string `json:"owner"`				   //所有者
}


/*
TODO
需求响应者
*/
type ReplyData struct {
	Owner      string `json:"owner"`					//所有者
}

// 用户
type User struct {
	Name string `json:"name"`
	ID string `json:"id"`
	Assets []string `json:"assets"`
}
// 资产
type Asset struct {
	Name string `json:"name"`
	ID string `json:"id"`
	Metadata string `json:"metadata"`
}
// 资产变更记录
type AssetHistory struct {
	AssetID string `json:"asset_id"`
	OriginOwnerID string `json:"origin_owner_id"`
	CurrentOwnerID string `json:"current_owner_id"`
}


