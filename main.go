package main

import(
	"proyectoanal/server"
	"proyectoanal/mac"
	"proyectoanal/idprocessor"
	"proyectoanal/iduser"
)

func main() {

	token := GetMacAddr()
	server.CreateLocalHost(token)
}
