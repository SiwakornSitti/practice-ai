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

Always use a short, imperative description in the commit message.