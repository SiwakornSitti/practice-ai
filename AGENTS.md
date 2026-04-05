# Agent Instructions

This file contains repository-specific context for OpenCode sessions working on the `practice-ai` repository. Prioritize executable sources of truth (like the Makefile or go.mod) over this file.

## Core Workflows & Commands

All core commands are managed via the `Makefile` inside the `be-agent` directory. Always `cd be-agent` before running these.

*   **Run Dev Server:** `make run`
*   **Build:** `make build` (outputs to `be-agent/bin/`)
*   **Test:** `make test`
*   **Tidy Modules:** `make tidy`
*   **Lint:** `make lint` (Requires `golangci-lint`)

## Project Structure & Architecture

The application is a Go backend service built using **Clean Architecture** and is organized by **feature** (also known as packaged by feature).

*   **Entrypoint:** `be-agent/cmd/api/main.go`
*   **Architecture Pattern:** Clean Architecture. When creating or modifying features, you must adhere to this layer structure inside `be-agent/internal/{feature}/`:
    *   `domain/`: Core business entities, structs, and repository/usecase interfaces. No external dependencies.
    *   `usecase/`: Business logic implementations. Depends only on `domain/`.
    *   `repository/`: Data access implementations (e.g., SQL, NoSQL). Implements interfaces defined in `domain/`.
    *   `delivery/`: Transport layer (e.g., HTTP handlers, gRPC, CLI). Depends on `usecase/`.

## Context & Operational Gotchas

*   **Before Starting New Features:** **Agents MUST use the `question` tool to ask the user which web framework they wish to use** before building new Delivery layers, unless it is already explicitly clear from existing code or prompts.
*   **Routing:** Currently using standard library `net/http`. Do not introduce external routers (like Gin, Chi, or Fiber) unless explicitly requested by the user.
*   **Dependencies:** Run `make tidy` after adding any new dependencies or creating new files that import external packages.
*   **Security Context:** The user has indicated a priority on making the application secure. Proactively consider input validation, safe database queries (avoid SQL injection), secure headers, and authentication middleware when building out the Delivery and Repository layers.

## References

*   **Go Modules:** `be-agent/go.mod`
*   **Task Runner:** `be-agent/Makefile`
