// Это helloHandler
package handlers

import (
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// DI + конструктор
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Метод, имплементящий http.Handler интерфейс
func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	h.l.SetOutput(rw)
	h.l.Println("Hello from client!")

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Error while reading the body: %e\n", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data: %s\n", string(data))

}
