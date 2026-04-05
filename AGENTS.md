# Agent Instructions

Context for the `practice-ai` repository. Prioritize executable sources of truth (`be-agent/Makefile`, `be-agent/go.mod`) over this file.

## Workflows & Commands

**CRITICAL:** The Go project is in a subdirectory. You must `cd be-agent` before running any commands.

*   **Dev Server:** `make run`
*   **Build:** `make build` (outputs to `be-agent/bin/`)
*   **Test:** `make test`
*   **Lint:** `make lint` (Requires `golangci-lint`)
*   **Dependencies:** Run `make tidy` after importing new packages.

## Architecture

Go backend using **Clean Architecture, packaged by feature**.

*   **Entrypoint:** `be-agent/cmd/api/main.go`
*   **Structure:** `be-agent/internal/{feature}/`
    *   `domain/`: Entities and interfaces. Zero external dependencies.
    *   `usecase/`: Business logic. Depends only on `domain/`.
    *   `repository/`: Data access. Implements `domain/` interfaces.
    *   `delivery/`: Transport (HTTP handlers, etc). Depends on `usecase/`.

## Agent Constraints & Gotchas

*   **Unit Tests Mandatory:** Proactively write unit tests (using the standard `testing` package) alongside all new `usecase`, `delivery`, and `repository` code. Do not skip testing.
*   **Ask Before Delivery:** Use the `question` tool to ask the user which web framework to use *before* building new Delivery layers, unless explicitly specified.
*   **Routing:** Stick to standard library `net/http` for existing features. Do not introduce Gin, Chi, Fiber, etc., without explicit permission.
*   **Security Focus:** Prioritize security. Proactively implement input validation, JWT middleware, secure headers, and SQL-injection-safe queries.
*   **Graceful Shutdown:** The `main` entrypoint must implement a graceful shutdown mechanism catching OS signals (`SIGINT`, `SIGTERM`) to ensure in-flight requests and connections are closed properly before exiting.
