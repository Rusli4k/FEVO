package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rusli4k/fevo/cfg"
)

type Server struct{ *http.Server }

type Handlers struct {
	TAHandler TAHandler
}

func NewServer(opt cfg.Options, handlers Handlers) *Server {
	router := new(mux.Router)

	attachEndpoints(router, handlers)

	handler := router

	return &Server{
		Server: &http.Server{
			Addr:         opt.Server.Host + ":" + opt.Server.Port,
			Handler:      http.TimeoutHandler(handler, handlerTimeoutSeconds*time.Second, MsgTimeOut),
			ReadTimeout:  readTimeoutSeconds * time.Second,
			WriteTimeout: writeTimeoutSeconds * time.Second,
		},
	}
}

// Run will run our server.
func (srv *Server) Run() error {
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error loading the server: %w", err)
	}

	return nil
}

func attachEndpoints(router *mux.Router, handlers Handlers) {
	router.Path("/transactions").Methods(http.MethodPost).Handler(handlers.TAHandler.UploadTransactions())
	router.Path("/transactions").Methods(http.MethodGet).
		Queries(transaction_id).
		Queries(terminal_id).
		Queries(status).
		Queries(payment_type).
		Queries(date_post).
		Queries(payment_narrative).
		Handler(handlers.TAHandler.GetTransactions())

	// var allowedSortArgs = strings.Join([]string{firstName, lastName, createdAt}, "|")
	// router.Path("/users").Methods(http.MethodGet).
	// 	Queries(offset, "{"+offset+":[0-9]+}").
	// 	Queries(limit, "{"+limit+":[0-9]+}").
	// 	Queries(sort, "{"+sort+":(?:"+allowedSortArgs+")(?:[,]{1}(?:"+allowedSortArgs+")*)*}").
	// 	Handler(handlers.UserHandler.GetUsers())
}
