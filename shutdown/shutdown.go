package shutdown

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	threshHoldForceShutdown = 10
)

var (
	sigHandler *sigtermHandler
	logger     *zap.Logger
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Println("shutdown handler: cannot init zap logger, use noop instead")
		logger = zap.NewNop()
	}
	sigHandler = initSigtermHandler()
}

type sigtermHandler struct {
	functions  []func() error
	sigCount   int
	sigChannel chan os.Signal
	timeout    time.Duration
	mu         sync.Mutex
	done       chan struct{}
	cleanOnce  sync.Once
}

// SetTimeout set timeout for running deferred functions. It should be set before all of RegisterDeferFunc
func (s *sigtermHandler) SetTimeout(duration time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.timeout = duration
}

// RegisterFunc registers functions for calling when a sigterm signal is received
func (s *sigtermHandler) RegisterFunc(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.functions = append(s.functions, func() error {
		f()
		return nil
	})
}

// RegisterErrorFunc registers functions for calling when sigterm signal is received
func (s *sigtermHandler) RegisterErrorFunc(f func() error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.functions = append(s.functions, f)
}

// RegisterErrorFuncContext registers functions for calling when sigterm signal is received
func (s *sigtermHandler) RegisterErrorFuncContext(ctx context.Context, f func(ctx context.Context) error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.functions = append(s.functions, func() error {
		return f(ctx)
	})
}

// Wait waits until all registered functions are done
func (s *sigtermHandler) Wait() {
	// Ensure that clean is called even if panic
	go s.clean()
	// Wait for cleaning
	logger.Info("sigterm Handler: waiting for gracefully finishing current works before shutdown...")
	<-s.done
	logger.Info("sigterm Handler: done!")
}

func (s *sigtermHandler) handleTermination() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sigCount++
	switch s.sigCount {
	case 1:
		go s.clean()
		if s.timeout > 0 {
			go func() {
				time.Sleep(s.timeout)
				logger.Info("sigterm Handler: timeout! force application to exit")
				os.Exit(1)
			}()
		}
	case threshHoldForceShutdown:
		logger.Info("sigterm Handler: force application to exit...")
		//nolint:gocritic
		os.Exit(1)
	}
}

func (s *sigtermHandler) clean() {
	s.cleanOnce.Do(func() {
		defer close(s.done)
		logger.Info("sigterm Handler: start cleaning functions")
		for i := len(s.functions) - 1; i >= 0; i-- {
			if err := s.functions[i](); err != nil {
				logger.Error("sigterm Handler: cleaning function error", zap.String("error", err.Error()))
			}
		}
	})
}

// SigtermHandler returns sigterm handler for graceful shutdown
//
//nolint:golint
func SigtermHandler() *sigtermHandler {
	return sigHandler
}

func initSigtermHandler() *sigtermHandler {
	sigtermHandler := &sigtermHandler{
		functions:  make([]func() error, 0),
		sigCount:   0,
		sigChannel: make(chan os.Signal, 1),
		timeout:    -1,
		done:       make(chan struct{}),
	}
	signal.Notify(sigtermHandler.sigChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			s := <-sigtermHandler.sigChannel
			logger.Info(fmt.Sprint("sigterm handler: received signal: ", s))
			sigtermHandler.handleTermination()
		}
	}()
	return sigtermHandler
}
