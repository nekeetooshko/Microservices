package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Пришлось вынести, т.к. передать его функции goodbyeHandler (а она в последующем)

передаст его на closeServer не получится, иначе пропадет имплементация интерфейса handler -
ServeHTTP. Пиздец жопа съела трусы
*/
var server http.Server = http.Server{Addr: ":9090"}

func main() {

	log.Println("\033[34mThe server is running!\033[0m")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {

		defer wg.Done()

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error with server running: %#v\n", err)
		}
	}()

	http.HandleFunc("/h", helloHandler)
	http.HandleFunc("/g", goodbyeHandler)

	gracefulShutdown(&server)

	wg.Wait()
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

	closeServer(&server)

}

func closeServer(server *http.Server) {
	go func() { // Без горутины последний log.Printf не успеет отработать

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Error while server is shutting down by handler: %#v\n", err)
		}
	}()

	log.Printf("The server id down by the client request")
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
