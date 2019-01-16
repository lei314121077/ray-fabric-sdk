package demo

import (
	"encoding/json"
	"event"
	"fmt"
	"ray/fsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//查询
func (d *DemoController) DemoSer(certNo, name string, s *fsdk.Application) ([]byte, error){

	req := channel.Request{ChaincodeID: s.Setup.ChaincodeID, Fcn: "queryEduByCertNoAndName", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := s.Setup.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

//事物更新
func  (d *DemoController) DemoUpdateEvent(data Demo,s fsdk.Application )(string, error){

	eventID := "eventModifyEdu"
	reg, notifier := event.EventRegitser(s.Setup.Client, s.Setup.ChaincodeID, eventID)
	defer s.Setup.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: s.Setup.ChaincodeID, Fcn: "updateEdu", Args: [][]byte{b, []byte(eventID)}}
	respone, err := s.Setup.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = event.EventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil

}
