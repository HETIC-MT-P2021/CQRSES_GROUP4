# Go API using CQRS and Event Sourcing

## Where am I ? 💡

This app allows you to build in a easy and fast way a go project using the combination of CQRS and EventSourcing patterns. It is built by the Group 4 of the MT5-P2021 team in Hétic.

* To submit bug reports and feature suggestions, or track changes:
   https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/issues

Don't forget to read/use commits and PR conventions before doing any change !

## Docs 📄
You can find our docs folder [here](https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/tree/master/docs)

## Health check ❤️

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

## Auth :lock:

There is 3 roles implemented. 

- Operator 
- Admin 
- Super admin

You can add a new account with the register routes.

You add have an JWT key for auth, with the login Routes. The token will be available in the response Header.

There is 3 roles implemented. Operator - Admin - Super admin

You can add a new account with the register routes.

You add have an JWT key for auth, with the login Routes. The token will be available in the response Header. 

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

## Article :newspaper:

### Get an article
**Request:**
```http request
GET /api/v1/articles/<aggregate_article_id>
Accept: application/json
Content-Type: application/json
{
  "title": "<title>",
  "description": "<description>"
}
```

**Successful Response:**
```http request
HTTP/1.1 201
{
  "id":  "<id>",
  "title": "<title>"
  "description": "<description>"
}
```

**Failed Response:**
```http request
HTTP/1.1 404 Not Found
```

### New article
**Request:**
```http request
POST /api/v1/articles
Accept: application/json
Content-Type: application/json
{
  "title": "<title>",
  "description": "<description>"
}
```

**Successful Response:**
```http request
HTTP/1.1 201
{
  "status": "created"
}
```

**Failed Response:**
```http request
HTTP/1.1 500 Internal Server Error
```

### Update an article
**Request:**
```http request
PUT /api/v1/articles/<aggregate_article_id>
Accept: application/json
Content-Type: application/json
{
  "title": "<title>",
  "description": "<description>"
}
```

**Successful Response:**
```http request
HTTP/1.1 201
{
  "status": "updated"
}
```

**Failed Response:**
```http request
HTTP/1.1 500 Internal Server Error
```

## Features 📘 (incoming)

- As a user, I want to read articles
- As an admin, I want to read, write and publish articles

## Libraries 📚 (incoming)

[Go-Swagger](https://github.com/go-swagger/go-swagger)

[Gin Gonic](https://github.com/gin-gonic/gin)

## Contributors ✨

<table>
  <tr>
    <td align="center"><a href="https://github.com/jibe0123"><img src="https://avatars.githubusercontent.com/u/13694014?s=400&u=979e9cdf62bcebe3e97740f83768fb41c8984a70&v=4" width="100px;" alt=""/><br /><sub><b>Jean Baptiste Agostin</b></sub></a><br /><a href="https://github.com/jibe0123" title="Developper">✏️</a>
    <td align="center"><a href="https://github.com/wyllisMonteiro"><img src="https://avatars2.githubusercontent.com/u/36091415?s=400&v=4" width="100px;" alt=""/><br /><sub><b>Wyllis Monteiro</b></sub></a><br /><a href="https://github.com/wyllisMonteiro" title="Developper">✏️</a>
    <td align="center"><a href="https://github.com/FaycalTOURE"><img src="https://avatars.githubusercontent.com/u/19931625?s=400&v=4" width="100px;" alt=""/><br /><sub><b>Fayçal Touré</b></sub></a><br /><a href="https://github.com/FaycalTOURE" title="Developper">✏️</a></td>
    <td align="center"><a href="https://github.com/valmrt77"><img src="https://avatars0.githubusercontent.com/u/36480710?v=4" width="100px;" alt=""/><br /><sub><b>Valentin Moret</b></sub></a><br /><a href="https://github.com/valmrt77" title="Developper">✏️</a></td>
  </tr>
</table>

## License 📑
[MIT](https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/blob/master/LICENSE)
