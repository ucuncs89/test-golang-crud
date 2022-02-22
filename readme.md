Test CRUD Golang

used tech :

    Framework : GIN
    ORM : GORM
    Database : Postgresql

API Documentation :

    Get Product list
        Method : GET
        Url Endpoint : http://localhost:9000/product
        Payload: -

    Create Product
        Method : POST
        Url Endpoint : http://localhost:9000/product
        Payload: {“name”: “Dummy”, “price”: 2000,“stock”: 9}

    Detail Product
        Method : GET
        Url Endpoint : http://localhost:9000/product/:id
        Payload: -

    Update Product
        Method : PUT
        Url Endpoint : http://localhost:9000/product/:id
        Payload: {“name”: “Dummy”, “price”: 2000,“stock”: 9}

    Delete Product
        Method : DELETE
        Url Endpoint : http://localhost:9000/product/:id
        Payload: -

Run Command : go run main.go

Set Configuration DB Connection at models/config.go
change in func ConfigDB
