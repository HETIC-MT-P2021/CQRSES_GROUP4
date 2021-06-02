# Go API using CQRS and Event Sourcing

## Where am I ? 💡

This app allows you to build in a easy and fast way a go project using the combination of CQRS and EventSourcing patterns. It is built by the Group 4 of the MT5-P2021 team in Hétic.

- To submit bug reports and feature suggestions, or track changes:
  https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/issues

Don't forget to read/use commits and PR conventions before doing any change !

## Docs 📄

You can find our docs folder [here](https://github.com/HETIC-MT-P2021/CQRSES_GROUP4/tree/master/docs)

## Swagger doc 📄

To generate swagger doc

```sh
$ swag init -g cmd/app/main.go
```

Go in [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) to see API doc

## Features 📘 (incoming)

- As a user, I want to read articles
- As an admin, I want to read, write and publish articles

## Main Libraries 📚 (incoming)

- [Go-Swagger](https://github.com/go-swagger/go-swagger)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [JWT](github.com/kyfk/gin-jwt)
- [Elastic search](github.com/olivere/elastic/v7)
- [RabbitMQ Client](github.com/streadway/amqp)
- [Mysql](github.com/go-sql-driver/mysql)

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
