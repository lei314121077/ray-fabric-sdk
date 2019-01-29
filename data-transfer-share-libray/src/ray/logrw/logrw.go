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

type RayLog struct{
	log *log.Logger
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

	this.log = loger
	return this.log
}

func (this *RayLog) Debug()*log.Logger {
	this.Obj()
	this.log.SetPrefix(LEVEL_DEBUG)
	return this.log
}


func (this *RayLog) Info()*log.Logger {
	this.Obj()
	this.log.SetPrefix(LEVEL_INFO)
	return this.log
}

func (this *RayLog) Waring()*log.Logger {
	this.Obj()
	this.log.SetPrefix(LEVEL_Waring)
	return this.log
}

func (this *RayLog) Error()*log.Logger {
	this.Obj()
	this.log.SetPrefix(LEVEL_Error)
	return this.log
}

//
//func main(){
//	rlog := RayLog{nil, file}
//	rlog.Error().Printf("hello kitty!!!!!!")
//
//}
