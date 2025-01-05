package main

import (
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/bxcodec/go-clean-arch/ilovepdf"

	"github.com/bxcodec/go-clean-arch/internal/rest"
	"github.com/bxcodec/go-clean-arch/internal/rest/middleware"
	_ "github.com/bxcodec/go-clean-arch/docs"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

// @title Swagger Example API
// @version 1.0

// @host            localhost:9090
// @BasePath        /v1
func main() {
	e := echo.New()
	e.Use(middleware.CORS)
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	svc := ilovepdf.NewService()
	rest.NewILovePdfHandlerHandler(e, svc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}
