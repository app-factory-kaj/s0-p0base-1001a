# greeter — PRD

## Problem Statement

Developers building demos, integrations, or test harnesses often need a tiny, reliable HTTP endpoint to call that returns a predictable, personalized response — without standing up or configuring a real business service. Today they either hardcode a fake response or reuse a heavier internal service that was never meant for this purpose. S0 marker s0-p0base-1001a.

## Solution

greeter is a small, standalone Go HTTP service exposing a single endpoint, `GET /hello`, that returns a JSON greeting personalized by an optional `name` query parameter. It is open to any caller, has no moving parts beyond the one endpoint, and always returns a sensible response.

## Actors

- **API Caller** — any developer or external client that calls the greeter service's HTTP endpoint directly. No sign-in or identity is involved; the actor is simply whoever issues the HTTP request.

## User Stories

1. As an API Caller, I want to call `GET /hello?name=X` and receive a JSON greeting that includes the name I supplied, so that I can verify the service is reachable and working with my own input.
2. As an API Caller, I want `GET /hello` to still return a usable JSON greeting when I omit the `name` parameter (or send it empty), so that the endpoint never fails just because I left it out.

## Product Decisions

- Authentication: none — `/hello` is a fully open endpoint, with no sign-in or API-level access control.
- Response format: JSON, containing the greeting message.
- Default behavior: when `name` is missing or empty, the service returns a generic greeting (e.g. "Hello, World!") rather than an error.
- Audience: the service is reachable by any developer or external client — it is not restricted to internal/org callers.

## Out of Scope

- Any endpoint other than `GET /hello`.
- Persistence, logging of requests, rate limiting, or analytics.
- Internationalization/localization of the greeting text.
- Authentication, authorization, or per-caller quotas.

## Open Questions

None at this time.