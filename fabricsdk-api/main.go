package main

import (
	"fmt"
	"httpsdk"
)

func main() {
	fmt.Println("开始启动+++")
	httpsdk.HttpStart()
	//if res := fsdk.InitSdk();res != nil {
	//	httpsdk.HttpStart()
	//	fmt.Println("服务器启动成功！")
	//}else{
	//	fmt.Println("服务器启动失败！")
	//}

}