package logrw

import "testing"

var (
	raylog RayLog
	ffname string
)
func TestRayLog_Obj(t *testing.T) {

}

func TestRayLog_Debug(t *testing.T) {
	ffname = "testdebug.log"
	raylog = RayLog{nil, ffname}
	raylog.Debug().Printf("hello DEBUG!!!!!!")
}

func TestRayLog_Info(t *testing.T) {
	ffname = "testinfo.log"
	raylog = RayLog{nil, ffname}
	raylog.Info().Printf("hello INFO!!!!!!")
}


func TestRayLog_Waring(t *testing.T) {
	ffname = "testwaring.log"
	raylog = RayLog{nil, ffname}
	raylog.Debug().Printf("hello WARING!!!!!!")
}

func TestRayLog_Error(t *testing.T) {
	ffname = "testerror.log"
	raylog = RayLog{nil, ffname}
	raylog.Debug().Printf("hello ERROR!!!!!!")
}