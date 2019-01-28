package main

<<<<<<< HEAD
import(
	"proyectoanal/server"
	"proyectoanal/mac"
	"proyectoanal/idprocessor"
	"proyectoanal/iduser"
)

func main() {

	token := GetMacAddr()
	server.CreateLocalHost(token)
=======
import "github.com/santiagobedoya/analisisDisenoSoftware/proyecto1/user"

func main() {

	/* token := mac.GetMacAddr()

	server.CreateLocalHost(token) */
	user.GetUser()
>>>>>>> santiago
}
