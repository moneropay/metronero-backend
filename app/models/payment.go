package models

import "time"

type Payment struct {
	// Invoice ID
	InvoiceId string `json:"invoice_id"`
	MerchantName string `json:"merchant_name"`
	Amount uint64 `json:"amount"`
	Fee uint64 `json:"fee"`

	// Merchant provided ID for this payment
	OrderId string `json:"order_id,omitempty"`

	// Possible statuses: pending, confirming, finished, cancelled, withdrawn
	Status string `json:"status"`
	LastUpdate time.Time `json:"last_update"`
}

type PaymentRequest struct {
	Payment
	AcceptUrl string `json:"accept_url"`
	CancelUrl string `json:"cancel_url"`
	CallbackUrl string `json:"callback_url"`

	// Index of subaddress that was used to accept this payment
	AddressIndex uint64 `json:"address_index"`

	// Callback data from MoneroPay
	CallbackData string `json:"callback_data,omitempty"`

	// Merchant provided extra data field
	ExtraData string `json:"extra_data,omitempty"`
}
