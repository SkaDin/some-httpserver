package app

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"net/http"
	"some-httpserver/internal/cfg"
	"time"
)

type Server struct {
	config cfg.Config
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Config, ctx context.Context) *Server {
	server := &Server{}
	server.ctx = ctx
	server.config = config
	return server
}

func (s *Server) Serve() {
	log.Printf("Starting server")
	var err error
	s.db, err = pgxpool.New(s.ctx, s.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}

	carStorage := db3.NewCarStorage(s.db)
	userStorage := db3.NewUserStorage(s.db)

	carsProcessor := processors.NewCarsProcessor(carStorage)
	userProcessor := processor.NewUserProcessor(userStorage)

	carsHandler := handlers.NewCarHandler(carsProcessor)
	userHandler := handlers.NewUserHandler(userProcessor)

	routes := api.CreateRoutes(userHandler, carsHandler)
	rotues.Use(middleware.RequestLog)

	s.srv = &http.Server{
		Addr:    ":" + s.config.DbPort,
		Handler: routes,
	}
	log.Println("Server started")

	err = s.srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (s *Server) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), time.Second*5)

	s.db.Close()
	defer func() {
		cancel()
	}()
	var err error
	if err = s.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shotdown: %v", err)
	}
	log.Println("server exited properly")

	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}
}
