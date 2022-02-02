package items

type Item struct {
	Id                string      `json:"id"`
	Seller            int         `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Pictures  `json:"pictures"`
	Video             string      `json:"video"`
	Price             float64     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Pictures struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}
