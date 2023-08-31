package services

import (
	"fmt"
	"nvc/types"
	"testing"
)

func TestParseInsertQuery(t *testing.T) {
	mSqlController := SQLProcesser{}

	mBaseTable := types.NewTable(
		"users",
		[]string{
			"id",
			"username",
			"password",
		},
	)

	mDataModel := types.NewDataModel(
		[]types.Table{mBaseTable})

	mIdUserNameTable := types.NewTable(
		"users",
		[]string{
			"id",
			"username",
		},
	)

	mNoReturnTable := types.NewTable(
		"users",
		[]string{},
	)

	queries := map[string]string{
		"withWildcard":            "INSERT INTO users (username, password) VALUES ('test', 'test') RETURNING *;",
		"withColumns":             "INSERT INTO users (username, password) VALUES ('test', 'test') RETURNING id, username;",
		"withoutReturnStmt":       "INSERT INTO users (username, password) VALUES ('test', 'test');",
		"withInvalidReturnColumn": "INSERT INTO users (username, password) VALUES ('test', 'test') RETURNING id, username, invalid_column;",
		"withInvalidTable":        "INSERT INTO invalid_table (username, password) VALUES ('test', 'test');",
	}

	var cases = []struct {
		testName      string
		query         string
		expectedTable *types.Table
		expectedError error
		expectedQuery string
	}{
		{
			testName:      "Already returning wildcard",
			query:         queries["withWildcard"],
			expectedTable: &mBaseTable,
			expectedQuery: queries["withWildcard"],
			expectedError: nil,
		},
		{
			testName:      "Returning valid columns",
			query:         queries["withColumns"],
			expectedTable: &mIdUserNameTable,
			expectedQuery: queries["withWildcard"],
			expectedError: nil,
		},
		{
			testName:      "No return statement",
			query:         queries["withoutReturnStmt"],
			expectedTable: &mNoReturnTable,
			expectedQuery: queries["withWildcard"],
			expectedError: nil,
		},
		{
			testName:      "Invalid return column",
			query:         queries["withInvalidReturnColumn"],
			expectedTable: nil,
			expectedQuery: "",
			expectedError: fmt.Errorf("Column invalid_column does not exist in table users"),
		},
		{
			testName:      "Invalid table",
			query:         queries["withInvalidTable"],
			expectedTable: nil,
			expectedQuery: "",
			expectedError: fmt.Errorf("Table invalid_table does not exist in database"),
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			table, query, err := mSqlController.parseInsertQuery(mDataModel, c.query)

			if err != nil && c.expectedError == nil {
				t.Errorf("Expected error to be nil, got %s", err)
			}

			if err == nil && c.expectedError != nil {
				t.Errorf("Expected error to be %s, got nil", c.expectedError)
			}

			if table != nil && c.expectedTable == nil {
				t.Errorf("Expected table to be nil, got %s", table)
			}

			if table == nil && c.expectedTable != nil {
				t.Errorf("Expected table to be %s, got nil", c.expectedTable)
			}

			if table != nil && c.expectedTable != nil {
				if table.GetName() != c.expectedTable.GetName() {
					t.Errorf("Expected table name to be %s, got %s", c.expectedTable.GetName(), table.GetName())
				}

				if len(table.GetColumns()) != len(c.expectedTable.GetColumns()) {
					t.Errorf("Expected table columns to be %s, got %s", c.expectedTable.GetColumns(), table.GetColumns())
				}

				for i, column := range table.GetColumns() {
					if column != c.expectedTable.GetColumns()[i] {
						t.Errorf("Expected table column to be %s, got %s", c.expectedTable.GetColumns()[i], column)
					}
				}
			}

			if c.expectedQuery == "" && query != nil {
				t.Errorf("Expected query to be nil, got %s", *query)
			} else if c.expectedQuery != "" && query == nil {
				t.Errorf("Expected query to be %s, got nil", c.expectedQuery)
			} else {
				if query != nil && *query != c.expectedQuery {
					t.Errorf("Expected query to be %s, got %s", c.expectedQuery, *query)
				}
			}
		})
	}
}
