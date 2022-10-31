package rest

import (
	"time"

	"github.com/rusli4k/fevo/app/entities"
)

// TAUsecase represents transaction use-case layer.
type TAUsecase interface {
	UploadTrans(entities.Transaction) error
	GetTransByID(int) (*entities.Transaction, error)
	GetTransByTermID([]int) ([]entities.Transaction, error)
	GetTransByStatus(string) ([]entities.Transaction, error)
	GetTransByPayType(string) ([]entities.Transaction, error)
	GetTransByDataPost(map[string]time.Time) ([]entities.Transaction, error)

	// GetTrByPaymentNarrativeUC()
}
