# Frontend — Daily Study Topic

Frontend estático (HTML/CSS/JS puro) que consome a API e exibe o tópico do dia num card com visual de ficha de catálogo carimbada.

## Rodando localmente

Sem build step — abra `index.html` direto no navegador ou sirva a pasta com qualquer servidor estático:

```sh
# qualquer servidor estático funciona, ex:
npx serve .
```

`scripts.js` detecta o ambiente pelo hostname:

- `localhost` / `127.0.0.1` → `http://localhost:8080/api/topic`
- produção → `https://daily-study-topic.onrender.com/api/topic`

Para testar contra a API local, suba o backend antes (veja [../backend/README.md](../backend/README.md)).

## Estrutura

```text
index.html    # markup — card do tópico, kicker, botão "Gerar novo tópico"
styles.css    # estilo "ficha de catálogo": papel/sépia, carimbo, perfuração lateral
scripts.js    # fetch em /api/topic, popula o card, trata erro de carregamento
images/       # logo e assets estáticos
```

## Comportamento

- Ao carregar a página, `fetchTopic()` busca um tópico aleatório e popula `title`, `difficulty`, `description` e o índice de catálogo (`Nº 0XX`).
- Botão **"Gerar novo tópico"** dispara novo fetch, trocando o conteúdo do card.
- Em caso de erro de rede/API, exibe mensagem de fallback ("Erro ao carregar") sem quebrar o layout.

## Deploy

Publicado como site estático no GitHub Pages (ver workflow de CI/CD na raiz do repositório).
