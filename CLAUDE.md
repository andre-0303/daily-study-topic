# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

"Daily Study Topic" — learning project for CI/CD (GitHub Actions), REST API in Go, deploy to Render, static frontend on GitHub Pages, Docker. Backend: Go + Gin + SQLite. Frontend: plain HTML/CSS/JS.

## Commands

Run from `backend/`:

```sh
go run cmd/api/main.go        # start API server on :8080
go test ./...                 # run all tests
go test ./tests -run TestGetRandomTopic -v   # run a single test
go build ./...                # build, surfaces type errors across packages
```

Docker (from `backend/`):

```sh
docker build -t study-topics-api .
docker run --rm -p 8080:8080 study-topics-api
```

Frontend (`frontend/`) is static — open `index.html` directly or serve the directory; `scripts.js` calls the backend at `http://localhost:8080`.

## Architecture

Backend follows a layered structure under `backend/internal/`:

- `cmd/api/main.go` — entrypoint; opens/seeds the SQLite DB (`db.Open`), wires `repository` → `service` → `handlers`, builds the Gin router, calls `routes.SetupRoutes`, listens on `:8080` (or `$PORT`).
- `db/` — `Open(path)` opens (or creates) a SQLite file via `modernc.org/sqlite`, applies `schema.sql` and `seed.sql` (both embedded with `//go:embed`); pass `:memory:` for an ephemeral DB.
- `repository/` — data access; `TopicRepository.GetRandom` runs `SELECT ... ORDER BY RANDOM() LIMIT 1` against the `topics` table.
- `routes/` — registers HTTP routes and wires them to handlers (e.g. `GET /health`, `GET /api/topic`).
- `handlers/` — Gin handler functions (HTTP layer); translate requests to service calls and write JSON responses.
- `service/` — business logic (e.g. `GetRandomTopic` delegates to the repository).
- `models/` — plain data structs shared across layers (e.g. `Topic` with `ID`, `Title`, `Difficulty`, `Description`, all JSON-tagged).
- `tests/` — package `tests`, exercises the service layer directly (not handler/HTTP tests).

Data lives in a SQLite database (file path from `$DB_PATH`, default `topics.db`), seeded at startup from embedded `schema.sql`/`seed.sql`. CORS is enabled via `gin-contrib/cors`. A `Dockerfile` in `backend/` builds the API on `golang:1.26-alpine` (matches `go.mod`'s `go 1.26.3` directive). Module path is `study-topics-cicd` (see `go.mod`); internal imports use `study-topics-cicd/internal/...`.

The frontend (`frontend/index.html`, `scripts.js`, `styles.css`) fetches `GET /api/topic` and renders `title`/`difficulty`/`description` into the page; the "Gerar novo tópico" button re-fetches a new random topic.
