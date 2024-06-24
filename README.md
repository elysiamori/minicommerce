# Minicommerce API 🛍

The purpose of this API is to understand the basic concepts of crud operations in a rest api by implementing a project for an online store. This api was also developed as a final project for a mobile programming course and will be consumed by a flutter-based frontend as part of the project.

#### Tech Stack 
- Golang : https://go.dev/doc/ 
- PostgreSQL : https://www.postgresql.org/docs/ 
- Fiber : https://docs.gofiber.io/
- GORM : https://gorm.io/docs/

#### Library
- Fiber : ```go get github.com/gofiber/fiber/v2```
- GORM Postgres : ```go get gorm.io/gorm gorm.io/driver/postgres```
- Validator : ```go get github.com/go-playground/validator/v10```
- Accounting Format : ```go get github.com/leekchan/accounting```

#### Deployment
- App : https://fly.io
- Database : https://railway.app

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### MINI-COMMERCE API [DOCS]

Key Feature : 
- CRUD product
- List product
- Search product
- Details product

API URL : https://minicommerce.fly.dev
  
<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### Product CRUD
#### [POST] /api/products  🚀
Request :
- description : add product
- header : multipart/form-data
- body :
```
    - product_name : string
    - img_product : byte (uploaded image)
    - type_product : string
    - desc : string
    - sold : int
    - location : string
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
                "sold" : "int",
                "location" : "string",
                "price" : "integer",
                "stock" : "integer",
                "created_at" : "date"
            }
    ```
#### [GET] /api/products 📩
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
                        "sold" : "int",
                        "location" : "string",
                        "price" : "integer",
                        "stock" : "integer",
                    },

                    {
                        "id" : "integer",
                        "product_name" : "string",
                        "img_product" : "byte",
                        "type_product" : "string",
                        "desc" : "string",
                        "sold" : "int",
                        "location" : "string",
                        "price" : "integer",
                        "stock" : "integer",
                    }
                ]
            }
    ```
###  [PUT] /api/products/:productID 📤
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
    - sold : int
    - location : string
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
                    "sold" : "int",
                    "location" : "string",
                    "price" : "integer",
                    "stock" : "integer",
                    "created_at" : "date",
                    "updated_at" : "date"
                }
            }
    ```
    
#### [DELETE] /api/products/:productID ❌
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
#### [GET] /api/list-products 📃
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
                        "sold" : "int",
                        "location" : "string",
                        "price" : "integer",
                        "stock" : "integer"
                    }
                ]
            }
    ```
#### [GET] /api/search-product 🔍
Request : 
- description : search product
- params : query paramaeter (product_name, type_product)
Response :
- status [200] success
- data :
    ```json
            {
                "product_name" : "string",
                "img_product" : "byte",
                "type_product" : "string",
                "sold" : "int",
                "location" : "string",
                "price" : "integer",
                "stock" : "integer"
            }
    ```
<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### Detail product
#### [GET] /api/list-products/detail-product/productID 📝
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
                "sold" : "int",
                "location" : "string",
                "price" : "integer",
                "stock" : "integer",
                "desc" : "string"
            }
    ```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

# Importing Postman API Documentation

### Steps to Import API Documentation into Postman

##### 1. Install Postman
- Download and install Postman from [the official website](https://www.postman.com/downloads/).
- Follow the installation instructions for your operating system.

#### 2. Download the `minicommerce.json` File
- Ensure you have the `minicommerce.json` file, download or copied into your local machine.

#### 3. Import the Collection into Postman
- Open Postman.
- Click on the "Import" button in the top left corner of the Postman app.
- Select the "File" tab.
- Click on "Choose Files" and select the `minicommerce.json` file from your local machine.
- Postman will import the collection, and you will see it appear in the Collections sidebar.

#### 4. Using the Imported Collection
- Once imported, you will see the collection in the Collections sidebar.
- Click on the collection to browse through the API requests.
- Select any request and click the "Send" button to execute it and view the response.

By following these steps, you will have the API documentation and requests ready to use in Postman, making it easier to interact with the API endpoints provided in the `minicommerce.json` file.


