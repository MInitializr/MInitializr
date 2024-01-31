package initializers

import (
	"log"
	"example.com/minitializr/types"
)

type BaseIntializer struct {
	ServiceName string
	Service types.MIService
}

func (BaseIntializer) Initialize(){
	log.Println("BaseIntializer")
}