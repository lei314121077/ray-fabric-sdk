package data_transfer_chaincode

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type DataExchangeCenter struct{}

/*
@name init chaincode
 */
func (d *DataExchangeCenter) Init(stub shim.ChaincodeStubInterface) pb.Response {


	return shim.Success(nil)
}

/*
@name invoke chaincode
*/
func (d *DataExchangeCenter) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fn, args := stub.GetFunctionAndParameters()
	fmt.Sprintf("invoke 正在运行:%s " , fn)

	switch fn {
		//返回调用者信息
		case "":
			return d.addData(stub, args)
		default:
			fmt.Println("invoke 未找到函数名为: " + fn)
			return shim.Error(fmt.Sprintf("Invalid invoke function name...%s",fn ))
	}

}

/*
@name adddata
*/
func (d *DataExchangeCenter)addData(stub shim.ChaincodeStubInterface, args []string)pb.Response{

	return shim.Error("...")
}




func main() {
	if err := shim.Start(new(DataExchangeCenter)); err != nil {
		fmt.Printf("启动链码失败...！错误: %s", err)
	}
}








































