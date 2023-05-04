# E-Market

>  Sample REST API server built with Go, Gin and MongoDB


## Usage

1. Clone the repository and change to the project directory.

    ```bash
    git clone https://github.com/jenishayadav/go-gin-e-market.git
    cd go-gin-e-market
    ```

2. Install dependencies.

    ```bash
    go get .
    ```

3. Start a mongodb server at local using Docker.

    ```bash
    docker run --name mongodb -d -p 27017:27017 mongodb/mongodb-community-server
    ```

4. Mongo initial steps.

    4.1 Start mongo shell.
    ```bash
    docker exec -it mongodb mongosh
    ```
    
    4.2 Create a database and two mongo collections, one for `Order` and another for `Product`
    
    ```
    use josh_assignment
    ```
    ```
    db.createCollection("order")
    db.createCollection("product")
    ```

    4.3 Verify the collections.
    ```
    db.getCollectionNames()
    ```

    Exit `mongosh`.


4. Run the server.

    ```bash
    go run .
    ```


## Sample requests (cURL)

0. Health check
    ```bash
    curl --location --request GET 'http://localhost:8080/api/healthcheck'
    ```

1. Create product

    ```bash
    curl --location --request POST 'http://localhost:8080/api/products/' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "name": "DermaCO Vitamin-C Serum",
            "availableQty": 6,
            "price": 1000.0,
            "category": "Premium"
        }'
    ```


2. Get all products

    ```bash
    curl --location --request GET 'http://localhost:8080/api/products/'
    ```

3. Get product by ID. (Replace `<PRODUCT_ID>` with a sample ID)

    ```bash
    curl --location --request GET "http://localhost:8080/api/products/<PRODUCT_ID>"
    ```


4. (Partial) Update product by ID. (Replace `<PRODUCT_ID>` with a sample ID)

    ```bash
    curl --location --request PATCH "http://localhost:8080/api/products/<PRODUCT_ID>" \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "name": "DermaCo Vitamin-C Serum",
            "price": 1100.0,
        }'
    ```


5. Delete product by ID. (Replace `<PRODUCT_ID>` with a sample ID)

    ```bash
    curl --location --request DELETE "http://localhost:8080/api/products/<PRODUCT_ID>"
    ```

6. Create order. (Replace product ids placeholders)

    ```bash
    curl --location --request POST 'http://localhost:8080/api/orders/' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "orderItems": [
                {
                    "productId":"<PRODUCT_ID_1>",
                    "quantity": 3
                },
                {
                    "productId":"<PRODUCT_ID_2>",
                    "quantity": 5
                }
            ]
        }'
    ```

    **NOTE**: This takes care of maximum order quantity and also checks if ordered quantity is less than or equal to available quantity of order item.

    **NOTE**: You will see change in available quantity of ordered products as well. Try the request for fetching all products and validate.


7. Get all orders
    ```bash
    curl --location --request GET "http://localhost:8080/api/orders/"
    ```

8. Get order by ID
    ```bash
    curl --location --request GET "http://localhost:8080/api/orders/<ORDER_ID>"
    ```

9. Update order by ID

    ```bash
    curl --location --request PATCH "http://localhost:8080/api/orders/<ORDER_ID>" \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "orderStatus": "Dispatched"
        }'
    ```

    **NOTE**: You will see the update in dispatch date as well.


10. Delete order by ID
    ```bash
    curl --location --request DELETE "http://localhost:8080/api/orders/<ORDER_ID>"
    ```
