package mac

import (
	"bytes"
	"net"
)

// GetMacAddr retorna un string de la dirección Mac de la interfaz acitva
func GetMacAddr() (mac string) { // puede devolver wired mac or wireless mac

	// net.Interfaces() devuelve  un Slice de tipo Interface de interfaces del equipo
	interfaces, err := net.Interfaces()

	if err != nil {
		mac = err.Error()
		return
	}
	// Si no hay error el Slice, este se analiza

	for _, interfaz := range interfaces {

		// Se analiza cuál de las interfaces está activa, y esta se asigna a var mac
		if interfaz.Flags&net.FlagUp != 0 && bytes.Compare(interfaz.HardwareAddr, nil) != 0 {

			// Se usa método de HardwareAddr.String() para recibir un string
			mac = interfaz.HardwareAddr.String()
			break
		}
	}

	return
}
