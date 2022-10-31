// Package usecase holds on services
// that build all together the business flow,
// and represents so-called business logic layer of the app
// usecase only depends on package ent.
package usecase

import (
	"time"

	"github.com/rusli4k/fevo/app/entities"
)

// TARepository interface can be implemented
// in any kind of repositories like Postgres, MySQL etc.
type TARepository interface {
	AddTr(entities.Transaction) error
	GetTrByID(int) (*entities.Transaction, error)
	GetTrByTermID([]int) ([]entities.Transaction, error)
	GetTrByStatus(string) ([]entities.Transaction, error)
	GetTrByPayType(string) ([]entities.Transaction, error)
	GetTrByDataPost(map[string]time.Time) ([]entities.Transaction, error)
	// GetTrByStatus()
	// GetTrByPaymentType()
	// GetTrByDataPost()
	// GetTrByPaymentNarrative()
}
