# Frontend Architecture

How the frontend is structured and the decisions behind it. Vue is the default; if a
project uses React instead, mirror this doc with the React equivalents.

## Stack

- Framework: Vue (Composition API, `<script setup>`) — default in this template.
- Build: Vite. Package manager: yarn.

### Using React instead

The scaffold ships Vue. To use React for a project, replace the scaffold with
`yarn create vite . --template react-ts`, then keep the same structural rules below
(`src/api` as the only HTTP layer, `views/` + `components/`, a store lib, a router).
The principles are framework-agnostic; only the libraries change
(Pinia→Redux/Zustand, vue-router→react-router). Record the choice as an ADR.

## Structure

```
src/
├── views/        # route-level pages
├── components/   # reusable UI
├── stores/       # shared state
├── api/          # backend client/service layer (the ONLY place HTTP happens)
├── router/       # routes
└── main.ts
```

## Principles

- Components stay presentational; data fetching goes through `src/api`.
- Type everything; keep components small and focused.
- Shared state in stores; local state otherwise (don't over-globalize).

## Decisions

- Record significant frontend choices as ADRs in `docs/architecture/decisions/`
  (e.g. "chose Pinia over Vuex", "Vue vs React for this project").
