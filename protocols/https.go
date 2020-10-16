package protocols

import (
	"github.com/bigbroproject/bigbrocore/models"
	"net/http"
	"strconv"
	"time"
)

type Https struct {}

var clientHttps =  &http.Client{Timeout: time.Second * 10}
// If error is nil, then service is up
func (https Https) CheckService(Protocol models.Protocol) error {
	// Init
	url := "https://" + Protocol.Server
	if Protocol.Port == 0 {
		url += ":443"
	} else {
		url += ":" + strconv.Itoa(Protocol.Port)
	}
	resp, err := clientHttps.Get(url)
	if resp != nil {
		_ = resp.Body.Close()
		resp = nil
	}
	return err
}
