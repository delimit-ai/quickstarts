# TaskForge — Ruby on Rails

Minimal Rails API with Delimit governance.

## Setup

```bash
bundle install
rails server
```

Requires Ruby 3.1+.

## Endpoints

```
GET    /health        — Health check
GET    /tasks         — List tasks
POST   /tasks         — Create task (body: {"title": "..."})
GET    /tasks/{id}    — Get task
DELETE /tasks/{id}    — Delete task
```

## Governance

The `delimit.yml` and `.github/workflows/api-governance.yml` are pre-configured. Push this directory to a GitHub repo, open a PR that modifies `api/openapi.yaml`, and Delimit will validate the change automatically.
