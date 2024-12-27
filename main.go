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

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter() // Можно и просто ручками указать "GET"
	getRouter.HandleFunc("/", productHandler.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProducts)
	/* Возьмет id-шник, значения которого - цифры от 0 до 9 + значит, что их может быть много.
	А вообще, под капотом мы создаем переменную - id и далее этот id-шник пойдет в mux.Vars, и на UpdateProduct'e мы его извлечем
	*/

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
