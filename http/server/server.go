package server

import (
	"context"
	"net/http"
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
	notify          chan error
	shutdownTimeout time.Duration
}

func ListenAndServe(opts *Options, handler http.Handler) *Server {
	s := &http.Server{
		Handler:      handler,
		Addr:         opts.Listen,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
	}
	server := &Server{
		server:          s,
		notify:          make(chan error, 1),
		shutdownTimeout: opts.ShutdownTimeout,
	}

	server.start()

	return server
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
