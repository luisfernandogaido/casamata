package main

import (
	"log"
	"os"

	"github.com/luisfernandogaido/casamata/img"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("esperados os subcomandos: 'img'")
	}
	switch os.Args[1] {
	case "img":
		img.Img()
	default:
		log.Fatal("digite um comando válido")
	}
}
