package main

import (
	"blog/server/config"
	"blog/server/routes"
)

func main() {
	_ = routes.InitRouter().Run(config.HttpPort)

}
