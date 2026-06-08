# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

"Daily Study Topic" — learning project for CI/CD (GitHub Actions), REST API in Go, deploy to Render, static frontend on GitHub Pages, Docker (planned). Backend: Go + Gin. Frontend: plain HTML/CSS/JS.

## Commands

Run from `backend/`:

```sh
go run cmd/api/main.go        # start API server on :8080
go test ./...                 # run all tests
go test ./tests -run TestGetRandomTopic -v   # run a single test
go build ./...                # build, surfaces type errors across packages
```

Frontend (`frontend/`) is static — open `index.html` directly or serve the directory; `scripts.js` calls the backend at `http://localhost:8080`.

## Architecture

Backend follows a layered structure under `backend/internal/`:

- `cmd/api/main.go` — entrypoint; builds the Gin router, calls `routes.SetupRoutes`, listens on `:8080`.
- `routes/` — registers HTTP routes and wires them to handlers (e.g. `GET /health`, `GET /api/topic`).
- `handlers/` — Gin handler functions (HTTP layer); translate requests to service calls and write JSON responses.
- `service/` — business logic (e.g. `GetRandomTopic` picks a random topic from an in-memory slice).
- `models/` — plain data structs shared across layers (e.g. `Topic` with `ID`, `Title`, `Difficulty`, `Description`, all JSON-tagged).
- `tests/` — package `tests`, exercises the service layer directly (not handler/HTTP tests).

Data currently lives in an in-memory slice in `service/topic_service.go` — no database. Module path is `study-topics-cicd` (see `go.mod`); internal imports use `study-topics-cicd/internal/...`.

The frontend (`frontend/index.html`, `scripts.js`, `styles.css`) fetches `GET /api/topic` and renders `title`/`difficulty`/`description` into the page; the "Gerar novo tópico" button re-fetches a new random topic.
