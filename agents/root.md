# System Instructions

**Role:** You are an expert AI assistant managing a monorepo consisting of two distinct applications: a Go backend (`app/`) and a Next.js frontend (`fe-agent/`).

## Routing Rules
**CRITICAL:** Do not rely on this root file for implementation details.

*   If working on the Go backend service (`app/`), you MUST read `agents/backend.md` before taking action.
*   If working on the Next.js frontend (`fe-agent/`), you MUST read `agents/frontend.md` before taking action.

## Commit Convention
We follow [Conventional Commits](https://www.conventionalcommits.org/):
* `feat`: New feature
* `fix`: Bug fix
* `docs`: Documentation changes
* `chore`: Build/tooling updates (e.g., Makefile, Dockerfile)
* `refactor`: Code changes neither fixing a bug nor adding a feature
* `test`: Adding missing tests
* `style`: Formatting, missing semi-colons, etc.

## Global Constraints
*   Always `cd` into the appropriate directory (`app/` or `fe-agent/`) before running build tools or scripts.
*   Follow 12-Factor App principles. Ensure applications are container-ready, use environment variables for config, log to stdout/stderr, and include optimized Dockerfiles.

## Project-Scoped Skills
Custom skills specific to this repository are stored in `.agents/skills/`. 

When working on a specialized task:
1. Check `.agents/skills/` for existing custom workflows.
2. If you need to use one, load it via the `skill` tool using its name (the directory name in `.agents/skills/`).
3. If no skill exists, feel free to use `npx skills init` to create a new one inside this directory to share with the team.