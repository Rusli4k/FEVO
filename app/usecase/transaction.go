package usecase

import (
	"fmt"
	"time"

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
func (tr Transaction) UploadTrans(transaction entities.Transaction) error {
	if err := tr.Repo.AddTr(transaction); err != nil {
		return fmt.Errorf("repository error: %w", err)
	}

	return nil
}

func (tr Transaction) GetTransByID(id int) (*entities.Transaction, error) {
	ts, err := tr.Repo.GetTrByID(id)
	if ts == nil {
		if err != nil {
			return nil, fmt.Errorf("repository error: %w", err)
		}
		return nil, nil
	}
	return ts, nil
}

func (tr Transaction) GetTransByTermID(id []int) ([]entities.Transaction, error) {
	ts, err := tr.Repo.GetTrByTermID(id)
	if ts == nil {
		if err != nil {
			return nil, fmt.Errorf("repository error: %w", err)
		}
		return nil, nil
	}
	return ts, nil
}

func (tr Transaction) GetTransByStatus(st string) ([]entities.Transaction, error) {
	ts, err := tr.Repo.GetTrByStatus(st)
	if ts == nil {
		if err != nil {
			return nil, fmt.Errorf("repository error: %w", err)
		}
		return nil, nil
	}
	return ts, nil
}

func (tr Transaction) GetTransByPayType(st string) ([]entities.Transaction, error) {
	ts, err := tr.Repo.GetTrByPayType(st)
	if ts == nil {
		if err != nil {
			return nil, fmt.Errorf("repository error: %w", err)
		}
		return nil, nil
	}
	return ts, nil
}

func (tr Transaction) GetTransByDataPost(date map[string]time.Time) ([]entities.Transaction, error) {
	ts, err := tr.Repo.GetTrByDataPost(date)
	if ts == nil {
		if err != nil {
			return nil, fmt.Errorf("repository error: %w", err)
		}
		return nil, nil
	}
	return ts, nil
}

// // func (t Transaction) GetTrByPaymentNarrativeUC()
