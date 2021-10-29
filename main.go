package main

import (
	"log"
	"github.com/jnarinoc/twittor/bd"
	"github.com/jnarinoc/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
