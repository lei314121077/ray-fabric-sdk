package qtool


import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/**
@NAME GetListResult 获取查询区块中的结果集
@PARAM StateQueryIteratorInterface
@DESC StateQueryIteratorInterface 接口的扩展 返回一个string的byte数组形式
@RETURN []byte,error
**/
func GetListResult(resultsIterator shim.StateQueryIteratorInterface) ([]byte,error){

	defer resultsIterator.Close()

	// 返回值JSON数组定义
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bufferArrayMemberAlreadyWritten := false //控制开关

	for resultsIterator.HasNext(){

		// 返回一个Hash值
		queryResp, err := resultsIterator.Next()
		if err != nil{
			return nil, err
		}

		if bufferArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}

		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResp.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResp.Value))
		buffer.WriteString("}")
		bufferArrayMemberAlreadyWritten = true

	}

	buffer.WriteString("]")
	fmt.Printf("queryResult：\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

/**
@NAME GetHistoryListResult 获取历史区块中的结果集
@PARAM resultsIterator HistoryQueryIteratorInterface
@DESC HistoryQueryIteratorInterface 接口的扩展函数  返回一个string的byte数组形式
@RETURN []byte, error
**/
func GetHistoryListResult(resultsIterator shim.HistoryQueryIteratorInterface)([]byte, error){
	defer  resultsIterator.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		item,_:= json.Marshal( queryResponse)
		buffer.Write(item)
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Printf("queryResult:\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}


