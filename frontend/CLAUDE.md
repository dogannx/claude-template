# CLAUDE.md — Frontend (Vue)

Frontend-specific context; loaded on top of the root `CLAUDE.md`.

## Overview

Vue SPA built with Vite, managed with **yarn**. Talks to the backend API over HTTP/JSON.

## Layout

```
frontend/
├── src/
│   ├── views/        # pages (route targets)
│   ├── components/   # reusable UI components
│   ├── stores/       # shared state
│   ├── api/          # backend client / service layer
│   ├── router/       # route definitions
│   └── main.ts       # app entry
├── package.json
└── Dockerfile
```

## Conventions (see also `docs/coding-standards.md`)

- Composition API with `<script setup>`.
- Components: PascalCase, one per file.
- All HTTP calls go through `src/api` — never call fetch/axios directly in components.
- Shared state in `stores/`; local state otherwise.
- Read the API base URL from env (`VITE_API_BASE_URL`).

## Commands

```bash
yarn install
yarn dev        # dev server (http://localhost:5173)
yarn build      # production build
yarn test       # tests
yarn lint && yarn format
```

## When adding a feature

1. Add the API call(s) in `src/api`.
2. Add store state if shared.
3. Build component(s) in `src/components`, page in `src/views`.
4. Wire the route in `src/router`.
5. Write tests; keep components small and typed.
