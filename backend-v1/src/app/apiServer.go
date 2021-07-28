package apiServer

import (
	"antriin/src/util/debug"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func httpServer(e *echo.Echo, address string) {

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Start HTTP Server
	e.HideBanner = true
	go func() {
		if err := e.Start(address); err != nil {
			fmt.Printf("[HTTP Server] - %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		debug.Error("MongoConnection", err.Error())
	}
}

func RunHttpServer(e *echo.Echo) {

	address := fmt.Sprintf("0.0.0.0:8030")

	// start server
	httpServer(e, address)

}
