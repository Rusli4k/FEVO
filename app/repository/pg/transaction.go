// Package pg represents adapter layer for operations with postgres.
package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq" // standard use of driver for postgres
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

// AddTr method implements adding single transaction into the database.
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

// GetTrByID get transaction ID, returns single entity and error.
func (r Repo) GetTrByID(id int) (*entities.Transaction, error) {
	const sqlStatement = `SELECT * FROM "transaction" WHERE id = $1`

	var tr entities.Transaction

	if err := r.DB.QueryRow(sqlStatement, id).Scan(&tr.ID, &tr.RequestID, &tr.TerminalID, &tr.PartnerObjectID,
		&tr.AmountTotal, &tr.AmountOriginal, &tr.CommissionPS, &tr.CommissionClient, &tr.CommissionProvider,
		&tr.DateInput, &tr.DatePost, &tr.Status, &tr.PaymentType, &tr.PaymentNumber, &tr.ServiceID,
		&tr.Service, &tr.PayeeID, &tr.PayeeName, &tr.PayeeBankMFO, &tr.PayeeBankAccount,
		&tr.PaymentNarrative); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting rows: %w", err)
	}

	return &tr, nil
}

// GetTrByTermID get some terminal ID`s, returns slice of entities and error.
func (r Repo) GetTrByTermID(id []int) ([]entities.Transaction, error) {
	var sqlStatement = `SELECT * FROM "transaction" WHERE terminal_id IN (`
	for _, v := range id {
		sqlStatement = sqlStatement + fmt.Sprint(v) + ","
	}

	sqlStatement = sqlStatement[:len(sqlStatement)-1]
	sqlStatement += ")"

	rows, err := r.DB.Query(sqlStatement)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting rows: %w", err)
	}
	defer rows.Close()

	var tas []entities.Transaction

	for rows.Next() {
		var tr entities.Transaction

		err := rows.Scan(&tr.ID, &tr.RequestID, &tr.TerminalID, &tr.PartnerObjectID,
			&tr.AmountTotal, &tr.AmountOriginal, &tr.CommissionPS, &tr.CommissionClient, &tr.CommissionProvider,
			&tr.DateInput, &tr.DatePost, &tr.Status, &tr.PaymentType, &tr.PaymentNumber, &tr.ServiceID,
			&tr.Service, &tr.PayeeID, &tr.PayeeName, &tr.PayeeBankMFO, &tr.PayeeBankAccount,
			&tr.PaymentNarrative)
		if err != nil {
			return nil, fmt.Errorf("error getting rows: %w", err)
		}

		tas = append(tas, tr)
	}

	return tas, nil
}

// GetTrByStatus get status string, returns slice of entities and error.
func (r Repo) GetTrByStatus(st string) ([]entities.Transaction, error) {
	var sqlStatement = `SELECT * FROM "transaction" WHERE status = $1`

	rows, err := r.DB.Query(sqlStatement, st)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting rows from db: %w", err)
	}
	defer rows.Close()

	var trs []entities.Transaction

	for rows.Next() {
		var tr entities.Transaction

		err := rows.Scan(&tr.ID, &tr.RequestID, &tr.TerminalID, &tr.PartnerObjectID,
			&tr.AmountTotal, &tr.AmountOriginal, &tr.CommissionPS, &tr.CommissionClient, &tr.CommissionProvider,
			&tr.DateInput, &tr.DatePost, &tr.Status, &tr.PaymentType, &tr.PaymentNumber, &tr.ServiceID,
			&tr.Service, &tr.PayeeID, &tr.PayeeName, &tr.PayeeBankMFO, &tr.PayeeBankAccount,
			&tr.PaymentNarrative)
		if err != nil {
			return nil, fmt.Errorf("error getting rows: %w", err)
		}

		trs = append(trs, tr)
	}

	return trs, nil
}

// GetTrByPayType get payType string, returns slice of entities and error.
func (r Repo) GetTrByPayType(st string) ([]entities.Transaction, error) {
	var sqlStatement = `SELECT * FROM "transaction" WHERE payment_type = $1`

	rows, err := r.DB.Query(sqlStatement, st)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting rows: %w", err)
	}
	defer rows.Close()

	var tas []entities.Transaction

	for rows.Next() {
		var tr entities.Transaction
		err := rows.Scan(&tr.ID, &tr.RequestID, &tr.TerminalID, &tr.PartnerObjectID,
			&tr.AmountTotal, &tr.AmountOriginal, &tr.CommissionPS, &tr.CommissionClient, &tr.CommissionProvider,
			&tr.DateInput, &tr.DatePost, &tr.Status, &tr.PaymentType, &tr.PaymentNumber, &tr.ServiceID,
			&tr.Service, &tr.PayeeID, &tr.PayeeName, &tr.PayeeBankMFO, &tr.PayeeBankAccount,
			&tr.PaymentNarrative)

		if err != nil {
			return nil, fmt.Errorf("error getting rows: %w", err)
		}

		tas = append(tas, tr)
	}

	return tas, nil
}

// GetTrByDataPost get map contains "from" and "to" values, format is YYYY-MM-DD HH:MM:SS,
// returns slice of entities and error.
func (r Repo) GetTrByDataPost(date map[string]time.Time) ([]entities.Transaction, error) {
	var sqlStatement = `SELECT * FROM "transaction" WHERE date_post >= $1 and date_post < $2`

	rows, err := r.DB.Query(sqlStatement, date["from"], date["to"])
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting rows: %w", err)
	}
	defer rows.Close()

	var tas []entities.Transaction

	for rows.Next() {
		var tr entities.Transaction
		err := rows.Scan(&tr.ID, &tr.RequestID, &tr.TerminalID, &tr.PartnerObjectID,
			&tr.AmountTotal, &tr.AmountOriginal, &tr.CommissionPS, &tr.CommissionClient, &tr.CommissionProvider,
			&tr.DateInput, &tr.DatePost, &tr.Status, &tr.PaymentType, &tr.PaymentNumber, &tr.ServiceID,
			&tr.Service, &tr.PayeeID, &tr.PayeeName, &tr.PayeeBankMFO, &tr.PayeeBankAccount,
			&tr.PaymentNarrative)

		if err != nil {
			return nil, fmt.Errorf("error getting rows: %w", err)
		}

		tas = append(tas, tr)
	}

	return tas, nil
}

// GetTrByPayNar get part of payNarrative as string, returns slice of entities and error.
func (r Repo) GetTrByPayNar(st string) ([]entities.Transaction, error) {
	var sqlStatement = `SELECT * FROM "transaction" WHERE payment_narrative LIKE $1`

	st = "%" + st + "%"

	rows, err := r.DB.Query(sqlStatement, st)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting rows: %w", err)
	}
	defer rows.Close()

	var tas []entities.Transaction

	for rows.Next() {
		var tr entities.Transaction

		err := rows.Scan(&tr.ID, &tr.RequestID, &tr.TerminalID, &tr.PartnerObjectID,
			&tr.AmountTotal, &tr.AmountOriginal, &tr.CommissionPS, &tr.CommissionClient, &tr.CommissionProvider,
			&tr.DateInput, &tr.DatePost, &tr.Status, &tr.PaymentType, &tr.PaymentNumber, &tr.ServiceID,
			&tr.Service, &tr.PayeeID, &tr.PayeeName, &tr.PayeeBankMFO, &tr.PayeeBankAccount,
			&tr.PaymentNarrative)
		if err != nil {
			return nil, fmt.Errorf("error getting rows: %w", err)
		}

		tas = append(tas, tr)
	}

	return tas, nil
}
