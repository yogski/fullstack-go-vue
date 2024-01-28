# Fullstack App Made with Go + Vue

## What Is This? 
This is a simple CRUD app. Under the hood, it runs on Go as backend, Vue as frontend, and PostgreSQL as database.

## Prerequisites
- Go 1.18 or above
- Vue 3+
- PostgreSQL 9 or above
- NPM 8 or above

## Quick Start Guide
- General 
  1. Clone the repository from `https://github.com/yogski/fullstack-go-vue`.
  2. Run `cd /fullstack-go-vue`.
  3. The repository contains 2 modules: `backend` and `frontend`. For each module, follow the setup below.

- Backend
  1. Run `cd/backend`.
  2. To set up `env` file, run `make env`. If `make` command is not available (like in Windows command), simply copy `.env.sample` then paste in same directory, then rename it to `.env`
  3. Fill required values in `.env`. It is recommended to leave `APP_PORT` as 5000 for development purpose.
  4. For first time setup, run `go run main.go --migration`. The `--migration` flag tells the app to run migration setup only (e.g create necessary PostgreSQL tables and insert seed data from CSV files in `/backend/connections`).
  5. Run `go run main.go` to start backend server.

- Frontend
  1. Run `cd /frontend`
  2. Run `npm install`
  3. Run `npm run dev`
  4. Now access the web app from browser and enjoy!

## Todo
Still many things to do. This is MVP with plenty of room for improvements
- Use proper ORM using Gorm
- Refactor seeding process