package utilities

import (
	"github.com/fatih/color"
	"github.com/moneye/internal/models"
	"log"
	"strconv"
	"time"
)

/*
PrintStatus prints the result of a CheckService result and returns the same error, if there was, otherwise nil will be returned.
*/
func PrintStatus(service *models.Service, protocol *models.Protocol, err error) error {

	now := time.Now()
	port := strconv.Itoa(protocol.Port)
	if port == "0" {
		port = "No port"
	}
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		log.Printf("[%s] [%s] [%s] [%s - %s - %s] An error as accured: %s", red("ERRO"), now.Format(time.RFC3339), service.Name, protocol.Type, protocol.Server, port, err.Error())
	} else {
		green := color.New(color.FgHiGreen).SprintFunc()
		log.Printf("[ %s ] [%s] [%s] [%s - %s - %s] Service seems OK.", green("OK"), now.Format(time.RFC3339), service.Name, protocol.Type, protocol.Server, port)
	}
	return err
}
