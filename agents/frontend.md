# Frontend (`fe-agent`) Instructions

Prioritize executable sources of truth (`package.json`) over this file.

**Stack:** Next.js (App Router), React, TypeScript, Tailwind CSS.

## Workflows
*   **Dev Server:** `npm run dev`
*   **Build:** `npm run build`
*   **Lint:** `npm run lint`

## 🔺 Vercel React & Next.js Best Practices (Integrated Skill)

### 1. Architecture & Components
*   **Server Components by Default:** Always default to React Server Components (RSC). 
*   **Client Components:** Only use `'use client'` at the absolute leaves of the component tree when interactivity (e.g., `useState`, `useEffect`, `onClick`) is strictly required.
*   **Modularity:** Keep `page.tsx` minimal. Extract complex UI into modular, reusable components within `src/components/`.

### 2. Data Fetching & Mutations
*   **Async Server Components:** Fetch data directly in Server Components using `async`/`await` and the native `fetch` API.
*   **Server Actions:** Use Server Actions (`'use server'`) inside `src/actions/` for data mutations and form submissions instead of creating traditional API routes.
*   **Suspense & Streaming:** Wrap slow asynchronous components in `<Suspense>` boundaries with appropriate skeleton fallbacks to enable UI streaming.

### 3. Optimization & Core Web Vitals
*   **Images:** Strictly use `next/image` (`<Image />`) for all images to ensure automatic optimization and prevent Cumulative Layout Shift (CLS).
*   **Navigation:** Strictly use `next/link` (`<Link />`) for internal routing.
*   **Fonts:** Use `next/font` to optimize web fonts and avoid layout shifts.

### 4. State & Styling
*   **URL as State:** Prefer using URL search parameters (`useSearchParams` or tools like `nuqs`) for shareable/filterable state rather than local `useState`.
*   **Styling:** Strictly use Tailwind CSS utility classes. Avoid creating custom `.css` files. Use `clsx` and `tailwind-merge` (via a `cn()` utility) for conditional class joining.

### 5. API Integration
*   **Environment Variables:** All API calls to `be-agent` must use environment variables for the base URL (e.g., `NEXT_PUBLIC_API_URL`). Never hardcode `http://localhost:8080`.