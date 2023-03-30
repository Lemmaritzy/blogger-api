package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var (
	DB  *pgxpool.Pool
	Ctx context.Context
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "9461"
	dbname   = "blogger_db"
)

func init() {
	fmt.Println("postgres initializing")
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s pool_max_conns=%d", host, port, user, password, dbname, 5)
	DB, err = pgxpool.Connect(context.Background(), conn)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connected successfully")
	Ctx = context.Background()
}
