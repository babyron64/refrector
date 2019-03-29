package main

import (
	"github.com/kanatatsu64/refrector/mongodb"
	"github.com/kanatatsu64/refrector/server"
)

func main() {
	mongodb.InitMongoDB()
	server.InitServer(server.GetPort())
}
