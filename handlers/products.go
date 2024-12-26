// Это goodbyeHandler
package handlers

import (
	"BuildingMicroservicesWithGo/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	case http.MethodPut:

		// Дабы обновить что-то по id-шнику, его нужно достать. Регулярки. Дада, блять. Регулярки
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(req.URL.Path, -1) // Будет хранить наш id-шник

		if len(group) != 0 && len(group[0]) != 2 {
			http.Error(rw, "Error with regExp", http.StatusBadRequest)
		}

		// Отынтуем id
		id, err := strconv.Atoi(group[0][1])
		if err != nil {
			http.Error(rw, "Error with converting string data into int", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, req)

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

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
