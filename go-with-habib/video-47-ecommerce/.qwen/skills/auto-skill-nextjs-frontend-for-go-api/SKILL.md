---
name: nextjs-frontend-for-go-api
description: Scaffold a Next.js 16 frontend (App Router, TypeScript) that consumes an existing Go HTTP API on localhost via the absolute backend URL — no rewrites, no proxy. Verify Go routes from main.go before scaffolding.
source: auto-skill
extracted_at: '2026-06-28T14:34:09.530Z'
---

# Next.js 16 frontend for a Go HTTP API

When the user has an existing Go HTTP server (e.g. `main.go` exposing `GET /product` and `POST /product` on `:8080`) and asks for a frontend in a sibling folder, scaffold Next.js **by hand** instead of `create-next-app`. `create-next-app` is interactive and slow; the file set is small and predictable.

## Why hand-scaffold
- `create-next-app` prompts and downloads hundreds of MB. The minimal App Router project is ~7 files.
- The user is on Windows + Next.js + npm; see `user/environment.md` memory.

## ⚠️ Direct-URL approach is the default — NOT rewrites
Earlier versions of this skill recommended Next.js **rewrites** in `next.config.ts` to hide the backend URL behind `/api/*`. The user has explicitly rejected that pattern in this project:

> "I will explore how to change main.go... You won't change any of my files in the server. You just keep them as they are: main.go and my two files, go.mod, maybe not. Keep them as they are. You won't change any of those. What you will do is just put it in the frontend UI folder."

Translated: the UI calls `http://localhost:8080/product` **directly** from the browser. `next.config.ts` stays vanilla (empty config). A Next.js API proxy route at `app/api/product/route.ts` was also tried and rejected — they want the absolute URL visible in `lib/api.ts`, not hidden behind any indirection.

**The user is responsible for adding CORS headers to `main.go` themselves.** If a CORS error appears in the browser console, surface it back to the user — do NOT "fix" it by editing `main.go`, adding rewrites, or reintroducing a proxy route.

## File set to create
```
UI/
├── package.json              next@^16 (latest patched), react@19, typescript
                                          ^ pin to ^16.x; never pin exact 16.0.0 (CVE-2025-66478). See `feedback/apply-dependency-security-patches.md`.
├── tsconfig.json             strict, paths: { "@/*": ["./*"] }
├── next.config.ts            VANILLA — const nextConfig: NextConfig = {};  (no rewrites)
├── next-env.d.ts             // standard reference comment file
├── .gitignore                node_modules, .next, .env*
├── app/
│   ├── layout.tsx            imports globals.css, sets <html>/<body>
│   ├── globals.css           minimal dark theme (slate-800/900)
│   └── page.tsx              "use client" — fetch + form
└── lib/
    └── api.ts                typed fetch helpers, API_BASE = "http://localhost:8080"
```

## Key decisions
1. **App Router** (not `pages/`). It's the Next.js 16 default.
2. **Client component** for the page (`"use client"`) because we need `useState`/`useEffect` for the fetch-and-render loop.
3. **`cache: "no-store"` on GET** so the list refreshes after each POST without manual cache busting.
4. **`API_BASE = "http://localhost:8080"`** in `lib/api.ts`. No rewrites, no proxy routes, no `/api/*` prefix.
5. **One helper per endpoint** in `lib/api.ts` — typed, throws on non-2xx, returns parsed JSON. Don't inline `fetch` calls in the page.
6. **JSON Content-Type header** must be set explicitly on POST; the Go `json.NewDecoder` requires it.

## API client pattern (`lib/api.ts`)
```ts
const API_BASE = "http://localhost:8080";

export async function getProducts(): Promise<User[]> {
  const res = await fetch(`${API_BASE}/product`, { method: "GET", cache: "no-store" });
  if (!res.ok) throw new Error(`GET /product failed: ${res.status}`);
  return res.json();
}

export async function createProduct(p: Omit<User, "id">): Promise<User> {
  const res = await fetch(`${API_BASE}/product`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(p),
  });
  if (!res.ok) throw new Error(await res.text());
  return res.json();
}
```

## ⚠️ Verify the path of every route in `main.go` before writing `lib/api.ts`
Do **not** assume both endpoints share the same path. In this project, the user's `main.go` has:
```go
mux.HandleFunc("/", getProducts)        // GET -> /
mux.HandleFunc("/product", createProduct) // POST -> /product
```
GET on `/product` would 404. Always read `main.go` and replicate each `mux.HandleFunc(...)` exactly in `lib/api.ts`. If the routes don't match the API client pattern above, fix the client — do NOT edit `main.go`.

## next.config.ts — keep it vanilla
```ts
import type { NextConfig } from "next";

const nextConfig: NextConfig = {};

export default nextConfig;
```

## Run order (tell the user)
1. Start Go: `cd <project> && go run .` (keep running on `:8080`)
2. Start UI: `cd UI && npm install && npm run dev` (defaults to `:3000`)
3. Open `http://localhost:3000`

If port 3000 is held by a stale `next dev`, see `feedback/port-conflict-resolution.md` — kill and restart on `:3000`, do not move to `:3001`.

## Common pitfalls
- **Forgetting `"use client"`** → `useState`/`useEffect` errors at build time.
- **Reintroducing rewrites or a proxy route** → user rejection. `next.config.ts` stays vanilla, `lib/api.ts` keeps `API_BASE = "http://localhost:8080"`.
- **Forgetting `Content-Type: application/json`** on POST → Go's `json.NewDecoder` returns an error and returns 400.
- **Running only one of the two servers** → `fetch failed` in the UI. Confirm Go is on `:8080` first.
- **CORS error in the browser console** → that is the user's problem to solve in `main.go`; do not fix it on the frontend.
- **Hydration warnings mentioning `data-gr-ext-installed`** → Grammarly browser extension mutates `<body>` after server render. If the user finds it noisy, add `suppressHydrationWarning` to `<body>` in `app/layout.tsx` — silences the warning without affecting functionality. Don't add it pre-emptively; only when asked.
- **Go's struct tags** (separate gotcha, found in `main.go` of this project) — Go uses backtick struct tags `` `json:"id"` ``, not string-literal tags `"json:\"id\""`. The latter compiles but is ignored by the JSON encoder.
- **Inverted method-check guards in Go handlers** (separate gotcha) — `mux.HandleFunc("GET /x", ...)` already restricts the route to GET. A guard like `if r.Method == "GET" { return 400 }` then rejects the only method that could ever reach it. Use `if r.Method != http.MethodGet { return 400 }`.
- **CVE on Next.js 16.0.0** — scaffold with `next@^16` (latest patched). If a fresh `npm install` flags a CVE, propose upgrading right away (see `feedback/apply-dependency-security-patches.md`); the user treats CVEs as something to act on immediately.