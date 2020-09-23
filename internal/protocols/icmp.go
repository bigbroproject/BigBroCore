package protocols

import (
	"github.com/moneye/internal/models"
	"os/exec"
)

type Icmp struct {
	Protocol models.Protocol
}

// If error is nil, then service is up
func (icmp *Icmp) CheckService() error {
	// CHECK
	_, err := exec.Command("ping", icmp.Protocol.Server, "-c1").Output()
	//log.Printf("%s", out)
	return err
}
