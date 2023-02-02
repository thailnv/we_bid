package mock

import "we_bid/models"

var ProductList = []models.Product{
	{
		ID:                     1,
		Name:                   "Product_1",
		Price:                  10000,
		Image:                  "hppts://imagebb.com/image_1.png",
		Images:                 []string{},
		PriceMax:               0,
		PriceMin:               0,
		PriceBeforeDiscount:    0,
		PriceMinBeforeDiscount: 0,
		PriceMaxBeforeDiscount: 0,
		Description:            "",
		Sold:                   0,
		Status:                 0,
		Rating:                 []int{},
		Seller:                 0,
	},
}
