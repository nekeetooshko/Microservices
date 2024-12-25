package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	log.Println("The server is running!")

	// Хэндлеры (регистрируют функции в DefaultServeMux)
	http.HandleFunc("/h", helloHandler)
	http.HandleFunc("/g", goodbyeHandler)

	http.ListenAndServe(":9090", nil) // Используем DefaultServeMux
}

// Приветственный хендлер
func helloHandler(rw http.ResponseWriter, req *http.Request) {
	log.SetOutput(rw) // Переустанавливаю базовый вывод инфы
	log.Println("Hello from client!")

	// Считывам данные из запроса (req.Body - интерфейс io.ReadCloser => нужно что-то, умеющее читать из него)
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Error while reading the body: %e\n", http.StatusBadRequest)
		return // http.Error не прерывает сервачок
	}
	log.Printf("Data: %s\n", string(data)) // Можно было бы через rw.Write, но там нельзя писать в кавычках

}

// Прощальный хендлер
func goodbyeHandler(rw http.ResponseWriter, req *http.Request) {
	log.SetOutput(rw)
	log.Println("Goodbye from client!")
}
