package main

import (
	"BuildingMicroservicesWithGo/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	log.Println("\033[34;2mThe server is running!\033[0m")

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags) // 1-куда выводим, 2-префикс, 3-настройки флагов
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	sm := http.NewServeMux() // Новый роутер
	sm.Handle("/h", helloHandler)
	sm.Handle("/g", goodbyeHandler)

	// Перепишем наш сервер
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second, // макс. время, пока коннект с клиентом активен
		ReadTimeout:  1 * time.Second,   // макс. вр. ожидания запроса на чтение от клиента
		WriteTimeout: 1 * time.Second,   // макс. вр. ожидания запроса на запись от клиента
	}

	server.ListenAndServe()
}
