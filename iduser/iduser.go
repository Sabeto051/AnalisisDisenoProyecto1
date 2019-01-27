package iduser

import "os/user"

// GetUsername retorna el nombre de usuario de la maquina
func GetUsername() (username string){
	structuser, error := user.Current() //Obtiene el usuario que esta en sesion
	if error != nil {
		username = "UserError"
	}
	username = structuser.Username	
	return
}