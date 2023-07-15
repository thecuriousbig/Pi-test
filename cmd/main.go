package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pi/cmd/httpserver"
	"pi/config"
	"pi/infrastructures"
	"pi/internal/core/services/usersvc"
	"pi/internal/handlers/userhdl"
	"pi/internal/repositories"
	"syscall"
	"time"
)

// @title                     pi test API
// @version                   1.0
// @description               This is a pi test API server.
// @termsOfService            http://swagger.io/terms/
// @contact.name              Tanatorn Nateesanpraser
// @contact.email             tanatorn.nateesanprasert@gmail.com
// @license.name              Apache 2.0
// @license.url               http://www.apache.org/licenses/LICENSE-2.0.html
// @host                      localhost:8080
// @BasePath                  /api/v1
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func init() {
	config.New()
}

func main() {
	// infrastructures
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		config.Get().Database.Host,
		config.Get().Database.Port,
		config.Get().Database.Username,
		config.Get().Database.Password,
		config.Get().Database.Database,
	)
	db := infrastructures.NewPostgres(dsn)
	rc := infrastructures.NewRedis()

	// repositories
	ur := repositories.NewUserRepository(db)
	ucr := repositories.NewUserCacheRepository(rc)
	// services
	us := usersvc.New(ur, ucr)
	// handlers
	uh := userhdl.New(us)

	e := httpserver.NewHTTPServer(uh)

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", config.Get().Endpoint.Port)); err != nil {
			e.Logger.Info("shutting down the server")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		defer e.Shutdown(ctx)
	}()

	// Graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
