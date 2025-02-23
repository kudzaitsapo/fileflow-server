package app

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/kudzaitsapo/fileflow-server/internal/auth"
	"github.com/kudzaitsapo/fileflow-server/internal/cache"
	"github.com/kudzaitsapo/fileflow-server/internal/config"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

var (
	currentApplication *Application
	mu sync.Mutex
)


type Application struct {
	mux *http.ServeMux
	middleware []func(http.Handler) http.Handler
	AppConfig config.ApplicationConfig
	Authenticator auth.Authenticator
	Store *store.Storage
	Cache *cache.Storage
}

func CreateApplication(config config.ApplicationConfig) *Application {
	return &Application{
		mux: http.NewServeMux(),
		middleware: []func(http.Handler) http.Handler{},
		AppConfig: config,
	}
}


func GetCurrentApplication() *Application {
	mu.Lock()
	defer mu.Unlock()
	return currentApplication
}

func SetCurrentApplication(app *Application) {
	mu.Lock()
	defer mu.Unlock()
	currentApplication = app
}

func (a *Application) SetAuthenticator(auth auth.Authenticator) {
	a.Authenticator = auth
}

func (a *Application) SetStore(store *store.Storage) {
	a.Store = store
}

func (a *Application) SetCache(cache *cache.Storage) {
	a.Cache = cache
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
