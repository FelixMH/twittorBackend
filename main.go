package main

import (
	"log"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No se pudo hacer conexión a la BD")
		return
	}
	handlers.Manejadores()
}
