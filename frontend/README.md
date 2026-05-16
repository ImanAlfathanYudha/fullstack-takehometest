# Durianpay Fullstack Payment System

This repository contains a full-stack application with a Go backend and a Next.js (React) frontend.

## 📁 Repository Structure
- `/backend`: Go service handling Auth and Payments.
- `/frontend`: Next.js application for the Dashboard UI.
- `openapi.yaml`: API documentation at the root.

## 🛠️ Prerequisites
- Go 1.22+
- Node.js 20+
- `make` (optional)

## 🚀 Step-by-Step Setup

### 1. Backend Setup
1. Navigate to `/backend`.
2. Install dependencies: `go mod download`.
3. Start the server: `go run main.go` (or `make run`).
   *The backend runs on http://localhost:8080*

### 2. Frontend Setup
1. Navigate to `/frontend`.
2. Install dependencies: `npm install`.
3. Start the development server: `npm run dev`.
   *The frontend runs on http://localhost:3001*

### 3. Login Credentials
Use any valid email/password combination (e.g., `cs@test.com` / `password`) as per the backend mock (you can check it on backend/main.go)

## 📖 API Documentation
The API is defined using OpenAPI v3. You can find the specification in `openapi.yaml` at the root of this repository.

## 🧪 Commands Summary
- **Start Backend**: `cd backend && go run main.go`
- **Start Frontend**: `cd frontend && npm run dev`
- **Build Frontend**: `cd frontend && npm run build`

## Flow
1. Access this link: http://localhost:3001/
2. Go to the login page by clicking login button, or you can access this link: http://localhost:3001/login
3. If you successfully login, you'll be directed to http://localhost:3001/dashboard (the dashboard page)
