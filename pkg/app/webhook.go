package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/juddbaguio/go-paymongo-webhook/pkg/controllers"
)

type PaymongoWebhookServer struct {
	mux    *mux.Router
	server http.Server
}

func InitWebhookServer() *PaymongoWebhookServer {
	mux := mux.NewRouter()
	SetupRoutes(mux)

	server := http.Server{
		Addr:         ":3000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return &PaymongoWebhookServer{
		mux:    mux,
		server: server,
	}
}

func (p *PaymongoWebhookServer) StartApp() error {
	serverError := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Webhook Ready on Port 3000")
		serverError <- p.server.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		return fmt.Errorf("server error: %v", err)
	case shutdownErr := <-shutdown:
		log.Println("graceful shutdown is starting")
		defer log.Printf("server successfully shtudown: %v\n", shutdownErr.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := p.server.Shutdown(ctx); err != nil {
			p.server.Close()
			return fmt.Errorf("could not stop the server gracefully: %v", err.Error())
		}

		<-ctx.Done()
	}

	return nil
}

func SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/webhook", controllers.ListenPaymongoWebhook).Methods("POST")
}
