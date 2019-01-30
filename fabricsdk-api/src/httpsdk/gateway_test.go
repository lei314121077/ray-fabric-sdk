package httpsdk

import (
	"io"
	"os"
	"log"
	"testing"
	"net/http"
	"crypto/tls"

)

func TestRunHttpStart(t *testing.T) {
	HttpStart()
}


func TestTlsCheck(t *testing.T){

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}

	if resp, err := c.Get("https://localhost:8000"); err != nil {
		log.Fatal("http.Client.Get: ", err)
	} else {
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	}
}








