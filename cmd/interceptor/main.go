package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/jmcshane/interceptor-sdk/pkg/server"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/signals"
)

const (
	// Port is the port that the port that interceptor service listens on
	Port         = 8082
	readTimeout  = 5 * time.Second
	writeTimeout = 20 * time.Second
	idleTimeout  = 60 * time.Second
)

func main() {
	// set up signals so we handle the first shutdown signal gracefully
	ctx := signals.NewContext()

	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		log.Fatalf("Failed to get the Kubernetes client set: %v", err)
	}

	zap, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %s", err)
	}
	logger := zap.Sugar()
	ctx = logging.WithLogger(ctx, logger)
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Fatalf("failed to sync the logger: %s", err)
		}
	}()

	service, err := server.NewWithInterceptors(kubeClient, logger)
	if err != nil {
		log.Fatalf("failed to initialize core interceptors: %s", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", service)

	// For handling Readiness Probe
	mux.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "ok")
	})

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      mux,
	}

	logger.Infof("Listen and serve on port %d", Port)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("failed to start interceptors service: %v", err)
	}
}
