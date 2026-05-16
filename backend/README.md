# Fullstack Payment Dashboard - Backend

This is the backend service for the Durianpay Fullstack Take-home test. It follows Clean Architecture principles (Entity-Repository-Usecase) to ensure scalability and maintainability.

## 🚀 Quick Start (Windows / PowerShell)

If you are on Windows, run these commands in order:

1. **Setup Environment:**
   ```powershell
   cp env.sample .env
   ```

2. **Install Dependencies:**
   ```powershell
   go mod tidy
   go mod vendor
   ```

3. **Run the App:**
   ```powershell
   $env:CGO_ENABLED="0"
   go run main.go
   ```

## 🍎 Quick Start (Mac / Linux / Makefile)

1. **Setup Environment:**
   ```bash
   cp env.sample .env
   make dep
   ```

2. **Run the App:**
   ```bash
   make run
   ```

## 🛠 Features

- **Database:** SQLite (using `modernc.org/sqlite` pure-Go driver to avoid CGO/GCC dependency issues).
- **Auto-Seeding:** On startup, the app automatically creates the database and populates it with 3 hard-coded payment records if the table is empty.
- **Authentication:** JWT-based authentication via custom middleware.
- **OpenAPI Validation:** Requests are validated against `openapi.yaml` using middleware.
- **CORS Support:** Pre-configured for frontend local development (ports 3001).

## 🧪 Testing

To run the unit tests:
```bash
make test
# OR (Windows)
$env:CGO_ENABLED="0"; go test ./... -v
```

## 📡 API Endpoints

- **POST `/dashboard/v1/auth/login`**: Authenticate and receive a JWT token.
- **GET `/dashboard/v1/payments`**: Fetch all payments (Requires Bearer Token).
  - Supports filtering by `status` parameter. the allowed parameters are (completed, processing, failed).
