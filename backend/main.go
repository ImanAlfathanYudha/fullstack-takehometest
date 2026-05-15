package main

import (
	"database/sql"
	"log"
	"time"

	"backend/internal/api"
	"backend/internal/config"
	authHandler "backend/internal/module/auth/handler"
	authRepository "backend/internal/module/auth/repository"
	authUsecase "backend/internal/module/auth/usecase"
	paymentHandler "backend/internal/module/payment/handler"
	paymentRepository "backend/internal/module/payment/repository"
	paymentUsecase "backend/internal/module/payment/usecase"
	srv "backend/internal/service/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

func main() {
	_ = godotenv.Load()
	db, err := sql.Open("sqlite", "dashboard.db?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := initDB(db); err != nil {
		log.Fatal(err)
	}
	JwtExpiredDuration, err := time.ParseDuration(config.JwtExpired)
	if err != nil {
		panic(err)
	}
	// ── Auth module (existing) ──
	userRepo := authRepository.NewUserRepo(db)
	authUC := authUsecase.NewAuthUsecase(userRepo, config.JwtSecret, JwtExpiredDuration)
	authH := authHandler.NewAuthHandler(authUC)
	// ── Payment module (NEW) ──
	paymentRepo := paymentRepository.NewPaymentRepo(db)
	paymentUC := paymentUsecase.NewPaymentUsecase(paymentRepo)
	paymentH := paymentHandler.NewPaymentHandler(paymentUC)
	apiHandler := &api.APIHandler{
		Auth:    authH,
		Payment: paymentH, // ← NEW
	}
	server := srv.NewServer(apiHandler, config.OpenapiYamlLocation, config.JwtSecret) // ← pass JwtSecret for middleware
	addr := config.HttpAddress
	log.Printf("starting server on %s", addr)
	server.Start(addr)
}

func initDB(db *sql.DB) error {
	// create tables if not exists
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  email TEXT NOT NULL UNIQUE,
		  password_hash TEXT NOT NULL,
		  role TEXT NOT NULL
		);`,
		//query for payment tables
		`CREATE TABLE IF NOT EXISTS payments (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  merchant TEXT NOT NULL,
		  status TEXT NOT NULL,
		  amount REAL NOT NULL,
		  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}
	// seed admin user if not exists
	var count int
	row := db.QueryRow("SELECT COUNT(1) FROM users")
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?)", "cs@test.com", string(hash), "cs"); err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?)", "operation@test.com", string(hash), "operation"); err != nil {
			return err
		}
	}

	row = db.QueryRow("SELECT COUNT(1) FROM payments")
	if err := row.Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		payments := []struct {
			merchant string
			status   string
			amount   float64
		}{
			{"Tokopedia", "completed", 150000},
			{"Shopee", "processing", 75000.50},
			{"Grab", "failed", 200000},
		}

		for _, p := range payments {
			_, err := db.Exec(
				"INSERT INTO payments(merchant, status, amount) VALUES (?, ?, ?)",
				p.merchant,
				p.status,
				p.amount,
			)
			if err != nil {
				return err
			}
		}
	}

	const dbLifetime = time.Minute * 5
	db.SetConnMaxLifetime(dbLifetime)

	return nil
}
