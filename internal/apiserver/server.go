package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/exdev-studio/requests-dashboard-api/internal/store"
	"github.com/exdev-studio/requests-dashboard-api/internal/store/memstore"
)

type server struct {
	logger *logrus.Logger
	store  store.Store
	router *mux.Router
}

func newServer(c *Config) *server {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logLevel, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(logLevel)
	logger.Debugf("log level set to %s", c.LogLevel)

	s := memstore.New()

	srv := &server{
		logger: logger,
		store:  s,
		router: mux.NewRouter(),
	}

	srv.configureRouter()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/requests", s.handleRequestsList()).Methods(http.MethodGet)
}

func (s *server) handleRequestsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requests, err := s.store.Request().List()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			s.logger.Error(err)
		}

		s.respond(w, http.StatusOK, requests)
	}
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.logger.Error(err)
		}
	}
}

func (s *server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, map[string]string{
		"error": err.Error(),
	})
}
