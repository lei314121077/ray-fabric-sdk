package ordercc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"strconv"
)

//TODO
func (o *Order) AddHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("必须符合规定的参数！")
	}

	return shim.Success([]byte("保存记录操作成功！"))

}

//TODO
func (o *Order) ModifyHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("必须符合规定的参数！")
	}

	return shim.Success([]byte("更新记录操作成功！"))

}

//TODO
func (o *Order) QueryUser(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	// step 1:检查参数个数
	if len(args) != 2 {
		return shim.Error("没有足够的args!")
	}

	// step 2:验证参数正确性
	workNo := args[0]
	userID := args[1]
	if userID == ""{
		return shim.Error("无效的args!")
	}else if workNo == ""{
		return shim.Error("无效的args!")
	}
	// step 3:验证数据是否存在
	userBytes, err := stub.GetState("")
	if err != nil || len(userBytes) == 0 {
		return shim.Error("用户未找到")
	}

	return shim.Success(userBytes)
}





