# be-agent

A Go backend service using Clean Architecture, organized by feature.

## Architecture

- `cmd/`: Application entrypoints.
- `internal/`: Private application code.
    - `{feature}/`: Feature-based packages.
        - `domain/`: Business entities and interfaces.
        - `usecase/`: Business logic.
        - `repository/`: Data access implementations.
        - `delivery/`: API/Transport layer (HTTP handlers, gRPC, etc.).
