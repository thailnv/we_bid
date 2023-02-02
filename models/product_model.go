package models

type Product struct {
	ID                     int      `json:"id"`
	Name                   string   `json:"name"`
	Seller                 int      `json:"seller"`
	Image                  string   `json:"image"`
	Images                 []string `json:"images"`
	Description            string   `json:"description"`
	Sold                   int      `json:"sold"`
	Rating                 []int    `json:"rating"`
	Status                 int      `json:"status"`
	Price                  int      `json:"price"`
	PriceMax               int      `json:"price_max"`
	PriceMin               int      `json:"price_min"`
	PriceBeforeDiscount    int      `json:"price_before_discount"`
	PriceMaxBeforeDiscount int      `json:"price_max_before_discount"`
	PriceMinBeforeDiscount int      `json:"price_min_before_discount"`
}
