# API Documentation

## Health check

**Request:**
```http request
GET /api/v1/health_check
```
**Successful Response:**
```http request
HTTP/1.1 200 OK
{
  "message": "ok"
}
```

## Auth

### Login
**Request:**
```http request
GET /api/v1/login
Accept: application/json
Content-Type: application/json
{
	"username": "<username>",
	"password": "<password>"
}
```

**Successful Response:**
```http request
HTTP/1.1 200
Authorization: 	Bearer <token>
```

**Failed Response:**
```http request
HTTP/1.1 403 Forbidden 
HTTP/1.1 500 Internal server error
``` 

### Register
*More type of account need to be implemented*


*Actually only `admin` role can create account of role `operator`*

**Request:**
```http request
POST /api/v1/auth/register
Accept: application/json
Content-Type: application/json
{
	"username": "<username>",
	"password": "<password>",
	"email": "<email>"
}
```

**Successful Response:**
```http request
HTTP/1.1 201
Authorization: 	Bearer <string>
{
  "message": "Account_created"
}
```

**Failed Response:**
```http request
HTTP/1.1 403 Forbidden 
HTTP/1.1 500 Internal server error
``` 