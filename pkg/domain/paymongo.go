package paymongo

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

type Billing struct {
	Address Address `json:"address"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
}

type Redirect struct {
	CheckoutUrl string `json:"checkout_url"`
	Failed      string `json:"failed"`
	Success     string `json:"success"`
}

type SourceAttributes struct {
	Amount    int64    `json:"amount"`
	Billing   Billing  `json:"billing"`
	Currency  string   `json:"currency"`
	LiveMode  bool     `json:"livemode"`
	Redirect  Redirect `json:"redirect"`
	Status    string   `json:"status"`
	Type      string   `json:"string"`
	CreatedAt int64    `json:"created_at"`
}

type EventTypeAttributes struct {
	Type     string           `json:"type"`
	LiveMode bool             `json:"live_mode"`
	Data     SourceAttributes `json:"data"`
}

type MainAttributes struct {
	Type     string              `json:"type"`
	LiveMode string              `json:"live_mode"`
	Data     EventTypeAttributes `json:"data"`
}

type WebhookData struct {
	Id              string         `json:"id"`
	Type            string         `json:"type"`
	Attributes      MainAttributes `json:"attributes"`
	PreviousData    *WebhookData   `json:"previous_data"`
	PendingWebhooks int            `json:"pending_webhooks"`
	CreatedAt       int64          `json:"created_at"`
	UpdatedAt       int64          `json:"updated_at"`
}
