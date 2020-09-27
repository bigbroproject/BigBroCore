package protocols

import (
	"github.com/bigbroproject/bigbrocore/models"
	"net/http"
	"strconv"
	"time"
)

type Https struct {}

// If error is nil, then service is up
func (https Https) CheckService(Protocol models.Protocol) error {
	// CHECK
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	url := "https://" + Protocol.Server
	if Protocol.Port == 0 {
		url += ":443"
	} else {
		url += ":"+strconv.Itoa(Protocol.Port)
	}
	_, err := client.Get(url)
	return err
}
