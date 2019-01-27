package main


import (
	"github.com/jpoirier/cpu"
)
// GetProcessorID retorna el id del procesador
func GetProcessorID() (idprocessor string){
	idprocessor = cpu.ProcessorFamily 
	return
}