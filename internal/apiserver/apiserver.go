package apiserver

import (
	"net/http"
)

func Start(c *Config) error {
	s := newServer(c)

	s.logger.Infof("server starting on %s", c.BindAddr)
	return http.ListenAndServe(c.BindAddr, s)
}
