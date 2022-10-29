// Package pg lives in repository dir and represents adapter layer
// which enables interaction through a specific port and with a certain technology.
// in this case pg will act for CRUD operations with postgres.
package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Standard blanc import for pq.
	"github.com/rusli4k/fevo/app/entities"
)

// Repo wraps a database handle.
type Repo struct {
	DB *sql.DB
}

// NewRepo will initialize new instance of Repo.
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

// AddTransaction method implements adding one transaction into the database.
// It get one exemplar of transaction, return error.
func (r Repo) AddTr(tr entities.Transaction) error {
	const sqlStatement = `INSERT INTO "transaction" (id, request_id, terminal_id, partner_object_id,
		amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input,
		date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, 
		payee_bank_mfo, payee_bank_account, payment_narrative)
 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`

	if _, err := r.DB.Exec(sqlStatement, tr.ID, tr.RequestID, tr.TerminalID, tr.PartnerObjectID,
		tr.AmountTotal, tr.AmountOriginal, tr.CommissionPS, tr.CommissionClient, tr.CommissionProvider, tr.DateInput,
		tr.DatePost, tr.Status, tr.PaymentType, tr.PaymentNumber, tr.ServiceID, tr.Service, tr.PayeeID, tr.PayeeName,
		tr.PayeeBankMFO, tr.PayeeBankAccount, tr.PaymentNarrative); err != nil {
		return fmt.Errorf("error inserting to DB: %w", err)
	}

	return nil
}

// // GetTrByID get ID, returns single entity and error.
// func (r Repo) GetTrByID(id uint) (entities.Transaction, error) { return entities.Transaction{}, nil }

// // GetTrByTerminalID
// func (r Repo) GetTrByTerminalID()

// // GetTrByStatus
// func (r Repo) GetTrByStatus() {}

// // GetTrByPaymentType
// func (r Repo) GetTrByPaymentType() {}

// // GetTrByDataPost
// func (r Repo) GetTrByDataPost() {}

// // GetTrByPaymentNarrative
// func (r Repo) GetTrByPaymentNarrative() {}

// // func (r Repo) GetUsers(offset, limit, sort string) ([]entities.User, error) {
// // 	const sqlStatement = `
// // 			SELECT id, first_name, last_name, created_at
// // 			FROM "user"
// // 			WHERE deleted_at IS NULL
// // 			ORDER BY $1
// // 			OFFSET $2
// // 			LIMIT $3;
// // 	`
// // пошук по transaction_id
// // пошук по terminal_id (можливість вказати декілька одночасно id)
// // пошук по status (accepted/declined)
// // пошук по payment_type (cash/card)
// // пошук по date_post по періодам (from/to), наприклад: from 2022-08-12, to 2022-09-01 повинен повернути всі транзакції за вказаний період
// // пошук по частково вказаному payment_narrative

// // 	rows, err := r.DB.Query(sqlStatement, sort, offset, limit)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("error occurred while executing query: %w", err)
// // 	}

// // 	defer rows.Close()

// // 	var users []entities.User

// // 	for rows.Next() {
// // 		var user entities.User

// // 		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedAt)
// // 		if err != nil {
// // 			return nil, fmt.Errorf("error occurred while scaning object from query: %w", err)
// // 		}

// // 		users = append(users, user)
// // 	}

// // 	if err = rows.Err(); err != nil {
// // 		return nil, fmt.Errorf("error occurred during iteration: %w", err)
// // 	}

// // 	return users, nil
// // }
