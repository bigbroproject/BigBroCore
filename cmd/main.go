package main

import (
	"github.com/moneye/internal/models"
	"github.com/moneye/internal/process"
	"github.com/moneye/internal/protocols"
	"github.com/moneye/internal/utilities"
	"log"
	"time"
)

func main() {

	conf, err := models.ConfigFromFile("config/conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*conf)
	processesChannel := make(chan string)

	for i := range conf.Services {
		service := conf.Services[i]
		for i2 := range service.Protocols {
			protocol := service.Protocols[i2]
			switch protocol.Type {
			case "https":
				https := protocols.Https{
					protocol,
				}
				proc := process.NewProcess(func() {
					utilities.PrintStatus(&service, &protocol, https.CheckService())
				}, processesChannel)
				process.ScheduleProcess(proc, protocol.Interval)
				//err = utilities.PrintStatus(&service, &protocol, https.CheckService())
				break
			case "http":
				http := protocols.Http{
					protocol,
				}
				proc := process.NewProcess(func() {
					utilities.PrintStatus(&service, &protocol, http.CheckService())
				}, processesChannel)
				process.ScheduleProcess(proc, protocol.Interval)

				//err = utilities.PrintStatus(&service, &protocol, http.CheckService())
				break
			case "icmp":
				icmp := protocols.Icmp{
					protocol,
				}
				proc := process.NewProcess(func() {
					utilities.PrintStatus(&service, &protocol, icmp.CheckService())
				}, processesChannel)
				process.ScheduleProcess(proc, protocol.Interval)
				//err = utilities.PrintStatus(&service, &protocol, icmp.CheckService())
				break
			case "icmp6":
				icmp6 := protocols.Icmp6{
					protocol,
				}
				proc := process.NewProcess(func() {
					utilities.PrintStatus(&service, &protocol, icmp6.CheckService())
				}, processesChannel)
				process.ScheduleProcess(proc, protocol.Interval)
				//err = utilities.PrintStatus(&service, &protocol, icmp6.CheckService())
				break
			}

		}

	}
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
