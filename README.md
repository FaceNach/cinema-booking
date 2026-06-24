# Cinema Seat Booking

Small Go app for booking cinema seats. It serves a static frontend, exposes a JSON API, and uses Redis so only one user can hold a seat at a time.

## Overview

- Go backend with `net/http`
- Static frontend in `static/index.html`
- Redis-backed seat holds and confirmations
- Holds expire after 2 minutes
- App URL: `http://localhost:8080`
- Redis URL: `localhost:6379`

## Structure

```text
cmd/main.go                  # Server, routes, sample movies
internal/adapters/redis/     # Redis client
internal/booking/            # Booking domain, service, handlers, stores
internal/utils/              # JSON helpers
static/index.html            # Frontend
docker-compose.yaml          # Redis and Redis Commander
```

## Run

Requirements: Go 1.26.3 or compatible, Docker, Docker Compose.

```sh
docker compose up -d
go run ./cmd
```

Then open:

```text
http://localhost:8080
```

Redis Commander is available at:

```text
http://localhost:8081
```

Stop services:

```sh
docker compose down
```

## Tests

Tests need Redis running locally.

```sh
docker compose up -d redis
docker compose exec redis redis-cli FLUSHDB
go test ./...
```

## API

```text
GET    /movies
GET    /movies/{movieID}/seats
POST   /movies/{movieID}/seats/{seatID}/hold
PUT    /sessions/{sessionID}/confirm
DELETE /sessions/{sessionID}
```

POST, PUT, and DELETE requests use:

```json
{
  "user_id": "user-123"
}
```

## Redis

Main keys:

```text
seat:{movieID}:{seatID}
session:{sessionID}
```

Seats are held with Redis `SET NX`, so only the first user can reserve a free seat.

---

# Cinema Seat Booking

Pequena app en Go para reservar asientos de cine. Sirve un frontend estatico, expone una API JSON y usa Redis para que solo un usuario pueda retener un asiento a la vez.

## Resumen

- Backend en Go con `net/http`
- Frontend estatico en `static/index.html`
- Retenciones y confirmaciones guardadas en Redis
- Las retenciones expiran despues de 2 minutos
- URL de la app: `http://localhost:8080`
- Redis: `localhost:6379`

## Estructura

```text
cmd/main.go                  # Servidor, rutas, peliculas de ejemplo
internal/adapters/redis/     # Cliente Redis
internal/booking/            # Dominio, servicio, handlers y stores
internal/utils/              # Helpers JSON
static/index.html            # Frontend
docker-compose.yaml          # Redis y Redis Commander
```

## Ejecutar

Requisitos: Go 1.26.3 o compatible, Docker y Docker Compose.

```sh
docker compose up -d
go run ./cmd
```

Luego abrir:

```text
http://localhost:8080
```

Redis Commander queda disponible en:

```text
http://localhost:8081
```

Detener servicios:

```sh
docker compose down
```

## Tests

Los tests necesitan Redis corriendo localmente.

```sh
docker compose up -d redis
docker compose exec redis redis-cli FLUSHDB
go test ./...
```

## API

```text
GET    /movies
GET    /movies/{movieID}/seats
POST   /movies/{movieID}/seats/{seatID}/hold
PUT    /sessions/{sessionID}/confirm
DELETE /sessions/{sessionID}
```

Las peticiones POST, PUT y DELETE usan:

```json
{
  "user_id": "user-123"
}
```

## Redis

Claves principales:

```text
seat:{movieID}:{seatID}
session:{sessionID}
```

Los asientos se retienen con `SET NX` de Redis, asi que solo el primer usuario puede reservar un asiento libre.
