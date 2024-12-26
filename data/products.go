package data

import (
	"encoding/json"
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

// задает дефолтный энкодер через writer, которому будет передан http.ResponseWriter
func (p *Products) ToJSON(writer io.Writer) error {

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ") // Отступы
	return encoder.Encode(p)
}

func GetProduts() Products {
	return productList
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
