# Delimit Quickstarts

[![npm](https://img.shields.io/npm/v/delimit-cli)](https://www.npmjs.com/package/delimit-cli)
[![GitHub Action](https://img.shields.io/badge/GitHub%20Action-v1-blue)](https://github.com/marketplace/actions/delimit-api-governance)
[![Docs](https://img.shields.io/badge/docs-delimit.ai-green)](https://delimit.ai/docs)
[![API Governance](https://delimit-ai.github.io/badge/pass.svg)](https://github.com/marketplace/actions/delimit-api-governance)

Pick your framework, clone the directory, push. Governance is automatic.

Each quickstart is a minimal **TaskForge API** with the same five endpoints, a valid OpenAPI 3.0 spec, and a ready-to-go Delimit GitHub Action workflow. Open a PR that touches your API and Delimit will catch breaking changes before they ship.

## Frameworks

| Framework | Language | Directory | Router |
|-----------|----------|-----------|--------|
| [Django REST Framework](django/) | Python | `django/` | DRF ViewSets |
| [FastAPI](fastapi/) | Python | `fastapi/` | FastAPI |
| [Spring Boot](spring-boot/) | Java | `spring-boot/` | Spring MVC |
| [Ruby on Rails](rails/) | Ruby | `rails/` | Rails API |
| [Go + Gin](go/) | Go | `go/` | Gin |
| [Express](express/) | Node.js | `express/` | Express |
| [NestJS](nestjs/) | TypeScript | `nestjs/` | NestJS |

## Quick start

```bash
# 1. Copy any framework directory into a new repo
cp -r django/ my-api && cd my-api

# 2. Initialize git and push
git init && git add -A && git commit -m "Initial API"
gh repo create my-org/my-api --public --push --source .

# 3. Open a PR that changes api/openapi.yaml — Delimit runs automatically
```

## What each directory contains

```
<framework>/
  api/openapi.yaml              # OpenAPI 3.0 spec (TaskForge API)
  delimit.yml                   # Delimit policy config (default preset)
  .github/workflows/
    api-governance.yml          # GitHub Action workflow
  <framework source files>      # Minimal working API
  README.md                     # Framework-specific setup
```

## TaskForge API endpoints

Every quickstart implements the same five endpoints:

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/tasks` | List all tasks |
| `POST` | `/tasks` | Create a task |
| `GET` | `/tasks/{id}` | Get a task by ID |
| `DELETE` | `/tasks/{id}` | Delete a task |

## Links

- [Delimit CLI](https://www.npmjs.com/package/delimit-cli) -- `npm i -g delimit-cli`
- [Delimit GitHub Action](https://github.com/marketplace/actions/delimit-api-governance)
- [Documentation](https://delimit.ai/docs)
- [GitHub](https://github.com/delimit-ai)
