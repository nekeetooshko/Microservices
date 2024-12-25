package main

import (
	"BuildingMicroservicesWithGo/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("\033[34;2mThe server is running!\033[0m")

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags) // 1-куда выводим, 2-префикс, 3-настройки флагов

	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	sm := http.NewServeMux() // Новый роутер
	sm.Handle("/h", helloHandler)
	sm.Handle("/g", goodbyeHandler)

	if err := http.ListenAndServe(":9090", sm); err != nil {
		l.Fatalf("Error with server running: %#v\n", err)
	} // Используем наш ServeMux
}
