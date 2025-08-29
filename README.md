# Mini-POS-API-Challenge

## Database Setup

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

> **notes** that we are using docker network `portofolio_network` in `docker-compose.yaml`. it is important to make sure running app with docker work

\[---\]

---

## App Setup

### Docker

### Local

## API
