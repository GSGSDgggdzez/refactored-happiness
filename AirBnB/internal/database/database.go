package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	GetUserByEmail(email string) (*User, error)

	GetUserByEmailAndPassword(email string, password string) (*User, error)

	CreateNewUser(user *User) error

	VerifyUserEmail(token string) error

	UpdatePassword(token string) error

	GetUserByVerificationToken(token string) (*User, error)

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

// these is the structure of the User database
type User struct {
	ID                int64
	Email             string
	Password          string
	FirstName         string
	LastName          string
	Country           string
	Phone             string
	VerifiedEmail     bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Profile_Url       string
	VerificationToken string
}

// I want to get the user data
func (s *service) GetUserByEmail(email string) (*User, error) {
	var user User
	query := `SELECT email
              FROM users WHERE email = ?`
	err := s.db.QueryRow(query, email).Scan(
		&user.Email,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *service) GetUserByEmailAndPassword(email string, password string) (*User, error) {
	var user User
	query := `SELECT  email, password 
              FROM users WHERE email = ?`
	err := s.db.QueryRow(query, email).Scan(
		&user.Email,
		&user.Password,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil
	}

	return &user, nil
}

func (s *service) CreateNewUser(user *User) error {
	query := `INSERT INTO users (email, password, first_name, last_name, country, phone, verifie_email, created_at, updated_at, Profile_Url, verificationToken) 
              VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), ?, ?)`

	result, err := s.db.Exec(query,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Country,
		user.Phone,
		user.VerifiedEmail,
		user.Profile_Url,
		user.VerificationToken,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

func (s *service) VerifyUserEmail(token string) error {
	query := `UPDATE users SET verifie_email = true WHERE verificationToken = ?`
	_, err := s.db.Exec(query, token)
	return err
}

func (s *service) UpdatePassword(token string) error {
	query := `UPDATE users SET verifie_email = true WHERE verificationToken = ?`
	_, err := s.db.Exec(query, token)
	return err
}

func (s *service) GetUserByVerificationToken(token string) (*User, error) {
	var user User

	query := `SELECT id, email, first_name, last_name, country, phone, profile_url
              FROM users 
              WHERE verification_token = ?`

	err := s.db.QueryRow(query, token).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Country,
		&user.Phone,
		&user.Profile_Url,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invalid verification token")
	}

	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	return &user, nil
}

var (
	dbname     = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}
