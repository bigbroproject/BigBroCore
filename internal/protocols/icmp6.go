package protocols

import (
	"github.com/moneye/internal/models"
	"os/exec"
)

type Icmp6 struct {
	Protocol models.Protocol
}

// If error is nil, then service is up
func (icmp *Icmp6) CheckService() error {
	// CHECK
	_, err := exec.Command("ping6", icmp.Protocol.Server, "-c1").Output()
	//log.Printf("%s", out)
	return err
}
