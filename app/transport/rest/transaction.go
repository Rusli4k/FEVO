package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rusli4k/fevo/app/entities"
	"github.com/rusli4k/fevo/pkg/parser"
)

// UploadTransaction will handle transactions upload from
// CSV file to repository.
func (th TAHandler) UploadTransactions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	r.ParseMultipartForm(32 << 20) // limit your max input length 32Mb
		file, _, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		ts, err := parser.CSVToTransactions(file)
		if err != nil {
			WriteJSONResponse(w, http.StatusUnprocessableEntity, Response{
				Message: "Error occurred while processing the file",
				Details: err.Error()})
			return
		}

		for _, v := range ts {
			if err := th.usecase.UploadTr(v); err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: "Error while adding to db: ",
					Details: err.Error()})
				return
			}
		}

		WriteJSONResponse(w, http.StatusOK, Response{
			Message: "Request processed with no errors.",
			Details: fmt.Sprint("Num of added rows:", len(ts))})
	})
}

// GetTransactions handle  filters in request.
func (ta TAHandler) GetTransactions() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var par entities.Filter
		var err error
		values := r.URL.Query()

		// Getting transaction ID from request.
		if value, ok := values[keyTaID]; ok {
			par.TransactionID, err = strconv.Atoi(value[0])
			if err != nil {
				WriteJSONResponse(w, http.StatusBadRequest, Response{
					Message: MSgBadURL,
					Details: err.Error()})

				return
			}
			ansTa, err := ta.usecase.GetTransByID(par.TransactionID)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			WriteJSONResponse(w, http.StatusOK, &ansTa)

			return
		}

		// Getting terminal ID from request - type: []int.
		if value, ok := values[keyTmID]; ok {
			{
				for _, v := range value {
					id, err := strconv.Atoi(v)
					if err != nil {
						WriteJSONResponse(w, http.StatusBadRequest, Response{
							Message: MSgBadURL,
							Details: err.Error()})

						return
					}
					par.TerminalID = append(par.TerminalID, id)
				}

				ansTa, err := ta.usecase.GetTransByTermID(par.TerminalID)
				if err != nil {
					WriteJSONResponse(w, http.StatusInternalServerError, Response{
						Message: MsgInternalSeverErr,
						Details: err.Error()})

					return
				}
				if ansTa == nil {
					WriteJSONResponse(w, http.StatusOK, Response{
						Message: MsgNotFound,
						Details: fmt.Sprint("No data found with input: ", par.TerminalID),
					})

					return
				}
				WriteJSONResponse(w, http.StatusOK, ansTa)

				return
			}

			// Getting parameter status from request - accepted|declined.
			if value, ok := values[keyStatus]; ok {
				{
					par.Status = value[0]
					if par.Status != statusValAccepted && par.Status != statusValDeclined {
						WriteJSONResponse(w, http.StatusBadRequest, Response{
							Message: MSgBadURL,
							Details: err.Error()})

						return
					}
				}
			}
			// Getting parameter payment type from request - cash|card.
			if value, ok := values[keyPayType]; ok {
				{
					par.Status = value[0]
					if par.PaymentType != payValCash && par.PaymentType != payValCard {
						WriteJSONResponse(w, http.StatusBadRequest, Response{
							Message: MSgBadURL,
							Details: err.Error()})

						return
					}
				}
			}
			// Getting parameter Date from request - fromXXXX-XX-XXtoXXXX-XX-XX.
			if value, ok := values[keyDatePost]; ok {
				{
					par.DatePost, err = parser.ParseDateFromString(value[0])
					if err != nil {
						WriteJSONResponse(w, http.StatusBadRequest, Response{
							Message: MSgBadURL,
							Details: err.Error()})

						return
					}
				}
			}
			// Getting parameter payment narrative from request - string.
			if value, ok := values[keyPayNar]; ok {

				{
					par.PaymentNarrative = value[0]

				}

			}
		}
	})
}

func (ta TAHandler) GetTransactionsCSV() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
