package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Postgresql struct {
	postgresql *pgx.Conn
	PoolName   string
}

// postgresql://user:password@netloc:port/dbname
func New(postgresEndpoint string) (*Postgresql, error) {
	conn, err := pgx.Connect(context.Background(), postgresEndpoint)

	if err != nil {
		return nil, err
	}
	return &Postgresql{
		postgresql: conn,
	}, nil
}

func Close(db *Postgresql) {
	db.postgresql.Close(context.Background())
}
