package main

import (
	"log"
)

func main() {

	regInterfaces,err := Initialize()
	if err != nil {
		log.Fatal(err.Error())
	}
	//protocols.RegisterInterface(&regInterfaces,"ftp", protocols.FTP{})
	Start(regInterfaces)

}


