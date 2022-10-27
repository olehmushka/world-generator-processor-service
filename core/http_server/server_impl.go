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

	address           string
	basePath          string
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
}

func New(cfg *config.Config, handlers Handlers) Server {
	return &server{
		handlers: handlers,

		address:           cfg.HTTPServer.Address,
		basePath:          cfg.HTTPServer.BasePath,
		readTimeout:       time.Duration(cfg.HTTPServer.ReadTimeoutInSec) * time.Second,
		readHeaderTimeout: time.Duration(cfg.HTTPServer.ReadHeaderTimeoutInSec) * time.Second,
		writeTimeout:      time.Duration(cfg.HTTPServer.WriteTimeoutInSec) * time.Second,
		idleTimeout:       time.Duration(cfg.HTTPServer.IdleTimeoutInSec) * time.Second,
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
		ReadTimeout:       s.readTimeout,
		ReadHeaderTimeout: s.readHeaderTimeout,
		WriteTimeout:      s.writeTimeout,
		IdleTimeout:       s.idleTimeout,
	}

	log := logrus.WithFields(logrus.Fields{"http_server_address": s.address})
	log.Info("starting http server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Error("can not start http server", err.Error())
	}
}

func (s *server) getPath(routeName string) string {
	var basePath string
	if s.basePath != "" {
		basePath = "/" + s.basePath
	}
	return fmt.Sprintf("%s/%s", basePath, routeName)
}
