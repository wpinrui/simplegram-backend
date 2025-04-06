import (
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	ID             int
	Username       string
	HashedPassword string
	CreatedAt      string
}

func main() {
	connStr := "user=postgres dbname=simplegram sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func createUsersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		hashed_password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, user User) error {
	query := `INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING id;`
	err := db.QueryRow(query, user.Username, user.HashedPassword).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
