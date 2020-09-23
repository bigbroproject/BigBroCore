package main

import (
	"github.com/moneye/internal/models"
	"github.com/moneye/internal/protocols"
	"github.com/moneye/internal/utilities"
	"log"
)

func main() {

	conf, err := models.ConfigFromFile("config/conf.yml")
	if err != nil {
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
				err = utilities.PrintStatus(&service, &protocol, https.CheckService())
				break
			case "icmp":
				icmp := protocols.Icmp{
					protocol,
				}
				err = utilities.PrintStatus(&service, &protocol, icmp.CheckService())
				break
			case "icmp6":
				icmp6 := protocols.Icmp6{
					protocol,
				}
				err = utilities.PrintStatus(&service, &protocol, icmp6.CheckService())
				break
			}

		}

	}
}
