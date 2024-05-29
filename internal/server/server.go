package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server

	AdditionDuration       time.Duration
	SubstractionDuration   time.Duration
	MultiplicationDuration time.Duration
	DivisionDuration       time.Duration
}

func New(addition, substraction, multiplication, division time.Duration) *Server {
	mux := http.NewServeMux()

	// mux handlerfuncs here

	return &Server{
		Server: &http.Server{
			Handler: mux,
			Addr:    ":8080",
		},

		AdditionDuration:       addition,
		SubstractionDuration:   substraction,
		MultiplicationDuration: multiplication,
		DivisionDuration:       division,
	}
}

func (s *Server) Run() error {
	err := s.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server crushed with err: %v", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.Server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("server shutdowned with error: %w", err)
	}

	return nil
}
