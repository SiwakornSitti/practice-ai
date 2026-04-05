# Agent Instructions

Context for the `practice-ai` repository. Prioritize executable sources of truth over this file.

## Global Constraints & Gotchas

*   **Subdirectories:** The repository is split into `be-agent` (backend) and `fe-agent` (frontend). Always `cd` into the appropriate directory before running commands.
*   **Container-Ready (12-Factor):** All applications must be built to be containerized effortlessly. This means:
    1.  **Configuration:** strictly use environment variables (`os.Getenv` or `process.env`).
    2.  **Logging:** Log entirely to `stdout`/`stderr`.
    3.  **Health Checks:** Provide readiness probes (e.g., `GET /health`).
    4.  **Dockerfiles:** Proactively create multi-stage `Dockerfile`s optimized for production.

---

## Backend (`be-agent`)

**Stack:** Go, standard `net/http`.

### Workflows
*   **Dev Server:** `make run`
*   **Build:** `make build`
*   **Test:** `make test`
*   **Dependencies:** `make tidy`

### Architecture
*   **Clean Architecture, packaged by feature.**
*   **Structure:** `be-agent/internal/{feature}/`
    *   `domain/`: Entities and interfaces. Zero external dependencies.
    *   `usecase/`: Business logic.
    *   `repository/`: Data access.
    *   `delivery/`: Transport (HTTP handlers).

### Backend Rules
*   **Unit Tests Mandatory:** Write unit tests for all `usecase`, `delivery`, and `repository` code.
*   **Routing:** Stick to standard library `net/http`. No external routers unless permitted.
*   **Security Focus:** Implement input validation, JWT middleware, secure headers, and SQL-injection-safe queries.
*   **Graceful Shutdown:** Catch `SIGINT`/`SIGTERM` to ensure in-flight requests close properly.

---

## Frontend (`fe-agent`)

**Stack:** Next.js (App Router), React, TypeScript, Tailwind CSS.

### Workflows
*   **Dev Server:** `npm run dev`
*   **Build:** `npm run build`
*   **Lint:** `npm run lint`

### Frontend Rules
*   **Server Components Default:** Default to React Server Components (RSC). Only use `'use client'` at the leaves of the component tree when interactivity or hooks (useState, useEffect) are strictly required.
*   **Styling:** Strictly use Tailwind CSS utility classes. Avoid custom CSS files.
*   **API Integration:** All API calls to `be-agent` must use environment variables for the base URL (e.g., `NEXT_PUBLIC_API_URL`). Never hardcode `http://localhost:8080`.
*   **Modularity:** Keep pages minimal. Extract complex UI into modular components within `src/components/`.
