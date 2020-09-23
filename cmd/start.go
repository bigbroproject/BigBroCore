package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/moneye/models"
	"github.com/moneye/process"
	"github.com/moneye/protocols"
	"github.com/moneye/utilities"
	"log"
	"time"
)

func Initialize() (map[string]protocols.ProtocolInterface,error) {
	return protocols.DefaultRegisteredInterfaces()
}

func Start(registeredInterfaces map[string]protocols.ProtocolInterface){
	conf, err := models.ConfigFromFile("config/conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*conf)
	processesChannel := make(chan string)

	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range conf.Services {
		service := conf.Services[i]
		for i2 := range service.Protocols {
			protocol := service.Protocols[i2]
			protocolInterface := registeredInterfaces[protocol.Type]

			if protocols.IsRegistered(&registeredInterfaces,protocol.Type) {
				proc := process.NewProcess(func() {
					utilities.PrintStatus(&service, &protocol, protocolInterface.CheckService(protocol))
				}, processesChannel)
				process.ScheduleProcess(proc, protocol.Interval)
			} else {
				red := color.New(color.FgRed).SprintFunc()
				log.Println(fmt.Sprintf("[%s] [%s] [%s] [%s - %s - %s] An error as accured: %s", red("ERRO"), time.Now().Format(time.RFC3339), service.Name, protocol.Type, protocol.Server, protocol.Port,"ProtocolInterface "+protocol.Type+" not registered!"))
			}

		}
	}
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
