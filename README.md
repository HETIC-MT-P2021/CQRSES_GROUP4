# API Documentation

## Where am I ? 💡

This app allows you to build in a easy and fast way a go project using the combination of CQRS and EventSourcing patterns.

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

## Auth 🔒

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

## Features 📘 (incoming)

- Start a go project
- Projects will be build using CQRS and EventSourcing 

## Libraries 📚 (incoming)

## Contributors ✨

<table>
  <tr>
    <td align="center"><a href="https://github.com/jibe0123"><img src="https://avatars.githubusercontent.com/u/13694014?s=400&u=979e9cdf62bcebe3e97740f83768fb41c8984a70&v=4" width="100px;" alt=""/><br /><sub><b>Jean Baptiste Agostin</b></sub></a><br /><a href="https://github.com/jibe0123" title="Developper">✏️</a>
    <td align="center"><a href="https://github.com/wyllisMonteiro"><img src="https://avatars2.githubusercontent.com/u/36091415?s=400&v=4" width="100px;" alt=""/><br /><sub><b>Wyllis Monteiro</b></sub></a><br /><a href="https://github.com/wyllisMonteiro" title="Developper">✏️</a>
    <td align="center"><a href="https://github.com/valmrt77"><img src="https://avatars0.githubusercontent.com/u/36480710?v=4" width="100px;" alt=""/><br /><sub><b>Valentin Moret</b></sub></a><br /><a href="https://github.com/valmrt77" title="Developper">✏️</a></td>
  </tr>
</table>

## License 📑
[MIT](https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/blob/master/LICENSE)
