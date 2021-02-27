package apiserver

import (
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

	return &server{
		logger: logger,
		store:  s,
		router: mux.NewRouter(),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
