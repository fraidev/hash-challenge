# hash-challenge

## Introduction 
This project proposes two microservices working together to display a product list with discounts per users.

The first microservice is the product service, written with Typescript, using NodeJS. 

The second microservice is the discount service, written with GO Language.

In this monolith we also have a another project, the proto-generator, written with Typescript and NodeJS, that project take the proto file and generate the GRPC interfaces for both services.

## Flow

The product service has a web service with only one endpoint, the "/product". This endpoint expect to receive user id as header request. This service will get all the products from database and send via GRPC only the product ids to discount service with the user id. The discount service will get the user and produts from database by ids, and will calculate the discount with a rule engine. After will response de product service with the value and percentage of the discount. And finaly product service will response the web request with a json of all products and their discounts.


## How to execute / Deploy

Runs product service, discount service, postgres database and apply some records in database. Everything inside the Docker.

```sh
git clone https://github.com/fraifelipe/hash-challenge

cd hash-challenge

docker-compose up --build
```

## How to Test

Birth date case
```sh
curl -H "USER-ID: 9f4abbe9-113c-4099-8bef-2184d030f2c3" localhost:8080/product
```

Not birth date case
```sh
curl -H "USER-ID: 3f045e1b-3ff7-429c-9ca1-e4e7585b48a6" localhost:8080/product
```

## Disclaimer
This project has tests, but all database intergration tests only make sense with a database with the records that is in populate script. In real world projects normaly has tests that add more records and check if this records were added. But in this project only get records is a requirement, thus, that don't have add records methods and I only check the records from script.

Normally at a bigger project I would add a library or struct for helps with DI and IOC, helping scaling the code and to test, but I didn't see a code smell or difficult to test in this small project.