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

	//添加交易信息
	case "CreateTransferApplication":
		return d.users.CreateTransferApplication(stub, args)

	default:
		fmt.Println("invoke 未找到函数名为: " + fn)
		return shim.Error(fmt.Sprintf("Invalid invoke function name...%s",fn ))
	}

}


