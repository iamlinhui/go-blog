package main

import (
	"go-blog/config"
	"go-blog/routes"
)

func main() {
	_ = routes.InitRouter().Run(config.HttpPort)
}
