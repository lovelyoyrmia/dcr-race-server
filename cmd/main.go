package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/lovelyoyrmia/dcr_race/domain"
	"github.com/lovelyoyrmia/dcr_race/domain/repositories"
	"github.com/lovelyoyrmia/dcr_race/domain/services"
	"github.com/lovelyoyrmia/dcr_race/domain/websockets"
	"github.com/lovelyoyrmia/dcr_race/internal/db"
	"github.com/lovelyoyrmia/dcr_race/pkg/config"
)

var interruptSignal = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	// Load configuration
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("error occurred : %v", err)
	}

	// create a new message dispatcher for websocket connections
	dispatcher := websockets.NewDispatcher()
	go dispatcher.Run()

	// Add interrupt signal
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignal...)
	defer stop()

	// Initialize Database
	database := db.NewDatabase(ctx, c)
	database.NewMigrations()

	store := db.NewStore(database.DB)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runHTTPServer(waitGroup, ctx, c, store, dispatcher)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runHTTPServer(
	waitGroup *errgroup.Group,
	ctx context.Context,
	c config.Config,
	db db.Querier,
	dispatcher *websockets.Dispatcher,
) {
	router := echo.New()
	router.Use(middleware.Logger())

	userRepo := repositories.NewUserRepositories(db)

	userServices := services.NewUserLocationServices(userRepo, c, dispatcher)

	// create websocket handlers
	wsHandler := websockets.NewHandlers(dispatcher)

	userRoutes := domain.NewUserLocationRoutes(userServices)

	group := router.Group("/v1")
	userRoutes.SetupRoutes(group)
	websockets.SetRoutes(group, wsHandler)

	s := &http.Server{
		Addr:         c.Address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info().Msgf("Starting HTTP server on port: %s", s.Addr)
	// Create go routines to serve http
	waitGroup.Go(func() error {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal().Msgf("cannot start HTTP server: %v", err)
			return err
		}
		return nil
	})

	// Waiting server to gracefully shutdown
	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown HTTP server")

		err := s.Shutdown(context.Background())
		if err != nil {
			log.Error().Msg("failed to shutdown http server")
			return err
		}
		log.Info().Msg("HTTP server is stopped")
		return nil
	})
}
