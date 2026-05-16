# Klyvi API

Klyvi is a personalized movie and TV discovery app that learns your taste and helps you find what to watch next. Track what you've watched, plan to watch, or are currently watching, and the recommendations build from there.

This repository (`klyvi`) contains the backend REST API, built with Go. The frontend lives in a separate repo: [klyvi-web](https://github/jacobkly/klyvi).

## Why Separate Repos?

- Easier deployment (separate API and web hosting)
- Simpler CI/CD and maintenance
- Scalable for future features and contributors

## Tech Stack

- **Backend:** [Go](https://go.dev/) REST API
- **Database:** [Supabase](https://supabase.com/) (PostgreSQL)
- **Authentication:** Supabase Auth
- **Storage:** Supabase Buckets (for user-uploaded content)
- **3rd-Party API:** [TMDB](https://www.themoviedb.org/) (The Movie Database)
- **Migrations:** [Goose](https://github.com/pressly/goose) for database schema management

## Hosting

- Single VM deployment, API and frontend behind an Nginx proxy

## Project Goal

Make it easy to find what to watch next, with recommendations that reflect a user's actual taste across all movies and TV rather than a single streaming catalog.

## Milestones

- ✅ Initial TMDB client layer
- ✅ Initial database design (movies, TV series/seasons, collections, tracking, etc.)
- ⬜ Catalog ingestion: full movie and TV data fetched, normalized, and cached
- ⬜ Supabase JWT auth middleware
- ⬜ Tracking: log, rate, watchlist, and history per user
- ⬜ Recommendation engine: tiered scorers, from cold fallback through fully personalized
- ⬜ Deployment

## Getting Started

**Prerequisites:** Go 1.24+, Docker, a Supabase project, a TMDB API key.

Create a `.env.dev` file:

```bash
SERVER_PORT=8080
SERVER_TIMEOUT_READ=3s
SERVER_TIMEOUT_WRITE=5s
SERVER_TIMEOUT_IDLE=5s
SERVER_DEBUG=true

DB_HOST=...           # Supabase pooler host
DB_PORT=5432
DB_USER=...
DB_PASS=...
DB_NAME=postgres
DB_SSLMODE=require
DB_DEBUG=true

TMDB_API_KEY=...
TMDB_BASE_URL=https://api.themoviedb.org/3
```

Run with Docker (applies migrations on boot, then starts the API on `:8080`):

```bash
docker compose up --build
```

Or run locally:

```bash
go run ./cmd/migrate up
go run ./cmd/api
```

Health check:

```bash
curl http://localhost:8080/health
```

Stay tuned for updates!