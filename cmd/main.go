package main

import (
	"log"
	"os"

	"github.com/whiteagleinc-meli/curso-testing-go/clase1"
)

func main() {
	suma := clase1.Add(2, 3)

	if suma != 5 {
		log.Printf("Error: Se esperaba %d, se obtuvo %d", 5, add)
		os.Exit(1)
	}

	log.Println("Todo salio bien")
}
