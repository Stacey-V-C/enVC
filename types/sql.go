package types

type Table struct {
	name    string
	columns []string
}

func NewTable(name string, columns []string) Table {
	return Table{name: name, columns: columns}
}

func (t *Table) GetColumns() []string {
	return t.columns
}

func (t *Table) GetName() string {
	return t.name
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

func NewDataModel(tables []Table) DataModel {
	return DataModel{tables: tables}
}

func (d *DataModel) GetTable(tableName string) *Table {
	for _, table := range d.tables {
		if table.name == tableName {
			return &table
		}
	}

	return nil
}

type SQLResults struct {
	table     Table
	statement string
	results   [][]string
}
