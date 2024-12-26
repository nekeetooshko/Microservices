// Это goodbyeHandler
package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

// DI + конструктор
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// Метод, имплементящий http.Handler интерфейс
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	g.l.SetOutput(rw)
	g.l.Println("Goodbye from client!")
}
