package httpsdk

import (
	"demo"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"order"
)


var(
	addr = flag.String("http", "localhost:8000", "endpoint of YourService")

)

var users []User

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


func init() {

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

	// tls验证
	//cfg := &tls.Config{
	//	MinVersion:               tls.VersionTLS12,
	//	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	//	PreferServerCipherSuites: true,
	//	CipherSuites: []uint16{
	//		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	//		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	//	},
	//}
	//
	//srv := &http.Server{
	//	Addr:         *addr,
	//	WriteTimeout: time.Second * 15,
	//	ReadTimeout:  time.Second * 15,
	//	IdleTimeout:  time.Second * 60,
	//	Handler: router, // Pass our instance of gorilla/mux in.
	//	TLSConfig:    cfg,
	//	TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	//}
	fmt.Println("ok start successfull !")

	log.Fatal(http.ListenAndServe(*addr, router))
	//log.Fatal(srv.ListenAndServeTLS("./server.rsa.crt", "./server.rsa.key"))

}

