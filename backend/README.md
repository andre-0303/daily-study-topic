# Backend — Daily Study Topic

API REST em Go (Gin) que serve tópicos de estudo aleatórios a partir de um banco SQLite.

## Stack

- Go
- [Gin](https://github.com/gin-gonic/gin) — router HTTP
- SQLite via `modernc.org/sqlite` (driver puro Go, sem cgo)
- `gin-contrib/cors` — CORS liberado pro frontend consumir a API

## Rodando localmente

```sh
go run cmd/api/main.go
```

Sobe o servidor em `:8080` (ou na porta definida em `PORT`). No boot, abre/cria o SQLite em `DB_PATH` (padrão `topics.db`), aplica `schema.sql` e popula com `seed.sql` — ambos embutidos no binário via `//go:embed`.

Variáveis de ambiente:

| Variável  | Padrão      | Descrição                                                 |
| --------- | ----------- | --------------------------------------------------------- |
| `PORT`    | `8080`      | Porta HTTP do servidor                                    |
| `DB_PATH` | `topics.db` | Caminho do arquivo SQLite (`:memory:` para banco efêmero) |

## Testes

```sh
go test ./...                                  # roda todos os testes
go test ./tests -run TestGetRandomTopic -v     # roda um teste específico
go build ./...                                 # build, expõe erros de tipo entre pacotes
```

## Endpoints

| Método | Rota          | Descrição                                  |
|--------|---------------|--------------------------------------------|
| GET    | `/health`     | Healthcheck — `{"status": "ok"}`           |
| GET    | `/api/topic`  | Retorna um tópico aleatório (JSON)         |

Resposta de `GET /api/topic`:

```json
{
  "id": 1,
  "title": "Goroutines e Channels",
  "difficulty": "Intermediário",
  "description": "..."
}
```

## Arquitetura

Estrutura em camadas sob `internal/`:

```text
cmd/api/main.go          # entrypoint — monta dependências, router, sobe servidor
internal/
├── db/                  # abre/seeda o SQLite (schema.sql + seed.sql via embed)
├── repository/          # acesso a dados — queries SQL (ex: GetRandom)
├── service/             # regra de negócio — orquestra repository
├── handlers/            # camada HTTP — traduz request/response em JSON
├── routes/              # registro de rotas, liga handlers ao router Gin
└── models/              # structs compartilhadas entre camadas (ex: Topic)
tests/                   # pacote `tests`, exercita a camada service
```

Fluxo de uma requisição: `routes` → `handlers` → `service` → `repository` → SQLite.

Módulo Go: `study-topics-cicd` (ver `go.mod`); imports internos usam `study-topics-cicd/internal/...`.
