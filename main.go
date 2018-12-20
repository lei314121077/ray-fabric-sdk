package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DataExchangeCenter struct{}

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
		return d.CreateTransferApplication(stub, args)

	default:
			fmt.Println("invoke 未找到函数名为: " + fn)
			return shim.Error(fmt.Sprintf("Invalid invoke function name...%s",fn ))
	}

}


/*
@name 将数据存储到交易账本中
*/
func (d *DataExchangeCenter)CreateTransferApplication(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Entering CreateTransferApplication")

	if len(args) < 2 {
		fmt.Println("args参数错误！< 2")
		return shim.Error("交易数据参数错误！")
	}

	var transferApplicationId, transferApplicationInput = args[0],args[1]
	err := stub.PutState(transferApplicationId, []byte(transferApplicationInput))
	if err != nil {
		fmt.Println("无法保存交易数据：", err)
		return shim.Error(fmt.Sprintf("保存交易数据失败！", err))
	}

	fmt.Println("保存交易数据成功")
	//TODO 未确定返回的格式
	return shim.Success(nil)
}


func main() {
	if err := shim.Start(new(DataExchangeCenter)); err != nil {
		fmt.Printf("启动链码失败...！错误: %s", err)
	}
	fmt.Println("Chaincode成功启动!")
}








































