package server

import (
	"encoding/json"
	"net/http"
)

// Mensaje ss
type Mensaje struct {
	MacAddress    string `json:"macAddress"`
	IDprocesador  string `json:"idProcesador"`
	NombreUsuario string `json:"nombreUser"`
}

/* func handler(w http.ResponseWriter, r *http.Request) {
	// librería io para convertir  datos en bytes
	io.WriteString(w, "Hola servidor Go\n")
} */

// CreateLocalHost crea un servidor local
func CreateLocalHost(mac string, idprocessor string, usr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		mensaje := Mensaje{mac, idprocessor, usr}

		json.NewEncoder(w).Encode(mensaje)
	})
	http.ListenAndServe(":8000", nil)
}
