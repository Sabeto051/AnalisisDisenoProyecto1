package main

import (
	"github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/server"

	"github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/mac"
)

func main() {

	token := mac.GetMacAddr()

	server.CreateLocalHost(token)
}
