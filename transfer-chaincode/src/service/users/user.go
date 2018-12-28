package users

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type User struct {}

/*
@name 将数据存储到交易账本中
*/
func (u *User)CreateTransferApplication(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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




























