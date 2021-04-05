package main

import (
	"blog/src/config"
	"blog/src/routes"
)

func main() {
	_ = routes.InitRouter().Run(config.HttpPort)
}
