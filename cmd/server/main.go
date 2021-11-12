package main

import (
	"flag"
	"log"

	"github.com/dollarkillerx/survive/internel/conf"
	"github.com/dollarkillerx/survive/internel/server"
)

var configFileName = flag.String("cfn", "configs/config.json", "name of configs file")

func init() {
	flag.Parse()

	conf.InitConfig(*configFileName)
}

func main() {
	ser := server.NewServer()
	if err := ser.Run(); err != nil {
		log.Fatalln(err)
	}
}
