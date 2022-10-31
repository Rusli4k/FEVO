package rest

import (
	"github.com/rusli4k/fevo/app/entities"
)

// TAUsecase represents transaction use-case layer.
type TAUsecase interface {
	UploadTr(entities.Transaction) error
	GetTransByID(int) (*entities.Transaction, error)
	GetTransByTermID([]int) ([]entities.Transaction, error)

	// GetTrByStatusUC()
	// GetTrByPaymentTypeUC()
	// GetTrByDataPostUC()
	// GetTrByPaymentNarrativeUC()
}
