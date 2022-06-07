package main

import (
	"log"

	"github.com/luxarasis/twittor/bd"
	"github.com/luxarasis/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin Conexion a la base de datos")
		return
	}

	handlers.Managers()
}
