package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// ----------------------------------------------------------------------------------------------------

// задает дефолтный энкодер через writer, которому будет передан http.ResponseWriter
func (p *Products) ToJSON(writer io.Writer) error {

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ") // Отступы
	return encoder.Encode(p)
}

// Создаем декодер и читаем все данные из io.Reader
func (p *Product) FromJSON(r io.Reader) error {
	newDecoder := json.NewDecoder(r) // Читает из r
	return newDecoder.Decode(p)      // Читает данные и десериализует их в p
}

// ----------------------------------------------------------------------------------------------------

func GetProduts() Products {
	return productList
}

// Добавляет товар к уже готовому списку товаров
func AddProduct(prod *Product) {

	prod.ID = getNextId()
	productList = append(productList, prod)
}

// Какой-то буллщит, буллщит, ультраМяу
func UpdateProduct(id int, prod *Product) error {

	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	productList[pos] = prod

	return nil

}

var ErrProductNotFound = fmt.Errorf("Product not found")

// Ищет товар по указанному id-шнику
func findProduct(id int) (*Product, int, error) {
	for i, v := range productList {
		if v.ID == id {
			return v, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// Получает id-шник ласт элемента
func getNextId() int {
	// Нужна она для корректного добавления данных в БД. Ведь в запросе можно указать любой id-шник, а
	// добавлять стоит по конкретному
	return productList[len(productList)-1].ID + 1
}

var productList = []*Product{

	{
		ID:          1,
		Name:        "Латте",
		Description: "Пенистый молочный кохфе",
		Price:       245,
		SKU:         "абв323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	{
		ID:          2,
		Name:        "Эспрессо",
		Description: "Маленький и крепкий кохфе без молока",
		Price:       199,
		SKU:         "фжд34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
