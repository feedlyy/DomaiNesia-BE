package app

type Dvd struct {
	Id 		 int 	`form:"id" json:"id"`
	Title    string `form:"title" json:"title"`
	Price    int    `form:"price" json:"price"`
	Quantity int    `form:"quantity" json:"quantity"`
	Category string `form:"category" json:"category"`
}
