# Project Improvement Ideas

This repo is already a clean minimal starter for Go + `templ` + HTMX + Bun/Tailwind. The biggest opportunity is not adding more demo pages, but making the starter more reusable, testable, and production-safe.

## Highest-Value Improvements

### 1. Add a real application skeleton

Right now most behavior is registered through `init()` functions and package-level globals in `handle/*.go`. That keeps the demo small, but it makes the project harder to extend once it grows past examples.

Suggested change:
- Introduce an `App` or `Server` struct that owns dependencies.
- Move route registration into an explicit `RegisterRoutes(mux, deps)` path.
- Replace globals like `embeddedResources`, `setupFuncs`, and `counterValue` with fields on that struct.

Why this matters:
- Easier testing.
- Clear dependency boundaries.
- Better foundation for adding config, services, storage, auth, and middleware later.

Relevant files:
- `handle/handle.go`
- `handle/home.go`
- `handle/alert.go`
- `handle/modal.go`
- `handle/progress.go`
- `handle/themeController.go`

### 2. Add test coverage for handlers and rendering

There are currently no Go tests. For a starter project, this is one of the most important missing pieces because new users will copy the structure you establish here.

Suggested change:
- Add `httptest` coverage for each route.
- Test status codes, method restrictions, and representative HTML output.
- Add tests for static asset serving and `ETag` behavior.
- Add at least one regression test for the HTMX counter flow.

Why this matters:
- Protects the starter from silent breakage.
- Gives users examples of how to test `templ`-driven handlers.
- Makes refactoring the route setup much safer.

Relevant files:
- `handle/*.go`

### 3. Harden the HTTP server for production

`cmd/serve.go` uses `http.ListenAndServe` directly with no server timeouts or graceful shutdown.

Suggested change:
- Use an `http.Server` with `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, and `IdleTimeout`.
- Add graceful shutdown on `SIGINT`/`SIGTERM`.
- Centralize config like port and environment.

Why this matters:
- Prevents slow-client and hanging-connection problems.
- Makes the project deployable without immediate rewrites.
- Gives the starter a more professional baseline.

Relevant files:
- `cmd/serve.go`

### 4. Improve frontend lifecycle behavior

A few client-side utilities work for the demos but will become fragile as the app grows.

Observed issues:
- `web/dialog.ts` adds new `click` and `keydown` listeners each time `dialogEventHandler` runs.
- Theme initialization is driven through a script call in the head plus a nested `DOMContentLoaded` listener in `web/theme.ts`.
- The app relies on global functions on `window`, which is fine for a demo but not ideal as a starter pattern.

Suggested change:
- Use delegated or one-time event wiring.
- Initialize theme immediately without extra listener nesting to reduce flash-of-wrong-theme behavior.
- Move toward module-driven bootstrapping for HTMX hooks and UI behavior.

Why this matters:
- Avoids event listener leaks.
- Produces a cleaner pattern for future components.
- Makes the starter feel intentional rather than demo-only.

Relevant files:
- `web/index.ts`
- `web/dialog.ts`
- `web/theme.ts`
- `ui/core.templ`

## Strong Next Improvements

### 5. Create a proper example data/service layer

The counter currently uses a package-level mutex and integer:
- `counterMutex`
- `counterValue`

That is acceptable for a demo, but not for a reusable starter.

Suggested change:
- Wrap the counter in a small service interface.
- Provide an in-memory implementation now.
- Make the handler depend on that interface.

Why this matters:
- Demonstrates the pattern users should follow for app state.
- Makes it easy to swap in database-backed or per-session state later.

Relevant files:
- `handle/home.go`

### 6. Add configuration and environment management

The app currently reads only `PORT` and otherwise relies on hardcoded defaults.

Suggested change:
- Add a config package or small typed config struct.
- Support environment selection such as `APP_ENV=development|production`.
- Document all env vars in the README.

Good candidates:
- `PORT`
- `APP_ENV`
- `LOG_FORMAT`
- `TRUST_PROXY`

Why this matters:
- Makes the starter usable across local dev, Docker, and deployment targets.
- Prevents configuration from spreading across the codebase.

### 7. Add CI, linting, and formatting enforcement

This repo would benefit from a baseline quality pipeline.

Suggested change:
- Add CI to run:
  - `go test ./...`
  - `templ generate ./ui/...`
  - `bun run js`
  - `bun run tw`
- Add linting for Go and TypeScript.
- Fail CI if generated assets or templates are out of date.

Why this matters:
- Keeps the starter healthy.
- Prevents accidental drift between source and generated output.

### 8. Clean up and tighten the README

The README is useful but still reads like an early project note.

Suggested change:
- Fix typos like `Prerequesites`.
- Clarify the difference between port `4000` and `4005`.
- Document the architecture briefly: Go handlers, `templ`, HTMX, Bun build.
- Add a section explaining how assets are embedded into the binary.
- Add a "how to add a page/component/route" section.

Why this matters:
- A starter project is judged heavily by its onboarding quality.

Relevant files:
- `readme.md`
- `Taskfile.yml`
- `Dockerfile.dev`
- `docker-compose.yml`

## Product and UX Improvements

### 9. Make the example pages more representative of real SSR + HTMX patterns

The current demos are useful, but they are mostly isolated widgets. A stronger starter would include one realistic feature slice.

Suggested addition:
- Add a small CRUD-style example such as notes, todos, or contacts.
- Include:
  - form validation
  - optimistic or partial HTMX updates
  - loading state
  - empty state
  - server-side error rendering

Why this matters:
- Shows how the stack handles real user flows.
- Makes the starter more compelling than a set of disconnected examples.

### 10. Improve accessibility across the demo UI

There are some good pieces already, but accessibility is not clearly a first-class part of the starter yet.

Suggested change:
- Audit focus states, keyboard navigation, and ARIA labeling.
- Ensure toast notifications are announced properly.
- Confirm modal focus management and focus return behavior.
- Add skip navigation and stronger semantic page landmarks.

Why this matters:
- Starters tend to be copied directly into production code.
- Accessibility should be part of the baseline, not a follow-up task.

Relevant files:
- `ui/core.templ`
- `ui/modal.templ`
- `ui/alert.templ`
- `ui/home.templ`

## Operational Improvements

### 11. Add a production Dockerfile

`Dockerfile.dev` is clearly for local development. The repo would benefit from a separate multi-stage production image.

Suggested change:
- Build JS/CSS assets.
- Build the Go binary.
- Copy only the binary into a small runtime image.

Why this matters:
- Smaller image size.
- Better deploy story.
- Cleaner distinction between development and production workflows.

### 12. Revisit asset caching strategy

Assets are served through the root handler fallback and use ETags with `Cache-Control: max-age=0`.

Suggested change:
- Route static assets explicitly under predictable prefixes.
- Fingerprint generated asset names where possible.
- Use longer-lived cache headers for immutable assets.

Why this matters:
- Better browser caching.
- Clearer routing model.
- More realistic production defaults.

Relevant files:
- `handle/home.go`
- `handle/handle.go`
- `package.json`

## Small, High-Leverage Fixes

These are quick wins:

- Replace the blocking `time.Sleep(4 * time.Second)` in `handle/progress.go` with a cancellable example job or a shorter simulated task.
- Add structured logging or at least a log mode without ANSI colors for non-TTY environments.
- Move repeated route-method branching into helpers so unsupported methods consistently return `405`.
- Consider a custom 404 page instead of the current bundled-file fallback behavior.
- Add `.gitignore` coverage for generated binaries like `starter` if not already present.

## Suggested Implementation Order

If you want to improve this repo incrementally, I would do it in this order:

1. Add tests.
2. Refactor route registration away from globals and `init()`.
3. Harden the HTTP server and config handling.
4. Clean up frontend event wiring and theme bootstrapping.
5. Add one realistic CRUD example.
6. Add CI and production Docker support.

## Best Single Feature To Build Next

If the goal is to make this starter feel substantially more complete, the best next feature is:

**A small server-rendered CRUD module with validation, partial HTMX updates, tests, and in-memory storage.**

That one addition would demonstrate:
- page composition
- forms
- server-side validation
- reusable components
- route organization
- testing patterns
- real HTMX interaction design

It would do more for the project than adding more isolated demo widgets.
