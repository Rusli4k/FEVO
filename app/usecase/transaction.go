package usecase

import (
	"fmt"

	"github.com/rusli4k/fevo/app/entities"
)

type Transaction struct {
	Repo TARepository
}

// NewUser is a famous  trick with accepting
// interfaces and returning struct.
func NewTA(repo TARepository) Transaction {
	return Transaction{Repo: repo}
}

// SaveTA represents business logic
// and will take care of creating user.
func (tr Transaction) UploadTr(transaction entities.Transaction) error {
	if err := tr.Repo.AddTr(transaction); err != nil {
		return fmt.Errorf("repository error: %w", err)
	}

	return nil
}

// func (t Transaction) GetTrByIDUC(id uint) (entities.Transaction, error) {
// 	return entities.Transaction{}, nil
// }

// func (t Transaction) GetTrByTerminalIDUC()

// func (t Transaction) GetTrByStatusUC()

// func (t Transaction) GetTrByPaymentTypeUC()

// func (t Transaction) GetTrByDataPostUC()

// func (t Transaction) GetTrByPaymentNarrativeUC()
