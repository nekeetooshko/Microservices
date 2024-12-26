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

	switch req.Method {

	case http.MethodGet:
		p.GetProducts(rw, req)
		return

	case http.MethodPost:
		p.addProduct(rw, req)
		return

	default:

		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {

	p.l.SetOutput(rw)
	p.l.Println("GET - handler")

	productsList := data.GetProduts()

	err := productsList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Serialization error", http.StatusInternalServerError) // 500 ошибка
	}
}

// Обработчик POST - запроса
func (p *Products) addProduct(rw http.ResponseWriter, req *http.Request) {

	p.l.SetOutput(rw)
	p.l.Println("POST - handler")

	product := &data.Product{} // Сюда положим десериализованные данные
	err := product.FromJSON(req.Body)

	if err != nil {
		http.Error(rw, "Error while deserialization json", http.StatusBadRequest)
	}

	p.l.Printf("Our new product: %#v\n", product)
}
