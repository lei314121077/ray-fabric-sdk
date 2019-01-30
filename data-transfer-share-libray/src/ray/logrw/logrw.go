package logrw

import (
	"log"
	"os"
	"time"
)

var(
	filename = "./" + time.Now().Format("20060102") + ".log"
)

const(
	LEVEL_DEBUG = "DEBUG :"
	LEVEL_INFO = "INFO :"
	LEVEL_Waring = "DARING :"
	LEVEL_Error = "ERROR :"

)

//var logobj *log.Logger

type RayLog struct{
	file string
}

func (this *RayLog) Obj()*log.Logger{

	fname := ""
	if this.file == ""{
		fname = filename
	}else{
		fname = this.file
	}

	logFile, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}

	loger := log.New(logFile, "Start", log.Ldate|log.Ltime|log.Lshortfile)
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return loger
}

func (this *RayLog) Debug()*log.Logger {
	resp := this.Obj()
	resp.SetPrefix(LEVEL_DEBUG)
	return resp
}


func (this *RayLog) Info()*log.Logger {
	resp := this.Obj()
	resp.SetPrefix(LEVEL_INFO)
	return resp
}

func (this *RayLog) Waring()*log.Logger {
	resp := this.Obj()
	resp.SetPrefix(LEVEL_Waring)
	return resp
}

func (this *RayLog) Error()*log.Logger {
	resp := this.Obj()
	resp.SetPrefix(LEVEL_Error)
	return resp
}

//
//func main(){
//	rlog := RayLog{nil, file}
//	rlog.Error().Printf("hello kitty!!!!!!")
//
//}
