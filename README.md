# Blog Backend API

## Introdunction
Backend API for a simple blog

## Quick start

1. `git clone http://github.com/asktao/blog && cd blog`
2. Run using `docker-compose up -d`


When the application is run, the following endpoints will be available:

##### Create a article

- Host: `http://127.0.0.1:8080` (by default)

- Method: `POST`

- Path: `/articles`

- Request Body:
```
{
    "title": "Hello World",
    "content": "Blog API"
    "author": "Asktao"
}
```
- Response Header: `HTTP 201`
- Response Body:
```
{
    "status": 201,
    "message": "Success",
    "data": {
        "id": <article_id>
    }
}
```

##### Get article by id

- Host: `http://127.0.0.1:8080`

- Method: `GET`

- Path: `/articles/{article_id}`

- Response Header `HTTP 200`

- Response Body:
```
{
    "status": 200,
    "message": "Success"
    "data": {
        "id": <article_id>,
        "title": <article_title>,
        "content": <article_content>,
        "author": <article_author>,
     }
}
```

##### Get all article

- Host: `http://127.0.0.1:8080`

- Method: `GET`

- Path: `/articles`

- Query:
```
limit=10
offset=0
```

- Response Header `HTTP 200`

- Response Body:
```
{

    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      },
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}
```

### Using with Docker
* [Docker](https://www.docker.com)



## Dependent package
* [Go MYSQL Drive](github.com/go-sql-driver/mysql)
* [mux](github.com/gorilla/mux)
* [gorm](https://github.com/jinzhu/gorm)
* [dotenv](https://github.com/joho/godotenv)
