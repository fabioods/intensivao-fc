<h1 align="center">Welcome to Intensivão Full Cycle 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://twitter.com/fabioods" target="_blank">
    <img alt="Twitter: fabioods" src="https://img.shields.io/twitter/follow/fabioods.svg?style=social" />
  </a>
</p>

> Aplicação de conceitos de microserviços e comunicação entre eles.

# Author

👤 **Fábio dos Santos**

- Twitter: [@fabioods](https://twitter.com/fabioods)
- Github: [@fabioods](https://github.com/fabioods)
- LinkedIn: [@fabioods](https://linkedin.com/in/fabioods)

# Configuration

## Subir RabbitMQ & Reddis

Dentro da pasta docker executar:

```bash
docker-compose up -d
```

## Microserviço Product

Para executar o MS execute no terminal:

```bash
go run product.go
```

## Microserviço checkout

Executar no terminal:

```bash
export RABBITMQ_ERLANG_COOKIE="SWQOKODSQALRPCLNMEQG"
export RABBITMQ_DEFAULT_USER="rabbitmq"
export RABBITMQ_DEFAULT_PASS="rabbitmq"
export RABBITMQ_DEFAULT_VHOST="/"
export PRODUCT_URL="http://localhost:8081"
```

Para executar o MS execute no terminal:

```bash
go run checkout.go
```

## Microserviço catalog

Executar no terminal

```bash
export PRODUCT_URL="http://localhost:8081"
```

Para executar o MS execute no terminal:

```bash
go run catalog.go
```

## Microserviço de Order

Executar no terminal:

```bash
export RABBITMQ_ERLANG_COOKIE="SWQOKODSQALRPCLNMEQG"
export RABBITMQ_DEFAULT_USER="rabbitmq"
export RABBITMQ_DEFAULT_PASS="rabbitmq"
export RABBITMQ_DEFAULT_VHOST="/"
export RABBITMQ_CONSUMER_QUEUE="checkout_queue"
export REDIS_HOST="localhost:6379"
```

Para o MS de Order serão necessários 2 terminais, um deles para a Order consumir os dados do checkout e enviar para o MS de Payment, e outro para o MS de Order consumir a resposta do MS de Payment.
Para o primeiro caso devemos executar no terminal o seguinte comando, além dos export's acima:

```bash
go run order.go -opt checkout
```

Já para o segundo caso devemos executar no terminal o seguinte comando, além dos export's acima:

```bash
go run order.go -opt payment
```

## Microserviço de Payment

Executar no terminal

```bash
export RABBITMQ_ERLANG_COOKIE="SWQOKODSQALRPCLNMEQG"
export RABBITMQ_DEFAULT_USER="rabbitmq"
export RABBITMQ_DEFAULT_PASS="rabbitmq"
export RABBITMQ_DEFAULT_VHOST="/"
export RABBITMQ_CONSUMER_QUEUE="order_queue"
```

Para executar o MS execute no terminal:

```bash
go run payment.go
```

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
