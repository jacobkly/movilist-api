# MoviList API

MoviList is a simple web app to track movies and TV shows you've watched, plan to watch, or are currently watching. Inspired by AniList, but for film and TV.

This repository (`movilist-api`) contains the backend REST API, built with Go. The frontend lives in a separate repo: [`movilist-web`](https://github.com/you/movilist-web).

## Why Separate Repos?

- Easier deployment (Heroku for API, Vercel for web)
- Simpler CI/CD and maintenance
- Scalable for future features and contributors

## Tech Stack

- **Backend:** [Golang](https://go.dev/) REST API  
- **Database:** [Supabase](https://supabase.com/) (PostgreSQL)  
- **Authentication:** Supabase Auth  
- **Storage:** Supabase Buckets (for user-uploaded content)  
- **3rd-Party API:** [TMDB](https://www.themoviedb.org/?language=en-US) (The Movie Database) API 
- **Migrations:** [Goose](https://github.com/pressly/goose) for database schema management  


## Hosting Plans

- **Backend (Development):** Local self-tunneling (e.g., zrok.io)
- **Backend (Production):** Heroku

## Project Goal

Get a hosted version online as soon as possible. MoviList aims to make it easy to organize and share your movie and show lists.

---


### Smaller Goals / Milestones

- **Build Initial TMDb Client Layer** ✅
  - ~~Return data directly from TMDb to the client.~~
- **Design Initial Database** ✅
  - ~~Tables include: movies, TV series/seasons, recommendations, collections, tracking, etc.~~
- **Develop DB Layer**
  - Implement the service layer of the API for inserting and querying the database.
- **Implement API Authentication Middleware**
  - Two methods of entry: Bearer Token (normal use) and API Key (admin use)
- **Refine Database Design**
  - Create space-efficient tables (fix suboptimal column types from the initial design).
  - Add activity-tracking tables.
  - Include additional smaller tables to reduce direct TMDb API calls (e.g., genres, production info).
- **Build Redis Caching Layer**
  - Adjust below caching strategy as needed with further research.
  - Initially cache trending/upcoming/popular/top-rated items directly from TMDb.
  - Cache individual movie/TV data selectively based on popularity.
- **Deployment**
  - Host the API on Heroku.

---

Stay tuned for updates!