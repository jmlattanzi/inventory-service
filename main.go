package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"

	h "github.com/jmlattanzi/rex/services/inventory-service/api"
	"github.com/jmlattanzi/rex/services/inventory-service/inventory"
	mr "github.com/jmlattanzi/rex/services/inventory-service/repository/mongodb"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "development" {
		fmt.Println("loading development env")
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	repo := chooseRepo()
	service := inventory.NewInventoryService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/entry/category/{category}", handler.GetCategory)
	r.Put("/entry/{id}", handler.Update)
	r.Post("/entry", handler.Post)
	r.Get("/entry/{id}", handler.Get)
	r.Delete("/entry/{id}", handler.Delete)
	r.Get("/entry/all", handler.GetAll)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8080")
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() inventory.InventoryRepository {
	switch os.Getenv("URL_DB") {
	// case "redis":
	// 	redisURL := os.Getenv("REDIS_URL")
	// 	repo, err := rr.NewRedisRepository(redisURL)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return repo
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))

		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
