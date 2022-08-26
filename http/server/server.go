package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Options struct {
	Listen          string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type Server struct {
	server          *http.Server
	notifyChan      chan error
	shutdownTimeout time.Duration
}

func Start(serve http.Handler, opts *Options) {
	s := &http.Server{
		Handler:      serve,
		Addr:         opts.Listen,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
	}
	server := &Server{
		server:          s,
		notifyChan:      make(chan error, 1),
		shutdownTimeout: opts.ShutdownTimeout,
	}

	server.start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error

	select {
	case s := <-interrupt:
		log.Print("app - Run - signal: " + s.String())
	case err = <-server.notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.notify: %w", err))
	}

	if err = server.shutdown(); err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.shutdown: %w", err))
	}
}

func (s *Server) start() {
	go func() {
		s.notifyChan <- s.server.ListenAndServe()
		close(s.notifyChan)
	}()
}

func (s *Server) notify() <-chan error {
	return s.notifyChan
}

func (s *Server) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
