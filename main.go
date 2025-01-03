package main

import (
	"BuildingMicroservicesWithGo/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("\033[34;3mThe server is running!\033[0m")

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	productHandler := handlers.NewProduct(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter() // Можно и просто ручками указать "GET"
	getRouter.HandleFunc("/", productHandler.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddleWareProductValidation)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProducts)
	putRouter.Use(productHandler.MiddleWareProductValidation)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() { // Если без горутины - залочимся
		err := server.ListenAndServe()
		time.Sleep(time.Second)
		if err != nil {
			l.Fatalf("Error with server starting: %#v", err)
		}
	}()

	gracefulShutdown(server)
}

// Graceful shutdown
func gracefulShutdown(server *http.Server) {

	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select { // Он тут по факту нахуй не нужен, но меня научили писать GS на нем

	case sig := <-signalChan: // Можно было бы без присваивания, но я вывожу сигнал, так что

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Error while server is shutting down: %#v\n", err)
		}

		// Здесь стоит escape-команда. Это моя шиза. Все хорошо.
		log.Printf("\033[31mConnection stop by signal: %v, %#v\n\033[0m", sig, sig)
	}
}
