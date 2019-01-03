package common

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"service/users"
	"service/payments"
)

//
type DataExchangeCenter struct{
	user users.User
	payment payments.Payment
}

//TODO
func (d *DataExchangeCenter) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("开始实例化链码....")
	return shim.Success(nil)
}

//TODO
func (d *DataExchangeCenter) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fn, args := stub.GetFunctionAndParameters()
	fmt.Sprintf("invoke 正在运行:%s " , fn)

	switch fn {

	// 用户注册（开户）
	case "userRegister":
		return d.user.UserRegister(stub,args)
	// 用户注销
	case "userDestroy":
		return d.user.UserDestroy(stub,args)
	// 资产登记
	case "assetEnroll":
		return d.user.AssetEnroll(stub,args)
	// 资产转让
	case "assetExchange":
		return d.user.AssetExchange(stub,args)
	// 用户查询
	case "queryUser":
		return d.user.QueryUser(stub,args)
	// 资产查询
	case "queryAsset":
		return d.user.QueryAsset(stub,args)
	// 资产交易记录查询
	case "queryAssetHistory":
		return d.user.QueryAssetHistory(stub,args)
	// 转账
	case "PaymentTransfer":
		return d.payment.PaymentTransfer(stub, args)

	default:
		fmt.Println("invoke 未找到函数名为: " + fn)
		return shim.Error(fmt.Sprintf("Invalid invoke function name...%s",fn ))
	}

}


