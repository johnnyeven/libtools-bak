package gorm

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type commonDialect struct{}

func (commonDialect) BinVar(i int) string {
	return "$$" // ?
}

func (commonDialect) SupportLastInsertId() bool {
	return true
}

func (commonDialect) HasTop() bool {
	return false
}

func (commonDialect) SqlTag(value reflect.Value, size int, autoIncrease bool) string {
	switch value.Kind() {
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		if autoIncrease {
			return "INTEGER AUTO_INCREMENT"
		}
		return "INTEGER"
	case reflect.Int64, reflect.Uint64:
		if autoIncrease {
			return "BIGINT AUTO_INCREMENT"
		}
		return "BIGINT"
	case reflect.Float32, reflect.Float64:
		return "FLOAT"
	case reflect.String:
		if size > 0 && size < 65532 {
			return fmt.Sprintf("VARCHAR(%d)", size)
		}
		return "VARCHAR(65532)"
	case reflect.Struct:
		if _, ok := value.Interface().(time.Time); ok {
			return "TIMESTAMP"
		}
	default:
		if _, ok := value.Interface().([]byte); ok {
			if size > 0 && size < 65532 {
				return fmt.Sprintf("BINARY(%d)", size)
			}
			return "BINARY(65532)"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s) for commonDialect", value.Type().Name(), value.Kind().String()))
}

func (commonDialect) ReturningStr(tableName, key string) string {
	return ""
}

func (commonDialect) SelectFromDummyTable() string {
	return ""
}

func (commonDialect) Quote(key string) string {
	return fmt.Sprintf(`"%s"`, key)
}

func (commonDialect) databaseName(scope *Scope) string {
	from := strings.Index(scope.db.parent.source, "/") + 1
	to := strings.Index(scope.db.parent.source, "?")
	if to == -1 {
		to = len(scope.db.parent.source)
	}
	return scope.db.parent.source[from:to]
}

func (c commonDialect) HasTable(scope *Scope, tableName string) bool {
	var count int
	dbName, realTableName := DBName(tableName)
	if dbName == "" {
		dbName = c.databaseName(scope)
	}
	scope.NewDB().Raw("SELECT count(*) FROM INFORMATION_SCHEMA.TABLES WHERE table_name = ? AND table_schema = ?", realTableName, dbName).Row().Scan(&count)

	return count > 0
}

func (c commonDialect) HasColumn(scope *Scope, tableName string, columnName string) bool {
	var count int
	dbName, realTableName := DBName(tableName)
	if dbName == "" {
		dbName = c.databaseName(scope)
	}
	scope.NewDB().Raw("SELECT count(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE table_schema = ? AND table_name = ? AND column_name = ?", dbName, realTableName, columnName).Row().Scan(&count)
	return count > 0
}

func (c commonDialect) HasIndex(scope *Scope, tableName string, indexName string) bool {
	var count int
	dbName, realTableName := DBName(tableName)
	if dbName == "" {
		dbName = c.databaseName(scope)
	}
	scope.NewDB().Raw("SELECT count(*) FROM INFORMATION_SCHEMA.STATISTICS where table_name = ? AND index_name = ? and table_schema = ?", realTableName, indexName, dbName).Row().Scan(&count)
	return count > 0
}

func (c commonDialect) IndexColumnCountMap(scope *Scope, tableName string) map[string]int {
	dbName, realTableName := DBName(tableName)
	if dbName == "" {
		dbName = c.databaseName(scope)
	}
	rows, err := scope.NewDB().Raw("SELECT INDEX_NAME, COUNT(INDEX_NAME) FROM INFORMATION_SCHEMA.STATISTICS WHERE table_name = ? and table_schema = ? GROUP BY INDEX_NAME", realTableName, dbName).Rows()
	if err != nil {
		panic(err.Error())
	}

	indexColumnMap := make(map[string]int, 16)
	for rows.Next() {
		var indexName string
		var indexCount int
		if err = rows.Scan(&indexName, &indexCount); err != nil {
			panic(err.Error())
		}
		indexColumnMap[indexName] = indexCount
	}
	return indexColumnMap
}

func (c commonDialect) IndexColumnMap(scope *Scope, tableName string, NonUnique int) map[string][]string {
	dbName, realTableName := DBName(tableName)
	if dbName == "" {
		dbName = c.databaseName(scope)
	}
	rows, err := scope.NewDB().Raw("SELECT INDEX_NAME, SEQ_IN_INDEX, COLUMN_NAME FROM INFORMATION_SCHEMA.STATISTICS WHERE NON_UNIQUE = ? AND table_name = ? and table_schema = ?",
		NonUnique, realTableName, dbName).Rows()
	if err != nil {
		panic(err)
	}

	indexCountMap := c.IndexColumnCountMap(scope, tableName)

	indexColumnMap := make(map[string][]string, 16)
	for rows.Next() {
		var indexName, columnName string
		var seqInIndex int
		if err = rows.Scan(&indexName, &seqInIndex, &columnName); err != nil {
			panic(err)
		}
		if _, ok := indexColumnMap[indexName]; !ok {
			indexColumnMap[indexName] = make([]string, indexCountMap[indexName], indexCountMap[indexName])
		}
		indexColumnMap[indexName][seqInIndex-1] = columnName
	}
	return indexColumnMap
}

func (commonDialect) RemoveIndex(scope *Scope, indexName string) {
	scope.NewDB().Exec(fmt.Sprintf("DROP INDEX %v ON %v", indexName, scope.QuotedTableName()))
}

func (c commonDialect) Columns(scope *Scope, tableName string) map[string]string {
	dbName, realTableName := DBName(tableName)
	if dbName == "" {
		dbName = c.databaseName(scope)
	}
	rows, err := scope.NewDB().Raw(
		"SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, COLUMN_DEFAULT, EXTRA FROM INFORMATION_SCHEMA.COLUMNS WHERE table_schema = ? AND table_name = ?",
		dbName, realTableName,
	).Rows()
	if err != nil {
		panic(err)
	}

	columns := make(map[string]string, 32)
	for rows.Next() {
		var columnName, columnType, isNullable, extra string
		var columnDefault sql.NullString
		if err = rows.Scan(&columnName, &columnType, &isNullable, &columnDefault, &extra); err != nil {
			panic(err)
		}

		column := columnType
		if extra != "" {
			column += " " + extra
		}

		if isNullable == "NO" {
			column += " NOT NULL"
		}

		if columnDefault.Valid {
			column += " DEFAULT " + columnDefault.String
		}
		columns[columnName] = column
	}
	return columns
}
