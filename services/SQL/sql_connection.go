package services

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConnection struct {
	ctx context.Context
	db  *sql.DB
}

func NewSQLConnection(ctx context.Context) SQLConnection {
	address, exists := os.LookupEnv("SQL_ADDRESS")

	if exists == false {
		panic("SQL_ADDRESS not found in .env")
	}

	db, err := sql.Open("mysql", address)

	if err != nil {
		panic(err)
	}

	return SQLConnection{
		ctx,
		db,
	}
}

func (s *SQLConnection) GetConnection() *sql.DB {
	return s.db
}

func (s *SQLConnection) GetResultSetAsStrings(q string, len int) ([][]string, error) {
	rows, err := s.db.QueryContext(s.ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	resultSet := make([][]string, 0)

	for rows.Next() {
		res := make([]string, len)

		err := rows.Scan(&res)

		if err != nil {
			return nil, err
		}

		resultSet = append(resultSet, res)
	}

	return resultSet, nil
}
