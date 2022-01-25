package controllers

import (
	"encoding/json"
	"net/http"

	paymongo "github.com/juddbaguio/go-paymongo-webhook/pkg/domain"
	paymongo_webhook "github.com/juddbaguio/go-paymongo-webhook/pkg/service"
)

func ListenPaymongoWebhook(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	var WebhookData paymongo.WebhookData

	err := json.NewDecoder(r.Body).Decode(&WebhookData)

	if err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	paymongo_webhook.Listen(WebhookData)
}
