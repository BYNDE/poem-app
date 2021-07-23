package poem

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

// Run - starting server
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 + time.Second,
		WriteTimeout:   10 + time.Second,
	}

	logrus.Info("Poem-APP is starting!")

	// ! s.httpServer.ListenAndServeTLS() https/tls пратокол, в будущем добавить. Нужен ключ для принятия данных...
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Poem-APP is stoping!")
	return s.httpServer.Shutdown(ctx)
}
