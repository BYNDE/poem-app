package poem

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"
)

type Server struct {
	httpServer  *http.Server
	httpsServer *http.Server
}

// ! https://eli.thegreenplace.net/2021/go-https-servers-with-tls/ - MTLS для безопасного сойденения

// Run - starting server
func (s *Server) Run(port string, portTLS string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.httpsServer = &http.Server{
		Addr:           ":" + portTLS,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
		},
	}

	if err := Check("cert.pem", "key.pem"); err != nil {
		if err := Generate("ByndeCorp", "localhost", "cert.pem", "key.pem"); err != nil {
			return err
		}
	}

	go s.httpsServer.ListenAndServeTLS("cert.pem", "key.pem")

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) [2]error {
	var errors [2]error

	errors[0] = s.httpServer.Shutdown(ctx)
	errors[1] = s.httpsServer.Shutdown(ctx)

	return errors
}
