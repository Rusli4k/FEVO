package rest

// TransactionHandler is Transaction HTTP handler
// which consist of embedded TransactionUsecase interface.
type TAHandler struct {
	usecase TAUsecase
}

// NewTAHandler will return a new instance
// of TAHandler struct accepting TransactionUsecase interface.
func NewTAHandler(usecase TAUsecase) TAHandler {
	return TAHandler{
		usecase: usecase,
	}
}
