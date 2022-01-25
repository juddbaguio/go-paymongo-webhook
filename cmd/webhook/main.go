package main

import (
	"log"
	"os"

	"github.com/juddbaguio/go-paymongo-webhook/pkg/app"
)

func main() {
	webhook := app.InitWebhookServer()

	if err := webhook.StartApp(); err != nil {
		log.Printf("error: %v", err.Error())
		os.Exit(1)
	}
}
