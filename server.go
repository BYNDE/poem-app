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
	httpServer  *http.Server
	httpsServer *http.Server
}

// Run - starting server
func (s *Server) Run(port string, portTLS string, enableTLS bool, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ErrorLog:       log.New(logrus.StandardLogger().Writer(), "", 0),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.httpsServer = &http.Server{
		Addr:           ":" + portTLS,
		Handler:        handler,
		ErrorLog:       log.New(logrus.StandardLogger().Writer(), "", 0),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if enableTLS {
		if err := httpscerts.Check("cert.pem", "key.pem"); err != nil {
			if err := httpscerts.Generate("cert.pem", "key.pem", "localhost"); err != nil {
				return err
			}
		}
		go s.httpsServer.ListenAndServeTLS("cert.pem", "key.pem")
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) [2]error {
	var errors [2]error

	errors[0] = s.httpServer.Shutdown(ctx)
	errors[1] = s.httpsServer.Shutdown(ctx)

	return errors
}
