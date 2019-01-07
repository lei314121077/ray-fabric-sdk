package demo

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"ray/fsdkapi"
	"../event"
)

func (s *fsdkapi.ServiceSetup) DemoSer(certNo, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: s.ChaincodeID, Fcn: "queryEduByCertNoAndName", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := s.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}



func (s *fsdkapi.ServiceSetup)DemoUpdateEvent(d, Demo)(string, error){

	eventID := "eventModifyEdu"
	reg, notifier := event.EventRegitser(s.Client, s.ChaincodeID, eventID)
	defer s.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: s.ChaincodeID, Fcn: "updateEdu", Args: [][]byte{b, []byte(eventID)}}
	respone, err := s.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = event.EventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil

}
