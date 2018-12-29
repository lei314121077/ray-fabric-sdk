package users

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// 原始用户占位符
const (
	originOwner = "originOwnerPlaceholder"
)

// 用户密钥
func constructUserKey(userId string)string{
	return fmt.Sprint("user_%s",userId)
}


// 资产密钥
func constructAssetKey(assetID string)string{
	return fmt.Sprint("asset_%s",assetID)
}


// 用户注册（开户）
func (u *User) UserRegister(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	// step 1:检查参数个数
	if len(args) != 2{
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	name := args[0]
	id := args[1]
	if name == "" || id == ""{
		return shim.Error("Invalid args")
	}
	// step 3:验证数据是否存在
	if userBytes, err := stub.GetState(constructUserKey(id));err != nil || len(userBytes) != 0{
		return shim.Error("User alreay exist")
	}
	// step 4: 写入状态
	user := User{
		Name:name,
		ID:id,
		Assets:make([]string,0),
	}
	// 序列化对象
	userBytes, err := json.Marshal(user)
	if err != nil{
		return shim.Error(fmt.Sprint("marshal user error %s",err))
	}
	err = stub.PutState(constructUserKey(id), userBytes)
	if err != nil {
		return shim.Error(fmt.Sprint("put user error %s", err))
	}
	return shim.Success(nil)
}

// 用户注销
func  (u *User) UserDestroy(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 1{
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	id := args[0]
	if id == ""{
		return shim.Error("Invalid args")
	}
	// step 3:验证数据是否存在
	userBytes, err := stub.GetState(constructUserKey(id));
	if err != nil || len(userBytes) == 0{
		return shim.Error("User not found")
	}
	// step 4: 写入状态
	if err := stub.DelState(constructUserKey(id)); err != nil {
		return shim.Error(fmt.Sprintf("delete user error: %s", err))
	}
	// 删除用户名下的资产
	user := new(User)
	err = json.Unmarshal(userBytes,user)
	if err != nil{
		return shim.Error(fmt.Sprintf("unmarshal user error: %s", err))
	}
	for _,assetId := range user.Assets{
		if err := stub.DelState(constructAssetKey(assetId)); err != nil {
			return shim.Error(fmt.Sprintf("delete asset error: %s", err))
		}
	}

	return shim.Success(nil)
}

// 资产登记
func  (u *User) AssetEnroll(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 4 {
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	assetName := args[0]
	assetId := args[1]
	metadata := args[2]
	ownerId := args[3]
	if assetName == "" || assetId == "" || ownerId == ""{
		return shim.Error("Invalid args")
	}
	// step 3:验证数据是否存在
	userBytes, err := stub.GetState(constructUserKey(ownerId))
	if err != nil || len(userBytes) == 0{
		return shim.Error("User not found")
	}
	if assetBytes, err := stub.GetState(constructAssetKey(assetId)); err == nil && len(assetBytes) != 0 {
		return shim.Error("Asset already exist")
	}
	// step 4: 写入状态
	asset := &Asset{
		Name:     assetName,
		ID:       assetId,
		Metadata: metadata,
	}
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal asset error: %s", err))
	}
	if err := stub.PutState(constructAssetKey(assetId), assetBytes); err != nil {
		return shim.Error(fmt.Sprintf("save asset error: %s", err))
	}

	user := new(User)
	// 反序列化user
	if err := json.Unmarshal(userBytes, user); err != nil {
		return shim.Error(fmt.Sprintf("unmarshal user error: %s", err))
	}
	user.Assets = append(user.Assets, assetId)
	// 序列化user
	userBytes, err = json.Marshal(user)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal user error: %s", err))
	}
	if err := stub.PutState(constructUserKey(user.ID), userBytes); err != nil {
		return shim.Error(fmt.Sprintf("update user error: %s", err))
	}

	// 资产变更历史
	history := &AssetHistory{
		AssetID:        assetId,
		OriginOwnerID:  originOwner,
		CurrentOwnerID: ownerId,
	}
	historyBytes, err := json.Marshal(history)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal assert history error: %s", err))
	}

	historyKey, err := stub.CreateCompositeKey("history", []string{
		assetId,
		originOwner,
		ownerId,
	})
	if err != nil {
		return shim.Error(fmt.Sprintf("create key error: %s", err))
	}

	if err := stub.PutState(historyKey, historyBytes); err != nil {
		return shim.Error(fmt.Sprintf("save assert history error: %s", err))
	}

	return shim.Success(historyBytes)
}

// 资产转让
func  (u *User) AssetExchange(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 3 {
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	ownerID := args[0]
	assetID := args[1]
	currentOwnerID := args[2]
	if ownerID == "" || assetID == "" || currentOwnerID == ""{
		return shim.Error("Invalid args")
	}
	// step 3:验证数据是否存在
	originOwnerBytes, err := stub.GetState(constructUserKey(ownerID))
	if err != nil || len(originOwnerBytes) == 0 {
		return shim.Error("user not found")
	}

	currentOwnerBytes, err := stub.GetState(constructUserKey(currentOwnerID))
	if err != nil || len(currentOwnerBytes) == 0 {
		return shim.Error("user not found")
	}

	assetBytes, err := stub.GetState(constructAssetKey(assetID))
	if err != nil || len(assetBytes) == 0 {
		return shim.Error("asset not found")
	}

	// 校验原始拥有者确实拥有当前变更的资产
	originOwner := new(User)
	// 反序列化user
	if err := json.Unmarshal(originOwnerBytes, originOwner); err != nil {
		return shim.Error(fmt.Sprintf("unmarshal user error: %s", err))
	}
	aidexist := false
	for _, aid := range originOwner.Assets {
		if aid == assetID {
			aidexist = true
			break
		}
	}
	if !aidexist {
		return shim.Error("asset owner not match")
	}
	// step 4: 写入状态
	assetIds := make([]string, 0)
	for _, aid := range originOwner.Assets {
		if aid == assetID {
			continue
		}

		assetIds = append(assetIds, aid)
	}
	originOwner.Assets = assetIds

	originOwnerBytes, err = json.Marshal(originOwner)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal user error: %s", err))
	}
	if err := stub.PutState(constructUserKey(ownerID), originOwnerBytes); err != nil {
		return shim.Error(fmt.Sprintf("update user error: %s", err))
	}

	// 当前拥有者插入资产id
	currentOwner := new(User)
	// 反序列化user
	if err := json.Unmarshal(currentOwnerBytes, currentOwner); err != nil {
		return shim.Error(fmt.Sprintf("unmarshal user error: %s", err))
	}
	currentOwner.Assets = append(currentOwner.Assets, assetID)

	currentOwnerBytes, err = json.Marshal(currentOwner)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal user error: %s", err))
	}
	if err := stub.PutState(constructUserKey(currentOwnerID), currentOwnerBytes); err != nil {
		return shim.Error(fmt.Sprintf("update user error: %s", err))
	}

	// 插入资产变更记录
	history := &AssetHistory{
		AssetID:        assetID,
		OriginOwnerID:  ownerID,
		CurrentOwnerID: currentOwnerID,
	}
	historyBytes, err := json.Marshal(history)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal asset history error: %s", err))
	}

	historyKey, err := stub.CreateCompositeKey("history", []string{
		assetID,
		ownerID,
		currentOwnerID,
	})
	if err != nil {
		return shim.Error(fmt.Sprintf("create key error: %s", err))
	}

	if err := stub.PutState(historyKey, historyBytes); err != nil {
		return shim.Error(fmt.Sprintf("save asset history error: %s", err))
	}

	return shim.Success(nil)
}

// 用户查询
func (u *User) QueryUser(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 1 {
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	userID := args[0]
	if userID == ""{
		return shim.Error("Invalid args")
	}
	// step 3:验证数据是否存在
	userBytes, err := stub.GetState(constructUserKey(userID))
	if err != nil || len(userBytes) == 0 {
		return shim.Error("user not found")
	}

	return shim.Success(userBytes)
}

// 资产查询
func  (u *User) QueryAsset(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 1 {
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	assetID := args[0]
	if assetID == ""{
		return shim.Error("Invalid args")
	}
	// step 3:验证数据是否存在
	assetBytes, err := stub.GetState(constructAssetKey(assetID))
	if err != nil || len(assetBytes) == 0 {
		return shim.Error("asset not found")
	}

	return shim.Success(assetBytes)
}

// 资产交易记录查询
func  (u *User) QueryAssetHistory(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 2 && len(args) != 1 {
		return shim.Error("Not enough args")
	}

	// step 2:验证参数正确性
	assetID := args[0]
	if assetID == ""{
		return shim.Error("Invalid args")
	}
	queryType := "all"
	if len(args) == 2 {
		queryType = args[1]
	}

	if queryType != "all" && queryType != "enroll" && queryType != "exchange" {
		return shim.Error(fmt.Sprintf("queryType unknown %s", queryType))
	}
	// step 3:验证数据是否存在
	assetBytes, err := stub.GetState(constructAssetKey(assetID))
	if err != nil || len(assetBytes) == 0 {
		return shim.Error("asset not found")
	}

	// 查询相关数据
	keys := make([]string, 0)
	keys = append(keys, assetID)
	switch queryType {
	case "enroll":
		keys = append(keys, originOwner)
	case "exchange", "all": // 不添加任何附件key
	default:
		return shim.Error(fmt.Sprintf("unsupport queryType: %s", queryType))
	}
	result, err := stub.GetStateByPartialCompositeKey("history", keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("query history error: %s", err))
	}
	defer result.Close()

	histories := make([]*AssetHistory, 0)
	for result.HasNext() {
		historyVal, err := result.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("query error: %s", err))
		}

		history := new(AssetHistory)
		if err := json.Unmarshal(historyVal.GetValue(), history); err != nil {
			return shim.Error(fmt.Sprintf("unmarshal error: %s", err))
		}

		// 过滤掉不是资产转让的记录
		if queryType == "exchange" && history.OriginOwnerID == originOwner {
			continue
		}

		histories = append(histories, history)
	}

	historiesBytes, err := json.Marshal(histories)
	if err != nil {
		return shim.Error(fmt.Sprintf("marshal error: %s", err))
	}

	return shim.Success(historiesBytes)
}






























