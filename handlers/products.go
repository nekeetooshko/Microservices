// Это goodbyeHandler
package handlers

import (
	"BuildingMicroservicesWithGo/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// Роутер
func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// Роутинг запросов
	switch req.Method {
	// Если это метод GET
	case http.MethodGet:
		p.GetProducts(rw, req)
		return

	default: // Если метод не определен
		// rw.Write([]byte("Method not specified"))
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// Метод GET
func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {

	p.l.SetOutput(rw)

	productsList := data.GetProduts()

	err := productsList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError) // 500 ошибка
	}
}
