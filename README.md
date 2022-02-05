<h1 align="center">Welcome to Intensivão Full Cycle 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://twitter.com/fabioods" target="_blank">
    <img alt="Twitter: fabioods" src="https://img.shields.io/twitter/follow/fabioods.svg?style=social" />
  </a>
</p>

> Aplicação de conceitos de microserviços e comunicação entre eles.

## Author

👤 **Fábio dos Santos**

- Twitter: [@fabioods](https://twitter.com/fabioods)
- Github: [@fabioods](https://github.com/fabioods)
- LinkedIn: [@fabioods](https://linkedin.com/in/fabioods)

## Configuration

## Subir RabbitMQ & Reddis

Dentro da pasta docker executar:

```bash
docker-compose up -d
```

## Microserviço checkout

Executar no terminal:

```bash
export RABBITMQ_ERLANG_COOKIE="SWQOKODSQALRPCLNMEQG"
export RABBITMQ_DEFAULT_USER="rabbitmq"
export RABBITMQ_DEFAULT_PASS="rabbitmq"
export RABBITMQ_DEFAULT_VHOST="/"
export PRODUCT_URL="http://localhost:8080"
```

## Microserviço catalog

Executar no terminal

```bash
export PRODUCT_URL="http://localhost:8080"
```

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
