package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Играюсь с цветами вывода (escape-команды). Чо кринж то сразу?
	log.Println("\033[34mThe server is running!\033[0m")

	// Хэндлеры (регистрируют функции в DefaultServeMux)
	http.HandleFunc("/h", helloHandler)
	http.HandleFunc("/g", goodbyeHandler)

	server := http.Server{Addr: ":9090"}

	/* Здесь выбрана такая реализция: сервак на горутине, GS на main'e, ведь если бы мы выбрали обратку
	(сервак на main, GS на горутине), то мы бы просто застопали main функцией ListenAndServe.
	Таким образом, он не сможет обрабатывать сигналы, пока сервер не завершит работу, что делает невозможным
	точное реагирование на Ctrl + C до тех пор, пока сервер не завершит свои операции.
	*/

	go func() {

		err := http.ListenAndServe(server.Addr, nil)
		if err != nil {
			log.Fatalf("Error with server running: %#v\n", err)
		}
	}()

	gracefulShutdown(&server)
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
