package request

type ProductAddRequest struct {
	ProductName string `json:"product_name" validate:"required"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product" validate:"required"`
	Desc        string `json:"desc" validate:"required"`
	Sold        int    `json:"sold" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}

type ProductUpdatedRequest struct {
	ProductName string `json:"product_name" validate:"required"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product" validate:"required"`
	Desc        string `json:"desc" validate:"required"`
	Sold        int    `json:"sold" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}
