package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/survive/pkg/request"
	"github.com/dollarkillerx/survive/pkg/utils"
	"github.com/dollarkillerx/urllib"
)

var serverUri = flag.String("uri", "127.0.0.1:8571", "server rui")
var token = flag.String("token", "ea_token", "server token")

func init() {
	flag.Parse()
}

func main() {
	uri := fmt.Sprintf("http://%s/heartbeat", *serverUri)
	fmt.Println("server: ", uri)
	info, err := utils.GetCpuInfo()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for {
		cpuUse, err := utils.CpuUse()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}

		memUse, err := utils.MemUse()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}

		_, _, err = urllib.Post(uri).SetHeader("token", *token).SetJsonObject(request.HeartbeatRequest{
			Time:           time.Now().UnixNano() / 1e6,
			CPU:            info,
			CPUUsedPercent: cpuUse,
			MemUsedPercent: memUse,
		}).ByteOriginal()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(time.Second * 2)
	}
}
