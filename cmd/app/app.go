package main

import (
	"context"
	"fmt"
	"github.com/Pxe2k/kaspi-task/internal/delivery/person"
	"github.com/Pxe2k/kaspi-task/internal/repository"
	person2 "github.com/Pxe2k/kaspi-task/internal/usecase/person"
	"github.com/Pxe2k/kaspi-task/pkg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	pg, err := pkg.NewPGXPool(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr: fmt.Sprint(":", os.Getenv("PORT")),
		Handler: person.New(
			person2.New(
				repository.New(pg),
			),
		),
	}

	log.Println("server started...")

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
