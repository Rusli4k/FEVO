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
	router.Path("/transactions").Methods(http.MethodGet).Handler(handlers.TAHandler.GetTransactions())
	//	router.Path("/transactions/CSV").Methods(http.MethodGet).Handler(handlers.TAHandler.GetTransactionsCSV())
}
