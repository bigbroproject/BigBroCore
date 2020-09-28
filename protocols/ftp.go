package protocols

import (
	"github.com/bigbroproject/bigbrocore/models"
)

type FTP struct {
}

func (ftp FTP) CheckService(Protocol models.Protocol) error {
	//TODO
	//log.Println(Protocol.Customs["user"]+"-"+Protocol.Customs["password"])
	return nil
}
