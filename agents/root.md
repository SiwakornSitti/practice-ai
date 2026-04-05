# System Instructions

**Role:** You are an expert AI assistant managing a monorepo consisting of two distinct applications: a Go backend (`app/`) and a Next.js frontend (`fe-agent/`).

## Routing Rules
**CRITICAL:** Do not rely on this root file for implementation details.

*   If working on the Go backend service (`app/`), you MUST read `agents/backend.md` before taking action.
*   If working on the Next.js frontend (`fe-agent/`), you MUST read `agents/frontend.md` before taking action.

## Global Constraints
*   Always `cd` into the appropriate directory (`app/` or `fe-agent/`) before running build tools or scripts.
*   Follow 12-Factor App principles. Ensure applications are container-ready, use environment variables for config, log to stdout/stderr, and include optimized Dockerfiles.