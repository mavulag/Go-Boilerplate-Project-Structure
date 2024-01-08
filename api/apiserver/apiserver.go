// apiserver/apiserver.go
package apiserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	//    gorilla/mux for our request router
	"github.com/gorilla/mux"
	"github.com/mavulag/trilabs/storage"

	//    logrus for our logging
	"github.com/sirupsen/logrus"

	"github.com/mavulag/trilabs/handlers"
)

var defaultStopTimeout = time.Second * 30

type APIServer struct {
	addr    string
	storage *storage.Storage
}

// NewAPIServer function returns an initialized server
// func NewAPIServer(addr string) (*APIServer, error) {
func NewAPIServer(addr string, storage *storage.Storage) (*APIServer, error) {
	if addr == "" {
		return nil, errors.New("addr cannot be blank")
	}

	return &APIServer{
		// addr: addr,
		addr:    addr,
		storage: storage,
	}, nil
}

// Start starts a server with a stop channel
func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
	}

	go func() {
		logrus.WithField("addr", srv.Addr).Info("starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStopTimeout)
	defer cancel()

	logrus.WithField("timeout", defaultStopTimeout).Info("stopping server")
	return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.DefaultRoute)
	return router
}
