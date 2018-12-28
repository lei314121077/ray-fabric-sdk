package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"service"
)

func main() {

	if err := shim.Start(new(common.DataExchangeCenter)); err != nil {
		fmt.Printf("启动链码失败...！错误: %s", err)
	}
	fmt.Println("Chaincode成功启动!")

}








































