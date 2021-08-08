package poem

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kabukky/httpscerts"
	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

// Run - starting server
func (s *Server) Run(port string, enableTLS bool, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ErrorLog:       log.New(logrus.StandardLogger().Writer(), "", 0),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if enableTLS {
		if err := httpscerts.Check("cert.pem", "key.pem"); err != nil {
			if err := httpscerts.Generate("cert.pem", "key.pem", "localhost:8080"); err != nil {
				return err
			}
		}
		return s.httpServer.ListenAndServeTLS("cert.pem", "key.pem")
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
