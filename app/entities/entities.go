package entities

import "time"

// Transaction is main entity of program.
type Transaction struct {
	ID                 int       `json:"transaction_id"`
	RequestID          int       `json:"request_id"`
	TerminalID         int       `json:"terminal_id"`
	PartnerObjectID    int       `json:"partner_object_id"`
	AmountTotal        float64   `json:"amount_total"`
	AmountOriginal     float64   `json:"amount_original"`
	CommissionPS       float64   `json:"commission_ps"`
	CommissionClient   float64   `json:"commission_client"`
	CommissionProvider float64   `json:"commission_provider"`
	DateInput          time.Time `json:"date_input"`
	DatePost           time.Time `json:"date_post"`
	Status             string    `json:"status"`
	PaymentType        string    `json:"payment_type"`
	PaymentNumber      string    `json:"payment_number"`
	ServiceID          int       `json:"service_id"`
	Service            string    `json:"service"`
	PayeeID            int       `json:"payee_id"`
	PayeeName          string    `json:"payee_name"`
	PayeeBankMFO       int       `json:"payee_bank_mfo"`
	PayeeBankAccount   string    `json:"payee_bank_account"`
	PaymentNarrative   string    `json:"payment_narrative"`
}

// Filter contains all acceptable parameters for filtering Transactions.
type Filter struct {
	TransactionID    int
	TerminalID       []int
	Status           string
	PaymentType      string
	DatePost         map[string]time.Time
	PaymentNarrative string
}
