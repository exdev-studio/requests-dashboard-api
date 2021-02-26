package apiserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type server struct {
	logger *logrus.Logger
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

	return &server{
		logger,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Fatal("ServeHTTP() is not implemented")
}
