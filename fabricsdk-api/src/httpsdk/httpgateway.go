package httpsdk

import (
	"crypto/tls"
	"flag"
	"demo"
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

	mux := http.NewServeMux()

	// 根目录
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("这是一个示例服务.\n"))
	})

	// 注册用户
	mux.HandleFunc("/reguser", func(w http.ResponseWriter, req *http.Request){})

	// demo
	d := demo.DemoController{}
	mux.HandleFunc("/demo", d.DemoApi)

	// user
	o := order.Order{}
	mux.HandleFunc("/addOrderApi", o.AddOrderHistoryApi)
	mux.HandleFunc("/modifyOrderApi", o.ModifyHistoryApi)
	mux.HandleFunc("/quseryHistoryApi", o.QueryHistoryApi)

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
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Fatal(srv.ListenAndServeTLS("./server.rsa.crt", "./server.rsa.key"))
}


