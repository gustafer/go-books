# Simple GO JSON Books Api using Fiber and PostgreSQL with GORM
Simple api which uses Fiber V2 (should update to v3 soon) to create, read and delete books.
## Setup
### 1. First start by running the docker compose for postgres:
```
docker compose up -d
```
### 2. Start up the server

```
go run ./cmd/api-server
```

## Routes:
### Check it out at [routes](./cmd/api-server/routes/routes.go)
```
POST /api/book
// create a book using JSON body with author, title and publisher

GET /api/books
// gets all books info in JSON

GET /api/book/:id
// get book by ID in JSON (you can check it out the books id with the /api/books)

DELETE /api/book/:id
// delete a book by ID
```
