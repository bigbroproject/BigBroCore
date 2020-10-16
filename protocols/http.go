package protocols

import (
	"github.com/bigbroproject/bigbrocore/models"
	"net/http"
	"strconv"
	"time"
)

type Http struct {}


var clientHttp =  &http.Client{Timeout: time.Second * 10}
// If error is nil, then service is up
func (httpVar Http) CheckService(Protocol models.Protocol) error {
	// CHECK
	url := "http://" + Protocol.Server
	if Protocol.Port == 0 {
		url += ":80"
	} else {
		url += ":" + strconv.Itoa(Protocol.Port)
	}
	resp, err := clientHttp.Get(url)
	if resp != nil {
		_ = resp.Body.Close()
		resp = nil
	}
	return err
}
