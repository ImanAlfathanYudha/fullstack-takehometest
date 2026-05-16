# Fullstack Payment Dashboard

This is a fullstack application for a Payment Dashboard. It uses **Go** for the backend (Clean Architecture) and **Nuxt/Next** for the frontend.

## 🛠 Tools & Versions

```bash
go version go1.25.5 # Supports pure-Go SQLite
node v24.13.1
```

## Installation

**Backend:**
```bash
cd backend
cp env.sample .env
go mod tidy
go mod vendor
```

**Frontend:**
```bash
cd frontend
npm install
```

## Running Locally

**Backend Server:**
```bash
cd backend
# Windows (PowerShell)
$env:CGO_ENABLED="0"; go run main.go

# Mac / Linux
make run
```

**Frontend Client:**
```bash
cd frontend
npm run dev
```

## Production Build

**Backend:**
```bash
cd backend
make build
./bin/backend
```

**Frontend:**
```bash
cd frontend
npm run build
npm run start
```

## API Documentation

Once the backend is running, you can view the API specification by opening the `openapi.yaml` file in any OpenAPI/Swagger viewer (e.g., [editor.swagger.io](https://editor.swagger.io)).

## Accessing the App

Login to the dashboard by visiting:
```bash
http://localhost:3000/login
```
**Credentials:**
- Email: `cs@test.com`
- Password: `password`

---

## 📂 Project Structure
- [Backend Documentation](backend/README.md)
- [Frontend Documentation](frontend/README.md)

