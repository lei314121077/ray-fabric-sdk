package fsdkapi

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
}


type Application struct {
	Setup *ServiceSetup
}