# Backend (`be-agent`) Instructions

Prioritize executable sources of truth (`Makefile`, `go.mod`) over this file.

**Stack:** Go, standard `net/http`.

## Workflows
*   **Dev Server:** `make run`
*   **Build:** `make build`
*   **Test:** `make test`
*   **Dependencies:** `make tidy`

## Architecture
*   **Clean Architecture, packaged by feature.**
*   **Structure:** `internal/{feature}/`
    *   `domain/`: Entities and interfaces. Zero external dependencies.
    *   `usecase/`: Business logic.
    *   `repository/`: Data access.
    *   `delivery/`: Transport (HTTP handlers).

## Backend Rules
*   **Unit Tests Mandatory:** Write unit tests for all `usecase`, `delivery`, and `repository` code.
*   **Routing:** Stick to standard library `net/http`. No external routers unless permitted.
*   **Security Focus:** Implement input validation, JWT middleware, secure headers, and SQL-injection-safe queries.
*   **Graceful Shutdown:** Catch `SIGINT`/`SIGTERM` to ensure in-flight requests close properly.