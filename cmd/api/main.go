package main

import (
	"log"
	"mncPaymentAPI/pkg/api"
)

func main() {
	app := api.Default()
	err := app.Start()
	if err != nil {
		log.Print(err)
		panic(err)
	}
}
