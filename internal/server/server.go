package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ysrckr/deep_bucket_client/internal/server/routes"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      routes.NewRouter().RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func StartServer(ctx context.Context) error {

	server := NewServer()

	shutDownChan := make(chan error, 1)

	go func() {
		log.Printf("Server is running on http://localhost%s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			shutDownChan <- err
		}
		defer close(shutDownChan)
	}()

	select {
	case err := <-shutDownChan:
		return err
	case <-ctx.Done():
		log.Println("Shutting down server...")
		timeoutDuration := 5 * time.Second
		timeout, cancel := context.WithTimeout(context.Background(), timeoutDuration)
		defer cancel()
		if err := server.Shutdown(timeout); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
			return err
		}

		return nil
	}

}
