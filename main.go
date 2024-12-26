package main

import (
	"BuildingMicroservicesWithGo/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	log.Println("\033[34;2mThe server is running!\033[0m")

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
<<<<<<< HEAD
	productHandler := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", productHandler)
=======
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/h", helloHandler)
	sm.Handle("/g", goodbyeHandler)
>>>>>>> d37ec1b5d1aaec1e4a89612c377983a039690594

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() { // Если без горутины - залочимся
		err := server.ListenAndServe()
<<<<<<< HEAD
		time.Sleep(time.Second)
=======
>>>>>>> d37ec1b5d1aaec1e4a89612c377983a039690594
		if err != nil {
			l.Fatalf("Error with server starting: %#v", err)
		}
	}()

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
<<<<<<< HEAD
	l.Printf("\033[31mThe interrupt command received: %v\033[0m\n", sig)
=======
	l.Printf("\033[31mThe interrupt command received\033[0m: %v\n", sig)
>>>>>>> d37ec1b5d1aaec1e4a89612c377983a039690594

	// Добавим закрытие сервака (аналог GS)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
