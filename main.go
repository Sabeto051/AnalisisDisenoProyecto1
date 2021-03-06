package main

import (
	"github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/idprocessor"
	"github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/iduser"
	"github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/mac"
	"github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/server"
)

func main() {

	macaddr := mac.GetMacAddr()
	idprocessor := idprocessor.GetProcessorID()
	usr := iduser.GetUsername()
	server.CreateLocalHost(macaddr, idprocessor, usr)
}
