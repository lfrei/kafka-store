# kafka-store

Small application to load and query Products from a store using REST or Kafka

## Store

In-Memory Store holding the Products

## Kafka

Add products to the store by sending messages to Kafka:
* Topic: `product`
* Key: `{id}`
* Value: `product` JSON-Payload

## REST

Add / Query products from the store using REST:
* Add: `POST /product/{id}` with `product` JSON-Payload
* Query: `GET /product/{id}`
    