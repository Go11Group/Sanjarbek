package main

import (
	"log"
	"module/handler"
	"module/storage/postgres"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)
	problemRepo := postgres.NewProblemRepo(db)
	solvedProblemRepo := postgres.NewSolvedProblemRepo(db)

	h := handler.Handler{
		User:          userRepo,
		Problem:       problemRepo,
		SolvedProblem: solvedProblemRepo,
	}

	r := handler.NewHandler(h)


	r.SetTrustedProxies([]string{"192.168.1.2", "192.168.1.3"})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
