package protocols

import (
	"github.com/moneye/models"
	"os/exec"
)

type Icmp struct {
}

// If error is nil, then service is up
func (icmp Icmp) CheckService(	Protocol models.Protocol) error {
	// CHECK
	_, err := exec.Command("ping", Protocol.Server, "-c1").Output()
	//log.Printf("%s", out)
	return err
}
