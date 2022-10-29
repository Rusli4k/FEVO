// Package usecase holds on services
// that build all together the business flow,
// and represents so-called business logic layer of the app
// usecase only depends on package ent.
package usecase

import (
	"github.com/rusli4k/fevo/app/entities"
)

// TARepository interface can be implemented
// in any kind of repositories like Postgres, MySQL etc.
type TARepository interface {
	AddTr(tr entities.Transaction) error
	// GetTrByID(id uint) (entities.Transaction, error)
	// GetTrByTerminalID()
	// GetTrByStatus()
	// GetTrByPaymentType()
	// GetTrByDataPost()
	// GetTrByPaymentNarrative()
}
