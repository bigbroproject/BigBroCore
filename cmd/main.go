package main

import (
	"github.com/moneye/internal/models"
	"github.com/moneye/internal/protocols"
	"log"
)


func main() {

	conf,err := models.ConfigFromFile("config/conf.yml")
	if(err!=nil){
		log.Fatal(err)
	}
	log.Println(*conf)

	for i := range conf.Services {
		service := conf.Services[i]
		for i2 := range service.Protocols {
			protocol := service.Protocols[i2]
			switch protocol.Type {
			case "https":
				https := protocols.Https{
				protocol,
				}
				err = https.CheckService()
				if err !=nil {
					log.Printf("[%s - %s - %s - %d] Status ERROR. %s",service.Name,protocol.Type,protocol.Server,protocol.Port,err.Error())

				}else {
					log.Printf("[%s - %s - %s - %d] Status OK.",service.Name,protocol.Type,protocol.Server,protocol.Port)
				}
			}
		}
	}
}

