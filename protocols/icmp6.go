package protocols

import (
	"github.com/moneye/models"
	"os/exec"
)

type Icmp6 struct {
}

// If error is nil, then service is up
func (icmp Icmp6) CheckService(	Protocol models.Protocol) error {
	// CHECK
	_, err := exec.Command("ping6", Protocol.Server, "-c1").Output()
	//log.Printf("%s", out)
	return err
}
