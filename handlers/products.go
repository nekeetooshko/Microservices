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

	data.AddProduct(product)

}

// Обработчик PUT - запроса
func (p *Products) updateProducts(id int, rw http.ResponseWriter, req *http.Request) {

	p.l.SetOutput(rw)
	p.l.Println("PUT - handler")

	product := &data.Product{}
	err := product.FromJSON(req.Body)

	if err != nil {
		http.Error(rw, "Error while deserialization json", http.StatusBadRequest)
	}

	data.UpdateProduct(id, product)

}
