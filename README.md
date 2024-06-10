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
- ENV : ```go get github.com/joho/godotenv```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

### MINI-COMMERCE API [DOCS]

Key Feature : 
- CRUD product
- List product
- Search product
- Details product
  
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
                        "price" : "integer",
                        "stock" : "integer"
                    }
                ]
            }
    ```
#### [POST] /api/search-product 🔍
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
                "price" : "integer",
                "stock" : "integer",
                "desc" : "string"
            }
    ```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif">

# Importing Thunder Client API Documentation

To open the documentation and test the API using the Thunder Client extension in Visual Studio Code, please follow these steps to import the `minicommerce.json` file:

## Steps to Import

1. **Install Thunder Client Extension**:
   - Open Visual Studio Code.
   - Go to the Extensions view by clicking on the Extensions icon in the Activity Bar on the side of the window.
   - Search for "Thunder Client" and click the Install button.

2. **Download the `minicommerce.json` File**:
   - Download or copy paste `minicommerce.json` file on your local machine.

3. **Import the Collection into Thunder Client**:
   - Open the Thunder Client extension by clicking on its icon in the Activity Bar.
   - Click on the "Collections" tab.
   - Click on the "Import" button.
   - Select the `minicommerce.json` file from your local machine.

4. **Using the Imported Collection**:
   - Once imported, you will see the collection in the Thunder Client.
   - You can now browse through the API requests, execute them, and see the responses.

By following these steps, you will have the API documentation and requests ready to use in Thunder Client, making it easier to interact with the API endpoints provided in the `minicommerce.json` file.

