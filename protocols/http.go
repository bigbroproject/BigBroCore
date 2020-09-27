package protocols

import (
	"github.com/bigbroproject/bigbrocore/models"
	httplib "net/http"
	"strconv"
	"time"
)

type Http struct {
}

// If error is nil, then service is up
func (http Http) CheckService(Protocol models.Protocol) error {
	// CHECK
	tr := &httplib.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &httplib.Client{Transport: tr}
	url := "http://" + Protocol.Server
	if Protocol.Port == 0 {
		url += ":80"
	} else {
		url += ":" + strconv.Itoa(Protocol.Port)
	}
	_, err := client.Get(url)
	return err
}
