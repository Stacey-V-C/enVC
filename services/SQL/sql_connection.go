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

	err = ConnectToDefault(db)

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

func ConnectToDefault(db *sql.DB) error {
	dbName, exists := os.LookupEnv("SQL_DB_NAME")

	if exists == false {
		dbName = "envc"
	}

	_, err := db.Exec(
		"CREATE DATABASE IF NOT EXISTS " + dbName + ";" +
			"USE " + dbName + ";",
	)

	return err // even if nil
}

func (s *SQLConnection) Init() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS topics (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(64) NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,

		`CREATE TABLE IF NOT EXISTS comments (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			p_type ENUM(
				'topic', 
				'tag', 
				'tag_type', 
				'comment', 
				'connection', 
				'directed_relation
			) NOT NULL DEFAULT 'topic',
			p_id INT NOT NULL,
			content TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,

		`CREATE TABLE IF NOT EXISTS tags (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			p_type ENUM(
				'topic',
				'tag',
				'tag_type',
				'comment',
				'connection',
				'directed_relation'
			) NOT NULL DEFAULT 'topic',
			p_id INT NOT NULL,
			type INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,

		`CREATE TABLE IF NOT EXISTS tag_types (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(128) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,

		`CREATE TABLE IF NOT EXISTS connections (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			p_1_type ENUM(
				'topic',
				'tag',
				'tag_type',
				'comment',
				'connection',
				'directed_relation'
			) NOT NULL DEFAULT 'topic',
			p_1_id INT NOT NULL,
			p_2_type ENUM(
				'topic',
				'tag',
				'tag_type',
				'comment',
				'connection',
				'directed_relation'
			) NOT NULL DEFAULT 'topic',
			p_2_id INT NOT NULL,
			content TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,

		`CREATE TABLE IF NOT EXISTS directed_relations (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			s_p_type ENUM(
				'topic',
				'tag',
				'tag_type',
				'comment',
				'connection',
				'directed_relation'
			) NOT NULL DEFAULT 'topic',
			s_p_id INT NOT NULL,
			e_p_type ENUM(
				'topic',
				'tag',
				'tag_type',
				'comment',
				'connection',
				'directed_relation'
			) NOT NULL DEFAULT 'topic',
			e_p_id INT NOT NULL,
			content TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`,
	}

	err := ConnectToDefault(s.db)

	if err != nil {
		return err
	}

	for _, q := range queries {
		_, err := s.db.ExecContext(s.ctx, q)

		if err != nil {
			return err
		}
	}

	return nil
}
