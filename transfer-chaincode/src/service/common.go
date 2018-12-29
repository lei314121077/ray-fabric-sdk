package common

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"service/users"
)

//
type DataExchangeCenter struct{
	users users.User
}

//TODO
func (d *DataExchangeCenter) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

//TODO
func (d *DataExchangeCenter) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fn, args := stub.GetFunctionAndParameters()
	fmt.Sprintf("invoke 正在运行:%s " , fn)

	switch fn {

	// 用户注册（开户）
	case "userRegister":
		return d.users.UserRegister(stub,args)
	// 用户注销
	case "userDestroy":
		return d.users.UserDestroy(stub,args)
	// 资产登记
	case "assetEnroll":
		return d.users.AssetEnroll(stub,args)
	// 资产转让
	case "assetExchange":
		return d.users.AssetExchange(stub,args)
	// 用户查询
	case "queryUser":
		return d.users.QueryUser(stub,args)
	// 资产查询
	case "queryAsset":
		return d.users.QueryAsset(stub,args)
	// 资产交易记录查询
	case "queryAssetHistory":
		return d.users.QueryAssetHistory(stub,args)

	default:
		fmt.Println("invoke 未找到函数名为: " + fn)
		return shim.Error(fmt.Sprintf("Invalid invoke function name...%s",fn ))
	}

}


