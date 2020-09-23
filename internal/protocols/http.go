package protocols

import (
	"github.com/moneye/internal/models"
	"net/http"
	"strconv"
	"time"
)

type Http struct {
	Protocol models.Protocol
}

// If error is nil, then service is up
func (https *Http) CheckService() error {
	// CHECK
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	url := "http://" + https.Protocol.Server
	if https.Protocol.Port == 0 {
		url += ":80"
	} else {
		url += ":" + strconv.Itoa(https.Protocol.Port)
	}
	_, err := client.Get(url)
	return err
}
