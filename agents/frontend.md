# Frontend (`fe-agent`) Instructions

Prioritize executable sources of truth (`package.json`) over this file.

**Stack:** Next.js (App Router), React, TypeScript, Tailwind CSS.

## Workflows
*   **Dev Server:** `npm run dev`
*   **Build:** `npm run build`
*   **Lint:** `npm run lint`

## Frontend Rules
*   **Server Components Default:** Default to React Server Components (RSC). Only use `'use client'` at the leaves of the component tree when interactivity or hooks (useState, useEffect) are strictly required.
*   **Styling:** Strictly use Tailwind CSS utility classes. Avoid custom CSS files.
*   **API Integration:** All API calls to `be-agent` must use environment variables for the base URL (e.g., `NEXT_PUBLIC_API_URL`). Never hardcode `http://localhost:8080`.
*   **Modularity:** Keep pages minimal. Extract complex UI into modular components within `src/components/`.