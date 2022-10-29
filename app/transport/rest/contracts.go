package rest

import (
	"github.com/rusli4k/fevo/app/entities"
)

// TAUsecase represents transaction use-case layer.
type TAUsecase interface {
	UploadTr(tr entities.Transaction) error
	// GetTrByIDUC(id uint) (entities.Transaction, error)
	// GetTrByTerminalIDUC()
	// GetTrByStatusUC()
	// GetTrByPaymentTypeUC()
	// GetTrByDataPostUC()
	// GetTrByPaymentNarrativeUC()
}

// GetTrByID(id uint) (entities.Transaction, error)
// GetTrByTerminalID()
// GetTrByStatus()
// GetTrByPaymentType()
// GetTrByDataPost()
// GetTrByPaymentNarrative()
