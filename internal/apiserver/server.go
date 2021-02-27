package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/exdev-studio/requests-dashboard-api/internal/model"
	"github.com/exdev-studio/requests-dashboard-api/internal/store"
	"github.com/exdev-studio/requests-dashboard-api/internal/store/memstore"
)

const (
	ctxKeyRequestID ctxKey = iota
)

type ctxKey int8

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
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)

	requests := s.router.PathPrefix("/requests").Subrouter()
	requests.HandleFunc("", s.handleRequestsList()).Methods(http.MethodGet)
	requests.HandleFunc("/collect", s.handleRequestsCollect()).Methods(http.MethodPost)
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

func (s *server) handleRequestsCollect() http.HandlerFunc {
	type request struct {
		model.Request
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		request := &model.Request{
			Type:   req.Type,
			Fields: req.Fields,
		}

		if err := s.store.Request().Collect(request); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, request)
	}
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
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

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set(HeaderRequestID, id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Debugf("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Debugf(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}
