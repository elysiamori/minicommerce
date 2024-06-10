# Minicommerce API 🏪

#### Tech Stack 
- ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white)  Golang : https://go.dev/doc/ 
- ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=PostgreSQL&logoColor=white) PostgreSQL : https://www.postgresql.org/docs/ 
- Fiber : https://docs.gofiber.io/
- GORM : https://gorm.io/docs/

#### Library
- Fiber : go get github.com/gofiber/fiber/v2
- GORM Postgres : go get gorm.io/gorm gorm.io/driver/postgres
- Validator : go get github.com/go-playground/validator/v10
- Accounting Format : go get github.com/leekchan/accounting
- ENV : go get github.com/joho/godotenv


This API is an API about online shopping applications with minimalist features so that this API is made simple, this project was made to fulfill the final project of the mobile programming.

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### MINI-COMMERCE API [DOCS]

Key Feature : 
- CRUD product
- List product
- Details product
  
<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### Product CRUD
#### [POST] /api/products
Request :
- description : add product
- header : multipart/form-data
- body :
```
    - product_name : string
    - img_product : byte (uploaded image)
    - type_product : string
    - desc : string
    - price : integer
    - stock : integer
 ```

Response :
- status [201] created
- data :
    ```json
            {
                "id" : "integer",
                "product_name" : "string",
                "img_product" : "byte",
                "type_product" : "string",
                "desc" : "string",
                "price" : "integer",
                "stock" : "integer",
                "created_at" : "date"
            }
    ```
#### [GET] /api/products
Request : 
- description : get all products

Response :
- status [200] success
- data : 
    ```json
            {
                [
                    {
                        "id" : "integer",
                        "product_name" : "string",
                        "img_product" : "byte",
                        "type_product" : "string",
                        "desc" : "string",
                        "price" : "integer",
                        "stock" : "integer",
                    },

                    {
                        "id" : "integer",
                        "product_name" : "string",
                        "img_product" : "byte",
                        "type_product" : "string",
                        "desc" : "string",
                        "price" : "integer",
                        "stock" : "integer",
                    }
                ]
            }
    ```
###  [PUT] /api/products/:productID
Request :
- description : update product
- header : multipart/form-data
- params : productID(integer)
- body:
 ```
    - product_name : string
    - img_product : byte (uploaded image)
    - type_product : string
    - desc : string
    - price : integer
    - stock : integer
 ```
   
Response :
- status [200] success
- data :
    ```json
            {
                "Product":{
                    "id" : "integer",
                    "product_name" : "string",
                    "img_product" : "byte",
                    "type_product" : "string",
                    "desc" : "string",
                    "price" : "integer",
                    "stock" : "integer",
                    "created_at" : "date",
                    "updated_at" : "date"
                }
            }
    ```
    
#### [DELETE] /api/products/:productID
Request :
- description : delete product
- params : productID

Response :
- status [200] success
- data :
    ```json
            {
                "message" : "product successfully deleted"
            }
    ```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### List Product 
#### [GET] /api/list-products
Request :
- description : show all product to the menu

Response :
- status [200] success
- data :
    ```json
            {
                "Products": [
                    {
                        "id" : "integer",
                        "product_name" : "string",
                        "img_product" : "byte",
                        "type_product" : "string",
                        "price" : "integer",
                        "stock" : "integer"
                    }
                ]
            }
    ```
#### [POST] /api/search-product
Request : 
- description : search product
- header : application/json
- body :
    ```json
            {
                "product_name" : "string",
                "type_product" : "string"
            }
    ```
Response :
- status [200] success
- data :
    ```json
            {
                "product_name" : "string",
                "img_product" : "byte",
                "type_product" : "string",
                "price" : "integer",
                "stock" : "integer"
            }
    ```
<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### Detail product
#### [GET] /api/list-products/detail-product/productID
Request : 
- description : detail product
- params : productID(integer)

Response :
- status [200] success
- data :
    ```json
            {
                "product_name" : "string",
                "img_product" : "byte",
                "type_product" : "string",
                "price" : "integer",
                "stock" : "integer",
                "desc" : "string"
            }
    ```
