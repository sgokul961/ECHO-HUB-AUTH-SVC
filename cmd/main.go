package main

import (
	"log"

	"github.com/sgokul961/echo-hub-auth-svc/pkg/config"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/wire"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed at config", err)
	}
	server, err := wire.InitApi(c)

	if err != nil {
		log.Fatalln("error in intilizing server", err)
	} else {
		server.Start(c)
	}

}
