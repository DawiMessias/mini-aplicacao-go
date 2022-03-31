package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	app := app.Gerar()

	fmt.Println("Logando conforme parametro ", app.Usage)

	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
