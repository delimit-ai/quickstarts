# TaskForge — Spring Boot

Minimal Spring Boot API with Delimit governance.

## Setup

```bash
./mvnw spring-boot:run
# or: mvn spring-boot:run
```

Requires Java 17+.

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
