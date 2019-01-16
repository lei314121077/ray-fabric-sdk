package httpsdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)


type InitInfo struct {

	ChannelID     string							// 通道名称 ID
	ChannelConfig string							// 通道交易配置文件所在路径
	OrgName      string								// 组织的名称
	OrgAdmin       string							// 组织的管理员名称
	OrdererOrgName    string						// orderer 的名称
	OrgResMgmt *resmgmt.Client						// 资源管理端实例

	ChaincodeID    string							// 链码ID （即链码名称）
	ChaincodeGoPath    string						// 系统的GOPATH路径
	ChaincodePath    string							// 链码 源代码路径
	UserName    string								// 组织用户名称
	//SDK           *fabsdk.FabricSDK                	//SDK实例

}







