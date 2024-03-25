package server

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/fawzi2702/FizzBuzz/api/docs"
	"github.com/fawzi2702/FizzBuzz/api/helpers/environment"
	"github.com/fawzi2702/FizzBuzz/api/routers"
	"github.com/gin-gonic/gin"
)

type Server = http.Server

func SetupServer() (*Server, error) {
	var err error

	// Setup Gin
	r := gin.Default()

	// Set Gin mode
	envMode, err := environment.Get("MODE")
	if err != nil {
		return nil, err
	}

	if envMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Setup groups
	routers.SetupFizzbuzzRouter(r)
	routers.SetupStatsRouter(r)
	routers.SetupDocsRouter(r) // Docs

	// Start server
	apiPort, err := environment.Get("API_PORT")
	if err != nil {
		return nil, fmt.Errorf("API_PORT not found in environment: %w", err)
	}

	srv := &Server{
		Addr:    ":" + apiPort,
		Handler: r,
	}

	return srv, nil
}

func StartServer(srv *Server) error {
	if srv == nil {
		return fmt.Errorf("server is not set")
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	} else {
		fmt.Printf("Server started on %s\n", srv.Addr)

		return nil
	}
}

func ShutdownServer(srv *Server) {
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("Server shutdown with error: %s\n", err)
	} else {
		fmt.Println("Server shutdown")
	}
}

func RestartServer(srv *Server) error {
	if err := srv.Shutdown(context.Background()); err != nil {
		return err
	}

	return StartServer(srv)
}
