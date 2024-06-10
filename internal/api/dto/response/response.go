package response

type ProductAddResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product"`
	Desc        string `json:"desc"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
}

type ProductAllResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product"`
	Desc        string `json:"desc"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductUpdatedResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product"`
	Desc        string `json:"desc"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductListResponse struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
}

type ProductSearchResponse struct {
	ProductName string `json:"product_name"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product"`
	Desc        string `json:"desc"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
}

type ProductDetailProduct struct {
	ProductName string `json:"product_name"`
	ImgProduct  []byte `json:"img_product"`
	TypeProduct string `json:"type_product"`
	Price       string `json:"price"`
	Stock       int    `json:"stock"`
	Desc        string `json:"desc"`
}
