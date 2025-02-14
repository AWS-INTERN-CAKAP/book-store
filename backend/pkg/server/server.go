package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws-cakap-intern/book-store/pkg/response"
	"github.com/aws-cakap-intern/book-store/pkg/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

func NewServer(publicRoutes []*route.Route) *Server {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Hello, World!", nil))
	})

	v1 := e.Group("api")

	if len(publicRoutes) > 0 {
		for _, v := range publicRoutes {
			v1.Add(v.Method, v.Path, v.Handler)
		}
	}

	return &Server{e}
}

func (s *Server) Run(port string) {
	runServer(s, port)
	gracefulShutdown(s)
}

func runServer(srv *Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func gracefulShutdown(srv *Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal("Server Shutdown:", err)
		}
	}()
}
