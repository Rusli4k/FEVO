// Package provides core of the program
package entities

import "time"

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

// "id" INT  NOT NULL,
// "request_id" INT,
// "terminal_id" INT,
// "partner_object_id" INT,
// "amount_total" FLOAT,
// "amount_original" FLOAT,
// "commission_ps" FLOAT,
// "commission_client" FLOAT,
// "commission_provider" FLOAT,
// "date_input" Timestamp Without Time Zone,
// "date_post" Timestamp Without Time Zone,
// "status" VARCHAR(8),
// "payment_type" VARCHAR(4),
// "payment_number" VARCHAR(10),
// "service_id" INT,
// "service" VARCHAR(17),
// "payee_id" INT,
// "payee_name" VARCHAR(10),
// "payee_bank_mfo" INT,
// "payee_bank_account" VARCHAR(17),
// "payment_narrative" VARCHAR(255),
