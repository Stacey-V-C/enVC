package services

import (
	"errors"
	"fmt"
	ut "nvc/util"
	"regexp"
	"strings"
)

type SQLProcesser struct {
	//model *DataModel // maybe?  or the listener can just have the context to pass this
}

func NewSQLProcesser() SQLProcesser {
	return SQLProcesser{}
}

type Table struct {
	name    string
	columns []string
}

type DataModel struct {
	tables []Table
}

func (d *DataModel) GetTableNames() []string {
	tableNames := make([]string, len(d.tables))

	for i, table := range d.tables {
		tableNames[i] = table.name
	}

	return tableNames
}

func (d *DataModel) GetTable(tableName string) *Table {
	for _, table := range d.tables {
		if table.name == tableName {
			return &table
		}
	}

	return nil
}

var NVCDataModel = DataModel{
	tables: []Table{
		{
			name: "topics",
			columns: []string{
				"id",
				"title",
				"description",
				"created_at",
			},
		},
	},
}

// requestedColumns, err, updatedQuery := parseQueryReturning(model, query)

func (s *SQLProcesser) filterResults(requestedCols []string, resultSet [][]string, resultCols []string) ([][]string, error) {
	for _, col := range requestedCols {
		if ut.Contains(resultCols, col) == false {
			errorMessage := fmt.Sprintf("Column %s does not exist in result set", col)
			return nil, errors.New(errorMessage)
		}
	}

	filteredResultSet := make([][]string, len(resultSet))

	for i, row := range resultSet {
		filteredRow := make([]string, len(requestedCols))

		for j, col := range requestedCols {
			k := ut.GetIndex(resultCols, col)

			filteredRow[j] = row[k]
		}

		filteredResultSet[i] = filteredRow
	}

	return filteredResultSet, nil
}

func (s *SQLProcesser) parseQueryReturning(model DataModel, query string) (*Table, *string, error) {
	query = strings.ReplaceAll(query, "\n", " ")
	query = strings.ReplaceAll(query, "\t", " ")

	command := strings.Split(query, " ")[0]

	if command == "SELECT" {
		return s.parseSelectQuery(model, query)
	} else if command == "INSERT" {
		return s.parseInsertQuery(model, query)
	} else {
		return nil, nil, nil
	}
}

func (s *SQLProcesser) parseSelectQuery(model DataModel, query string) (*Table, *string, error) {
	getSelectDb, err := regexp.Compile(`(?i)SELECT\s+(.+)\s+FROM\s+([a-z_]+)\s`)
	if err != nil {
		print(err)
		print('\n')
	}

	matches := getSelectDb.FindStringSubmatch(query)
	if matches == nil {
		return nil, nil, errors.New("Invalid SELECT query")
	}
	columns := strings.Split(matches[1], ",")

	for i, column := range columns {
		columns[i] = strings.TrimSpace(column)
	}

	target := strings.TrimSpace(matches[2])

	tableNames := model.GetTableNames()

	if ut.Contains[string](tableNames, target) == false {
		errorMessage := fmt.Sprintf("Table %s does not exist in database", target)
		return nil, nil, errors.New(errorMessage)
	}

	table := model.GetTable(target)

	if columns[0] != "*" {
		err := validateColumns(table, columns)

		if err != nil {
			return nil, nil, err
		}
	} else {
		columns = table.columns
	}

	return &Table{
		name:    table.name,
		columns: columns,
	}, &query, nil
}

func (s *SQLProcesser) parseInsertQuery(model DataModel, query string) (*Table, *string, error) {
	getTableName, err := regexp.Compile(`(?i)INSERT\s+INTO\s+([a-z_]+)\s`)
	if err != nil {
		return nil, nil, err
	}

	getReturningStmt, err := regexp.Compile(`(?i)RETURNING\s+(.+);`)
	if err != nil {
		return nil, nil, err
	}

	tableNameMatch := getTableName.FindStringSubmatch(query)

	if tableNameMatch == nil {
		return nil, nil, errors.New("Invalid INSERT query")
	}

	tableName := strings.TrimSpace(tableNameMatch[1])

	table := model.GetTable(tableName)

	if table == nil {
		errorMessage := fmt.Sprintf("Table %s does not exist in database", tableName)
		return nil, nil, errors.New(errorMessage)
	}

	returningStmtMatches := getReturningStmt.FindStringSubmatch(query)

	var columns []string

	if returningStmtMatches == nil {
		newQuery := []string{
			query[:len(query)-1],
			" RETURNING *;",
		}

		query = strings.Join(newQuery, "")
	} else {
		if returningStmtMatches[1] == "*" {
			columns = table.columns
		} else {
			columns = strings.Split(returningStmtMatches[1], ",")

			for i, column := range columns {
				columns[i] = strings.TrimSpace(column)
			}

			err := validateColumns(table, columns)

			if err != nil {
				return nil, nil, err
			}

			query = strings.ReplaceAll(query, returningStmtMatches[0], "RETURNING *;")
		}
	}

	return &Table{
		name:    table.name,
		columns: columns,
	}, &query, nil
}

func validateColumns(table *Table, given_columns []string) error {
	for _, c := range given_columns {
		if ut.Contains[string](table.columns, c) == false {
			errorMessage := fmt.Sprintf("Column %s does not exist in table %s", c, table.name)
			return errors.New(errorMessage)
		}
	}

	return nil
}
