package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adeniyistephen/zeinamfi/app/business"
	"github.com/adeniyistephen/zeinamfi/app/routes"

	"github.com/gofiber/fiber/v2"
)

func newRouter() *fiber.App {
	// new fiber instance
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the ZeinaMFI-API :)")
	})

	return app
}

func main() {
	srv := newRouter()

	// initialize database
	database := business.DB{}
	database.ConnectDb()

	route := routes.Routes{
		Database: database,
	}

	// handle routes
	srv.Get("/account/:accntnum", route.GetBalance)
	srv.Post("/add_account", route.CreateAccount)
	srv.Patch("/deposit/:accntnum", route.CreateDeposit)
	srv.Get("/ledger/:accntnum", route.GetLedger)
	srv.Patch("/withdraw/:accntnum", route.CreateWithdraw)

	// handle graceful shutdown and server listening
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Println("service interrupt received")

		log.Println("http server shutting down")
		time.Sleep(5 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.ShutdownWithContext(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}

		log.Println("shutdown complete")

		close(idleConnsClosed)

	}()

	// server listening on port 10101
	log.Printf("Starting server on port 10101")
	if err := srv.Listen(":10101"); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("fatal http server failed to start: %v", err)
		}
	}

	<-idleConnsClosed
	log.Println("Service Stop")
}
