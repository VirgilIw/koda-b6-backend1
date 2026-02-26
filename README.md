# Simple User CRUD API (Go + Gin)

API sederhana untuk CRUD user menggunakan Go dan Gin.
Data disimpan sementara di memory.

## Run

```bash
go run main.go
```

Server berjalan di:

```
http://localhost:8888
```

## Endpoints

### Get All Users

```
GET /users
```

### Get User

```
GET /users/:id
```

### Create User

```
POST /users
```

Body

```json
{
  "email": "user@mail.com",
  "password": "123"
}
```

### Update User

```
PATCH /users/:id
```

### Delete User

```
DELETE /users/:id
```

## Note

Data akan hilang jika server direstart karena hanya disimpan di memory.
