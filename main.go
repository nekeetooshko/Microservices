package main

import (
	"log"
	"net/http"
)

func main() {

	http.ListenAndServe(":9090", nil) // Используем DefaultServeMux
	log.Println("The server is running!")

	// Хэндлеры (регистрируют функции в DefaultServeMux)
	http.HandleFunc("/h", helloHandler)
	http.HandleFunc("/g", goodbyeHandler)

}

// Приветственный хендлер
func helloHandler(rw http.ResponseWriter, req *http.Request) {
	log.SetOutput(rw) // Переустанавливаю базовый вывод инфы
	log.Println("Hello from client!")
}

// Прощальный хендлер
func goodbyeHandler(rw http.ResponseWriter, req *http.Request) {
	log.SetOutput(rw)
	log.Println("Goodbye from client!")
}
