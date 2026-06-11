# State Management & Routing

Project-specific conventions for state and navigation.

## State

- Library: TODO (e.g. Pinia for Vue / Redux Toolkit or Zustand for React).
- What goes in global state: TODO (auth/session, current user, cross-page data).
- What stays local: form state, UI toggles, anything one component owns.
- Async/data fetching pattern: TODO (e.g. composables wrapping `src/api`).

## Routing

- Library: TODO (vue-router / react-router).
- Route structure: TODO (public vs authenticated areas).
- Auth guard: TODO (where unauthenticated users are redirected).
- Code-splitting / lazy routes: TODO.

## API integration

- All calls go through `src/api`. Base URL from `VITE_API_BASE_URL`.
- Error/loading handling convention: TODO.
- Auth token handling: see `memory/project-facts.md`.
