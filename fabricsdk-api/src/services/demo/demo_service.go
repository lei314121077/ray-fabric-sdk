package demo

import (
	"fmt"
	"ray/fsdkapi"
)

func (s *fsdkapi.ServiceSetup)DemoSer(demo Demo)(string, error){

	fmt.Println("server process!")

	return "", nil

}

