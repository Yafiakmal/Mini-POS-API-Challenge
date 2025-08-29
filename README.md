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
docker volume create postgres_alpine_data
```

4. run docker compose

```bash
docker compose up -d
```

> - **notes** that we are using docker network `portofolio_network` in `docker-compose.yaml`. it is important to make sure running app with docker work
> - golang + gorm would **automatically migrate** models defined hardcoded.

\[---\]

---

## App Setup `Linux`

### Docker

> make sure docker daemon and docker cli is active and installed in your pc `sudo systemctl status docker --no-pager && echo "\n\n" && docker --version`

1. download app image i've built in the release
2. load image

```bash
docker load -i minipos-golang_1.0.0.tar
```

3. run image

```bash
docker run --network portofolio_network -p 8080:8080 --rm minipos-golang:1.0.0
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
DB_HOST= localhost
DB_USER= postgres
DB_PASSWORD= postgres
DB_NAME= pos
DB_PORT= 5432
DB_SSLMODE= disable
DB_TIMEZONE= Asia/Jakarta

# gorm would automatically migrate the models
MODE= production  # not gonna drop the tables first
# MODE= development # gonna drop the tables first for fresh models
```

3. Run with golang

   > - make sure **golang installed** on your local machine `go version`
   > - make sure you have **run the database** in the previous step `docker ps`

```bash
  go mod tidy # download library required in go.mod
  go run cmd/main.go
```

\[---\]

---

## API
