package order

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"ray/event"
	"ray/fsdk"
)


func (o *Order) AddOrderHistorySer(order Order, s *fsdk.Application)(string, error){

	eventID := "eventAddEdu"
	reg, notifier := event.EventRegitser(s.Setup.Client, s.Setup.ChaincodeID, eventID)
	defer s.Setup.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(order)
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: s.Setup.ChaincodeID, Fcn: "addHistory", Args: [][]byte{b, []byte(eventID)}}
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

func (o *Order) ModifyOrderHistorySer(order Order, s *fsdk.Application)(string, error){

	eventID := "eventModifyEdu"
	reg, notifier := event.EventRegitser(s.Setup.Client, s.Setup.ChaincodeID, eventID)
	defer s.Setup.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(order)
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: s.Setup.ChaincodeID, Fcn: "modifyHistory", Args: [][]byte{b, []byte(eventID)}}
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


func (o *Order) OrderHistorySer(workNo, userId,isHow string, s *fsdk.Application)([]byte, error){

	req := channel.Request{ChaincodeID: s.Setup.ChaincodeID, Fcn: "queryOrderHistry", Args: [][]byte{[]byte(workNo), []byte(userId), []byte(isHow)}}
	respone, err := s.Setup.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}




