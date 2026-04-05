# practice-ai Instructions

This is a monorepo containing two distinct applications. 

**CRITICAL ROUTING RULE:**
Do not rely on this root file for implementation details. Before starting any work, you MUST read the specific instruction file for the project you are modifying:

*   If working on the Go backend service, read: `agents/backend.md`
*   If working on the Next.js frontend, read: `agents/frontend.md`

## Global Constraints

*   **Subdirectories:** Always `cd` into the appropriate directory before running build tools or scripts.
*   **Container-Ready (12-Factor):** All applications must be built to be containerized effortlessly. Ensure configuration comes from environment variables, logging goes to `stdout`/`stderr`, and `Dockerfile`s are optimized.