package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/sync/errgroup"

	"github.com/thirteenths/test-bmstu23/cmd/configer/appconfig"
	"github.com/thirteenths/test-bmstu23/internal/app"
	"github.com/thirteenths/test-bmstu23/internal/domain/storage"
	"github.com/thirteenths/test-bmstu23/internal/infra/postgres"
	"github.com/thirteenths/test-bmstu23/pkg/handler"

	internalhttp "github.com/thirteenths/test-bmstu23/internal/ports/http"
)

//go:generate go run ../configer --apps api --envs local,prod,dev

func main() {
	cfg := appconfig.MustParseAppConfig[appconfig.APIConfig]()

	logger := logrus.New()

	lvl, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(lvl)

	if _, err := maxprocs.Set(maxprocs.Logger(logger.Infof)); err != nil {
		logger.WithError(err).Errorf("can't set maxprocs")
	}

	eg, ctx := errgroup.WithContext(context.Background())

	// HTTP servers.
	jsonRenderer := handler.NewJSONRenderer()

	servers := make([]*http.Server, 0)
	router := chi.NewRouter()

	// Storage
	pg, err := postgres.NewPostgres(
		"postgres://postgres:7dgvJVDJvh254aqOpfd@docker:5432/postgres?sslmode=disable",
	)
	if err != nil {
		logger.WithError(err).Errorf("can`t connect to postgres")
	}
	stg := storage.NewStorage(*pg)

	// services
	apiService := app.NewAPI(logger)
	eventService := app.NewEventService(logger, stg)

	// Main API router.
	mainGroupHandler := handler.NewGroupHandler("/",
		internalhttp.NewAPIHandler(jsonRenderer, apiService),
		internalhttp.NewEventHandler(jsonRenderer, *eventService),
	)

	mainHandler := handler.New(handler.MakePublicRoutes(
		router,
		handler.RoutesCfg{
			BasePath: cfg.Servers.Public.BasePath,
		},
		mainGroupHandler))

	servers = append(servers, &http.Server{
		Addr:     cfg.Servers.Public.ListenAddr,
		Handler:  mainHandler,
		ErrorLog: log.New(logger.Out, "api", 0),
	})

	logger.Debugf("Listing actual routes:\n")

	_ = chi.Walk(
		router,
		func(
			method string,
			route string,
			handler http.Handler,
			middlewares ...func(http.Handler) http.Handler,
		) error {
			logger.Debugf("[%s]: /rush-stand-up-club/%s%s\n", method, appconfig.APIAppName, route)
			return nil
		})

	for i := range servers {
		srv := servers[i]
		go func() {
			logger.Infof("starting server, listening on %s", srv.Addr)

			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				logger.WithError(err).Errorf("server can't listen and serve requests")
			}
		}()
	}

	logger.Infof("app started")

	sigQuit := make(chan os.Signal, 1)

	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("captured signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	if err = eg.Wait(); err != nil {
		logger.WithError(err).Infof("gracefully shutting down the server")
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	for _, srv := range servers {
		if err := srv.Shutdown(timeoutCtx); err != nil {
			logger.WithError(err).Fatalf("can't close server listening on '%s'", srv.Addr)
		}
	}

	logger.Info("app has been terminated")
}
