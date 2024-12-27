package main

import (
	"BuildingMicroservicesWithGo/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("\033[34;3mThe server is running!\033[0m")

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	productHandler := handlers.NewProduct(l)

	sm := mux.NewRouter() // Создаем роутер от гориллы (основной/корневой)

	// Создаем отдельный саброутер для обработки GET - запросов. Methods вернет rout, созданный специально для
	// GET - запросов. В конце, через .SubRouter - конвертим это в роутер
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	l.Printf("\033[31;3mThe interrupt command received: %v\033[0m\n", sig)

	// Добавим закрытие сервака (аналог GS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
