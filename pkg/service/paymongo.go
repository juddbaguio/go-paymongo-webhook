package paymongo_webhook

import (
	"log"

	paymongo "github.com/juddbaguio/go-paymongo-webhook/pkg/domain"
)

func Listen(webhook_data paymongo.WebhookData) {
	log.Println(webhook_data)
	// Do some business logic here
}
