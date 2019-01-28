package idprocessor

import (
	"strings"

	"github.com/jpoirier/cpu"
)

// GetProcessorID retorna el id del procesador
func GetProcessorID() (idprocessor string) {
	temp := cpu.ProcessorFamily

	idprocessor = strings.Replace(temp, " ", "_", -1)
	index := strings.LastIndex(idprocessor, "z")
	index++
	idprocessor = idprocessor[:index]
	return
}
