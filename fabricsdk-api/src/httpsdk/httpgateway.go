package httpsdk

import (
	"crypto/tls"
	"demo"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"order"
	"ray/logrw"
	"time"
)


var(
	addr = flag.String("http", "0.0.0.0:8000", "endpoint of YourService")

	file = "./" + time.Now().Format("20060102") + ".log"
)

var (
	users []User
	Rlog logrw.RayLog
)

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


func init() {
	//Rlog = logrw.RayLog{file}
	admin := User{LoginName:"Ray", Password:"123456", IsAdmin:"T"}
	alice := User{LoginName:"ChainDesk", Password:"123456", IsAdmin:"T"}
	bob := User{LoginName:"alice", Password:"123456", IsAdmin:"F"}
	jack := User{LoginName:"bob", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, alice)
	users = append(users, bob)
	users = append(users, jack)

}

func HttpStart(){


	flag.Parse()
	router := mux.NewRouter()

	// tls验证
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	srv := &http.Server{
		Addr:         *addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: router, // Pass our instance of gorilla/mux in.
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	}).Methods("GET")

	// demo
	d := demo.DemoController{}
	router.HandleFunc("/demo", d.DemoApi).Methods("POST")

	// user
	o := order.Order{}
	router.HandleFunc("/addOrderApi", o.AddOrderHistoryApi).Methods("POST")
	router.HandleFunc("/modifyOrderApi", o.ModifyHistoryApi).Methods("POST")
	router.HandleFunc("/quseryHistoryApi", o.QueryHistoryApi).Methods("POST")


	// test dev model
	//basePath := "/home/ray/go/data-transfer-chaincode/fabricsdk-api/src/httpsdk/"
	// output dev model
	basePath := "/root/go/data-transfer-chaincode/fabricsdk-api/src/httpsdk/"
	// just is http
	//log.Fatal(http.ListenAndServe(*addr, router))
	Rlog.Debug().Printf("debug",srv.ListenAndServeTLS(fmt.Sprintf("%v%v", basePath, "tls.crt"), fmt.Sprintf("%v%v", basePath, "tls.key")))
	log.Fatal(srv.ListenAndServeTLS(fmt.Sprintf("%v%v", basePath, "tls.crt"), fmt.Sprintf("%v%v", basePath, "tls.key")))
	fmt.Println("ok start http and tls successfull !")

}

