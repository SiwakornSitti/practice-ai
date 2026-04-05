# Agent Instructions

This file contains repository-specific context for OpenCode sessions. When in doubt, prioritize executable sources of truth (config files, scripts) over this file.

## Core Workflows & Commands

*   **Build:** `[INSERT BUILD COMMAND]`
*   **Test:** `[INSERT TEST COMMAND]`
*   **Lint/Format:** `[INSERT LINT/FORMAT COMMAND]`
*   **Check:** `[INSERT TYPECHECK/STATIC ANALYSIS COMMAND]`

*   **Execution Order:** If order matters for verification (e.g., `lint` must precede `test`), specify it here.

## Project Structure & Architecture

*   **Entrypoints:** `[Point to main app/lib entrypoints]`
*   **Monorepo/Boundaries:** `[If applicable, define package/directory boundaries]`
*   **Generated Code:** `[Identify paths containing generated code that should not be manually edited]`

## Context & Operational Gotchas

*   **Environment:** `[Specify required setup, env vars, or specific tool versions]`
*   **Framework Quirks:** `[Identify non-standard usage or project-specific patterns]`
*   **Testing:** `[Note prerequisites for tests, fixtures, or known flaky suites]`

## References

*   **Configuration:** `[List key config files like package.json, tsconfig.json, etc.]`
*   **CI Workflows:** `[Path to CI files]`
