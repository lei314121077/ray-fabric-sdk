package user

import (
	"fmt"
	"ray/fsdk"
)

func (u *User) UserDemoSer(name, passwd string, s *fsdk.Application)(string, error){

	fmt.Println("User Demo Server")

	return "", nil
}





