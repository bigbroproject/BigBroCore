package protocols

import (
	"github.com/moneye/models"
	"log"
)

type FTP struct {

}


func (ftp FTP) CheckService(Protocol models.Protocol) error {
	//TODO
	log.Println(Protocol.Customs["user"]+"-"+Protocol.Customs["password"])
	return nil
}



