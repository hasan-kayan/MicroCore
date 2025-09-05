package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/adapters/repository/mongo"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/config"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/services"
	httpserver "github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/http"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/http/handlers"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.New(cfg.App.Env, cfg.Logging.Level)
	if err != nil {
		log.Fatal(err)
	}
	defer logg.Sync()

	cb, err := mongo.Connect(cfg.Mongo.URI, cfg.Mongo.DB, cfg.Mongo.HealthCollection)
	if err != nil {
		logg.Fatal("mongo connect failed", zapErr(err))
	}
	healthRepo := mongo.NewHealthRepo(cb)
	healthSvc := services.NewHealthService(healthRepo)
	healthH := &handlers.HealthHandler{Svc: healthSvc}
	router := httpserver.NewRouter(healthH)

	srv := &http.Server{
		Addr:         cfg.App.HTTPAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		logg.Info("starting http server",
			zapField("addr", cfg.App.HTTPAddr),
			zapField("version", cfg.App.Version),
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logg.Fatal("http server error", zapErr(err))
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := time.WithTimeout(context.Background(), cfg.App.ShutdownTimeout)
	defer cancel()
	_ = srv.Shutdown(ctx)
	_ = cb.Client.Disconnect(ctx)
	logg.Info("shutdown complete")
}

// küçük yardımcılar (zap import’u yazmadan sade kullanmak için)
func zapErr(err error) interface{}         { return err }
func zapField(k string, v any) interface{} { return map[string]any{k: v} }
