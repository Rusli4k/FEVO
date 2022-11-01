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
			if err := th.usecase.UploadTrans(v); err != nil {
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

// GetTransactions handles filters in request.
func (th TAHandler) GetTransactionsByFilter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var par entities.Filter
		var err error
		values := r.URL.Query()

		// Handling "transaction ID" from request.
		if value, ok := values[keyTaID]; ok {
			par.TransactionID, err = strconv.Atoi(value[0])
			if err != nil {
				WriteJSONResponse(w, http.StatusBadRequest, Response{
					Message: MSgBadURL,
					Details: err.Error()})

				return
			}
			ansTa, err := th.usecase.GetTransByID(par.TransactionID)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			if ansTa == nil {
				WriteJSONResponse(w, http.StatusOK, Response{
					Message: MsgNotFound,
					Details: fmt.Sprint(MsgNotFoundWithIn, value),
				})

				return
			}
			WriteJSONResponse(w, http.StatusOK, ansTa)

			return
		}

		// Handling "terminal ID" from request.
		if value, ok := values[keyTmID]; ok {
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

			ansTa, err := th.usecase.GetTransByTermID(par.TerminalID)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			if ansTa == nil {
				WriteJSONResponse(w, http.StatusOK, Response{
					Message: MsgNotFound,
					Details: fmt.Sprint(MsgNotFoundWithIn, par.TerminalID),
				})

				return
			}
			WriteJSONResponse(w, http.StatusOK, ansTa)

			return
		}

		// Handling parameter "status" from request - accepted|declined.
		if value, ok := values[keyStatus]; ok {
			par.Status = value[0]
			if par.Status != statusValAccepted && par.Status != statusValDeclined {
				WriteJSONResponse(w, http.StatusBadRequest, Response{
					Message: MSgBadURL,
					Details: fmt.Sprint(MsgInvalidInput, value)})

				return
			}
			ansTa, err := th.usecase.GetTransByStatus(par.Status)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			if ansTa == nil {
				WriteJSONResponse(w, http.StatusOK, Response{
					Message: MsgNotFound,
					Details: fmt.Sprint(MsgNotFoundWithIn, par.Status),
				})

				return
			}
			WriteJSONResponse(w, http.StatusOK, ansTa)

			return
		}

		// Handling parameter "payment type" from request - cash|card.
		if value, ok := values[keyPayType]; ok {
			par.PaymentType = value[0]
			if par.PaymentType != payValCard && par.PaymentType != payValCash {
				WriteJSONResponse(w, http.StatusBadRequest, Response{
					Message: MSgBadURL,
					Details: fmt.Sprint(MsgInvalidInput, par.PaymentType)})

				return
			}
			ansTa, err := th.usecase.GetTransByPayType(par.PaymentType)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			if ansTa == nil {
				WriteJSONResponse(w, http.StatusOK, Response{
					Message: MsgNotFound,
					Details: fmt.Sprint(MsgNotFoundWithIn, par.PaymentType),
				})

				return
			}
			WriteJSONResponse(w, http.StatusOK, ansTa)

			return
		}

		// Handling parameter "date post" from request.
		if value, ok := values[keyDatePost]; ok {
			par.DatePost, err = parser.ParseDateFromString(value[0])
			if err != nil {
				WriteJSONResponse(w, http.StatusBadRequest, Response{
					Message: MSgBadURL,
					Details: fmt.Sprint(MsgInvalidInput, value)})

				return
			}

			ansTa, err := th.usecase.GetTransByDataPost(par.DatePost)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			if ansTa == nil {
				WriteJSONResponse(w, http.StatusOK, Response{
					Message: MsgNotFound,
					Details: fmt.Sprint(MsgNotFoundWithIn, value),
				})

				return
			}
			WriteJSONResponse(w, http.StatusOK, ansTa)

			return
		}

		// Handling parameter "payment narrative" from request.
		if value, ok := values[keyPayNar]; ok {
			par.PaymentNarrative = value[0]
			if len(par.PaymentNarrative) < payNarMinLen {
				WriteJSONResponse(w, http.StatusBadRequest, Response{
					Message: MSgBadURL,
					Details: fmt.Sprint(MsgInvalidInput, payNarMinLen)})

				return
			}
			ansTa, err := th.usecase.GetTransByPayNar(par.PaymentNarrative)
			if err != nil {
				WriteJSONResponse(w, http.StatusInternalServerError, Response{
					Message: MsgInternalSeverErr,
					Details: err.Error()})

				return
			}
			if ansTa == nil {
				WriteJSONResponse(w, http.StatusOK, Response{
					Message: MsgNotFound,
					Details: fmt.Sprint(MsgNotFoundWithIn, value),
				})

				return
			}
			WriteJSONResponse(w, http.StatusOK, ansTa)

			return
		}

		// Handling wrong parameter
		WriteJSONResponse(w, http.StatusBadRequest, Response{
			Message: MSgBadURL,
			Details: fmt.Sprint(MsgInvalidInput, values)})
	})
}
