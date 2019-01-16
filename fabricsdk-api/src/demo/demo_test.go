package demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


var (
	level1 = 10000		//一万级
	level2 = 100000		//十万级
	level3 = 1000000	//百万级
	level4 = 10000000	//千万级
)

func TestApplication_DemoApi(t *testing.T) {

	start := time.Now()
	var wg sync.WaitGroup
	for i:=0; i<=100000; i++{

		go func(){


			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("开始时间为：", start)
	fmt.Println("耗时为：", int64( time.Since(start).Nanoseconds()))

}


func TestServiceSetup_DemoSer(t *testing.T) {

	start := time.Now()
	var wg sync.WaitGroup
	for i:=0; i<=100000; i++{

		go func(){


			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("开始时间为：", start)
	fmt.Println("耗时为：", int64( time.Since(start).Nanoseconds()))

}

