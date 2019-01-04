package payments

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)


// 向指定的账户存入对应的金额
// -c '{"Args":["set", "账户名称", "要存入的金额"]}'
func (p *Payment) set(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("必须且只能指定账户名称及要存入的金额")
	}

	result, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据指定的账户查询信息失败")
	}

	if result == nil {
		return shim.Error("指定的账户不存在")
	}

	// 存入账户
	val, err := strconv.Atoi(string(result))
	if err != nil {
		return shim.Error("处理指定的账户金额时发生错误")
	}

	x, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("指定要存入的金额错误")
	}

	val = val + x

	// 保存信息
	err = stub.PutState(args[0], []byte(strconv.Itoa(val)))
	if err != nil {
		return shim.Error("存入账户金额时发生错误")
	}

	return shim.Success([]byte("存入操作成功"))

}



// 转账
// -c '{"Args":["Payment", "源账户名称", "目标账户名称", "转账金额"]}'
func (p *Payment)PaymentTransfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("必须且只能指定源账户及目标账户名称与对应的转账金额")
	}

	var source, target string
	var x string

	source = args[0]
	target = args[1]
	x = args[2]

	// 源账户扣除对应的转账金额
	// 目标账户加上对应的转账金额

	// 查询源账户及目标账户的余额
	sval, err := stub.GetState(source)
	if err != nil {
		return shim.Error("查询源账户信息失败")
	}
	// 如果源账户或目标账户不存在的情况下
	// 不存在的情况下直接return

	tval, err := stub.GetState(target)
	if err != nil {
		return shim.Error("查询目标账户信息失败")
	}

	// 实现转账
	s, err := strconv.Atoi(x)
	if err != nil {
		return shim.Error("指定的转账金额错误")
	}

	svi, err := strconv.Atoi(string(sval))
	if err != nil {
		return shim.Error("处理源账户余额时发生错误")
	}

	tvi, err := strconv.Atoi(string(tval))
	if err != nil {
		return shim.Error("处理目标账户余额时发生错误")
	}

	if svi < s {
		return shim.Error("指定的源账户余额不足, 无法实现转账")
	}

	svi = svi - s
	tvi = tvi + s

	// 将修改之后的源账户与目标账户的状态保存至账本中
	err = stub.PutState(source, []byte(strconv.Itoa(svi)))
	if err != nil {
		return  shim.Error("保存转账后的源账户状态失败")
	}

	err = stub.PutState(target, []byte(strconv.Itoa(tvi)))
	if err != nil {
		return  shim.Error("保存转账后的目标账户状态失败")
	}

	return shim.Success([]byte("转账成功"))

}




