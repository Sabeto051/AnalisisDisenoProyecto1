package iduser

import "os/user"

// GetUsername retorna el nombre de usuario de la maquina
func GetUsername() (username string) {
	structuser, error := user.Current() //Obtiene el usuario que esta en sesion
	if error != nil {
		username = "UserError" //Manejo de errores: retorna "UserError" como el nombre del usuario si por
		//algun motivo no hay ningun usuario en sesion
	}
	username = structuser.Username
	return
}
