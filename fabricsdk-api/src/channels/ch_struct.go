package channels

import "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"

// Request 包含查询和执行一个调用交易的参数
type Request struct {
	ChaincodeID  string
	Fcn          string
	Args         [][]byte
	TransientMap map[string][]byte

	// InvocationChai包含元数据，某些选择服务实现使用元数据来选择满足调用链中所有链码的背书
	// 策略的背书节点
	// Each chaincode may also be associated with a set of private data collection names
	// which are used by some Selection Services (e.g. Fabric Selection) to exclude endorsers
	// that do NOT have read access to the collections.
	// The invoked chaincode (specified by ChaincodeID) may optionally be added to the invocation
	// chain along with any collections, otherwise it may be omitted.
	InvocationChain []*fab.ChaincodeCall
}

//Response包含执行和查询一个调用交易的响应参数
type Response struct {
	Proposal         *fab.TransactionProposal
	Responses        []*fab.TransactionProposalResponse
	TransactionID    fab.TransactionID
	TxValidationCode pb.TxValidationCode
	ChaincodeStatus  int32
	Payload          []byte
}
