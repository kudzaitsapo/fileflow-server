package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kudzaitsapo/fileflow-server/internal/cache"
	"github.com/kudzaitsapo/fileflow-server/internal/config"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)


type Application struct {
	mux *http.ServeMux
	middleware []func(http.Handler) http.Handler
	AppConfig config.ApplicationConfig
	store *store.Storage
	cache *cache.Storage
}


func CreateApplication(config config.ApplicationConfig) *Application {
	return &Application{
		mux: http.NewServeMux(),
		middleware: []func(http.Handler) http.Handler{},
		AppConfig: config,
	}
}

func (app *Application) SetStore(store *store.Storage) {
	app.store = store
}

func (app *Application) SetCache(cache *cache.Storage) {
	app.cache = cache
}

func (app *Application) Use(middleware func(http.Handler) http.Handler) {
	app.middleware = append(app.middleware, middleware)
}

func (app *Application) Handle(pattern string, handler http.Handler) {
	finalHandler := handler
	for i := len(app.middleware) - 1; i >= 0; i-- {
		finalHandler = app.middleware[i](finalHandler)
	}
	app.mux.Handle(pattern, finalHandler)
}


func (app *Application) ListenAndServe() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.AppConfig.Config.Port),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
		Handler:      app.mux,
	}
	return srv.ListenAndServe()
}