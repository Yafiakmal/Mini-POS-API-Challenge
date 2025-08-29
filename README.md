# Mini-POS-API-Challenge

## Database Setup `Linux`

### Docker

> make sure docker daemon and docker cli is active and installed in your pc `sudo systemctl status docker --no-pager && echo "\n\n" && docker --version`

1. pull image

```bash
docker pull postgres:13.22-alpine3.22
```

2. create `docker-compose.yaml`

```yaml
services:
  postgres:
    image: postgres:13.22-alpine3.22
    container_name: postgres_container
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: pos
    volumes:
      - postgres_alpine_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    networks:
      - portofolio_network

networks:
  portofolio_network:
    driver: bridge
    external: true
volumes:
  postgres_alpine_data:
    name: postgres_alpine_data
    external: true
```

3. create docker volume

to make sure `docker-compose.yaml` above works and postgres data persistence

```bash
docker network create portofolio_network
docker volume create postgres_alpine_data
```

4. run docker compose

```bash
docker compose up -d postgres
```

> - **notes** that we are using docker network `portofolio_network` in `docker-compose.yaml`. it is important to make sure running app with docker work
> - golang + gorm would **automatically migrate** models defined hardcoded. so you don't need to migrate it manually

\[---\]

---

## App Setup `Linux`

chose one

### Docker

> make sure docker daemon and docker cli is active and installed in your pc `sudo systemctl status docker --no-pager && echo "\n\n" && docker --version`

1. download app image i've built in the [release](https://github.com/Yafiakmal/Mini-POS-API-Challenge/releases)
2. load image

```bash
docker load -i minipos-golang_v1.0.10.tar
```

3. Create `.env` file in project root
   > `DB_HOST` has different value if you run app locally

```bash
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pos
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Jakarta

# gorm would automatically migrate the models
# APP_ENV=production  # not gonna drop the tables first
APP_ENV=development # gonna drop the tables first for fresh models
```

4. run image

> - make sure you have **run the database** in the previous step `docker ps`

```bash
docker run --env-file .env --network portofolio_network -p 8080:8080 --rm minipos-golang:v1.0.10
```

### Local

if you want to run it locally

1. clone this repo

```bash
git clone https://github.com/Yafiakmal/Mini-POS-API-Challenge.git

cd Mini-POS-API-Challenge/
```

2. Create `.env` file in project root

```bash
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pos
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Jakarta
```

3. Run with golang

> - make sure **golang installed** on your local machine `go version`
> - make sure you have **run the database** in the previous step `docker ps`

```bash
  go mod tidy # download library required in go.mod
  # APP_ENV=development would reset all data inserted
  APP_ENV=development go run cmd/main.go
```

\[---\]

---

## API

### product CRUD

1. CREATE sample data

```bash
curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pencil",
    "price": 4000,
    "stock": 110
  }'

curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Spidol",
    "price": 7000,
    "stock": 70
  }'

curl -X POST http://localhost:8080/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "White Board",
    "price": 35000,
    "stock": 40
  }'

```

2. READ

```bash
curl -X GET http://localhost:8080/products

```

3. UPDATE

```bash
curl -X PUT http://localhost:8080/product/3 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Papan Tulis",
    "price": 35000,
    "stock": 55
  }'

```

4. DELETE

```bash
curl -X DELETE http://localhost:8080/product/1

```

### transaction

1. CREATE

```bash
curl -X POST http://localhost:8080/transaction \
  -H "Content-Type: application/json" \
  -d '[
    {
      "product_id": 2,
      "quantity": 3
    },
    {
      "product_id": 3,
      "quantity": 5
    }
  ]'

```

2. GET HISTORY

```bash
curl -X GET http://localhost:8080/transactions
```
