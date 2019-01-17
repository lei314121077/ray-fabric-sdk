package fsdk

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"os"
)


var App Application

//TODO https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/api
const (
	ChaincodeVersion     = "1.4"      				//指定Chaincode 版本
	configFile           = "config.yaml"   			// config.yaml 文件
	initialized          = false
	DataExchangeCenterCC = "DataExchangeCenter" 	//ChainCode 名称
	channdelId 			 = "A0000000000001"
)


/*
@name 安装和序列化链码
@param sdk fabricSDK句柄
@param info SDK 配置信息对象
@return channel.client  error
*/
func InstallAndInstantiateCC(sdk *fabsdk.FabricSDK, info *InitInfo) (*channel.Client, error) {

	fmt.Println("开始安装链码......")
	// 创建新的golang chaincode包
	ccPkg, err := gopackager.NewCCPackage(info.ChaincodePath, info.ChaincodeGoPath)
	if err != nil {
		return nil, fmt.Errorf("创建链码包失败: %v", err)
	}

	//包含安装链代码请求参数
	installCCReq := resmgmt.InstallCCRequest{Name: info.ChaincodeID, Path: info.ChaincodePath, Version: ChaincodeVersion, Package: ccPkg}

	// 允许管理员将链代码安装到节点的文件系统上
	_, err = info.OrgResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("安装链码失败: %v", err)
	}

	fmt.Println("指定的链码安装成功......")
	fmt.Println("开始实例化链码......")

	// 返回一个有效的策略
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.google.com"})

	instantiateCCReq := resmgmt.InstantiateCCRequest{Name: info.ChaincodeID, Path: info.ChaincodePath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}

	//使用可选的自定义选项（特定对等体，过滤的对等体，超时）实例化链码。如果未指定peer
	_, err = info.OrgResMgmt.InstantiateCC(info.ChannelID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("实例化链码失败: %v", err)
	}

	fmt.Println("链码实例化成功")

	clientChannelContext := sdk.ChannelContext(info.ChannelID, fabsdk.WithUser(info.UserName), fabsdk.WithOrg(info.OrgName))

	// 返回客户端实例。通道客户端可以查询链码，执行链码以及注册/取消注册特定通道上的链码事件。
	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("创建应用通道客户端失败: %v", err)
	}

	fmt.Println("通道客户端创建成功，可以利用此客户端调用链码进行查询或执行事务.")

	return channelClient, nil
}


//启动SDK
func startSdk()*channel.Client {
	fmt.Println("-----start sdk")
	initInfo := &InitInfo{

		ChannelID:     channdelId,
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/data-transfer-chaincode/.../channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.www.google.com",

		ChaincodeID:     DataExchangeCenterCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "transfer-chaincode",
		UserName:        "User1",

	}

	sdk, err := SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	defer sdk.Close()

	// 创建Channel
	if err := CreateChannel(sdk, initInfo); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// 安装链码
	channelClient, err := InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println("end start sdk",channelClient)
	return channelClient
}

func InitSdk()*Application{

	fmt.Println("开始启动+++")

	//启动SDK 判断channel句柄是否为nil
	if channelClient := startSdk(); channelClient != nil{
		serviceSetup := ServiceSetup{
			ChaincodeID: DataExchangeCenterCC,
			Client:      channelClient,
		}

		App := &Application{
			Setup: &serviceSetup,
		}

		if App.Setup.ChaincodeID != "" && App.Setup.Client != nil{
			fmt.Println("sdk启动成功！")
			return App
		}

		//启动https服务
	}else{
		fmt.Println("sdk启动失败!")
	}
}