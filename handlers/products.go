// Это goodbyeHandler
package handlers

import (
	"BuildingMicroservicesWithGo/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
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
func (p *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {

	p.l.SetOutput(rw)
	p.l.Println("POST - handler")

	product := &data.Product{} // Сюда положим десериализованные данные
	err := product.FromJSON(req.Body)

	if err != nil {
		http.Error(rw, "Error while deserialization json", http.StatusBadRequest)
	}

	p.l.Printf("Our new product: %#v\n", product)

	data.AddProduct(product)

}

// Обработчик PUT - запроса
func (p *Products) UpdateProducts(rw http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	// Отдает переменные, обнаруженные в URI, в виде map[string]string

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id from string to int", http.StatusBadRequest)
	}

	p.l.SetOutput(rw)
	p.l.Println("PUT - handler", id)

	product := &data.Product{}
	err = product.FromJSON(req.Body)

	if err != nil {
		http.Error(rw, "Error while deserialization json", http.StatusBadRequest)
	}

	data.UpdateProduct(id, product)
}
