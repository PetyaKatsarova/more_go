package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// CHANNEL: a type for communication accross go routines

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis, kuku, lala: %w", err)
	}
	fmt.Println("Starting server")

	defer func ()  {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis, kuku lala", err)
		}
	}()

	ch := make(chan error, 1) // chan has assoc type error: type sent across, buffer size 1: writer not
	// blocked till buffer is not full
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			// publishing val into the chan
			ch <- fmt.Errorf("failed to start server, KUKULALA: %w", err)
		}
		close(ch) // send signal for eveyone listening that its closed
	}()

	// select: like switch statement: lock on multiple chans at once
	select {
	case err = <-ch:
		return err // used buffer: if error exit
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
