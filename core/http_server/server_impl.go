package httpserver

import (
	"fmt"
	"net/http"
	"time"
	"world_generator_processor_service/config"
	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	_ "world_generator_processor_service/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpServerTools "github.com/olehmushka/golang-toolkit/http_server_tools"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
)

type server struct {
	handlers Handlers

	address                string
	basePath               string
	readTimeoutInSec       int
	readHeaderTimeoutInSec int
	writeTimeoutInSec      int
	idleTimeoutInSec       int
}

func New(cfg *config.Config, handlers Handlers) Server {
	return &server{
		handlers: handlers,

		address:                cfg.HTTPServer.Address,
		basePath:               cfg.HTTPServer.BasePath,
		readTimeoutInSec:       cfg.HTTPServer.ReadTimeoutInSec,
		readHeaderTimeoutInSec: cfg.HTTPServer.ReadHeaderTimeoutInSec,
		writeTimeoutInSec:      cfg.HTTPServer.WriteTimeoutInSec,
		idleTimeoutInSec:       cfg.HTTPServer.IdleTimeoutInSec,
	}
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(Register),
)

func Register(s Server) {
	go s.Register()
}

func (s *server) Register() {
	router := chi.NewRouter()

	router.Use(middleware.Logger, traceIDTools.SetTraceIDMiddleware)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", traceIDTools.TraceIDHeader},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           1000,
	}))
	// swagger UI
	router.Get(s.getPath(DocRouteName), httpSwagger.Handler(
		httpSwagger.URL(s.getPath(DocJsonRouteName)),
	))

	router.Get(s.getPath(InfoRouteName), httpServerTools.NewHandlesChain(s.handlers.GetInfo))
	router.Get(s.getPath(HealthCheckRouteName), httpServerTools.NewHandlesChain(s.handlers.GetHealthCheck))

	srv := &http.Server{
		Addr:              s.address,
		Handler:           router,
		ReadTimeout:       time.Duration(s.readTimeoutInSec),
		ReadHeaderTimeout: time.Duration(s.readHeaderTimeoutInSec),
		WriteTimeout:      time.Duration(s.writeTimeoutInSec),
		IdleTimeout:       time.Duration(s.idleTimeoutInSec),
	}

	logrus.WithFields(logrus.Fields{"http_server_address": s.address}).
		Info("starting http server...")
	if err := srv.ListenAndServe(); err != nil {
		logrus.
			Error("can not start http server", err.Error())
	}
}

func (s *server) getPath(routeName string) string {
	return fmt.Sprintf("/%s/%s", s.basePath, routeName)
}
