package sqlx

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"

	"golib/tools/sqlx/builder"
)

func NewDatabase(name string) *Database {
	return &Database{
		Database: builder.DB(name),
	}
}

type Database struct {
	*builder.Database
}

func (database *Database) Register(model Model) *builder.Table {
	database.mustStructType(model)
	rv := reflect.Indirect(reflect.ValueOf(model))
	table := builder.T(database.Database, model.TableName())
	ScanDefToTable(rv, table)
	database.Database.Register(table)
	return table
}

func (database Database) T(model Model) *builder.Table {
	database.mustStructType(model)
	return database.Database.Table(model.TableName())
}

func (database Database) mustStructType(model Model) {
	tpe := reflect.TypeOf(model)
	if tpe.Kind() != reflect.Ptr {
		panic(fmt.Errorf("model %s must be a pointer", tpe.Name()))
	}
	tpe = tpe.Elem()
	if tpe.Kind() != reflect.Struct {
		panic(fmt.Errorf("model %s must be a struct", tpe.Name()))
	}
}

func (database *Database) Insert(model Model) *builder.StmtInsert {
	table := database.T(model)

	fieldValues := FieldValuesFromStructByNonZero(model)

	if autoIncrementCol := table.AutoIncrement(); autoIncrementCol != nil {
		delete(fieldValues, autoIncrementCol.FieldName)
	}

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	return table.Insert().Columns(cols).Values(vals...)
}

func (database *Database) Update(model Model, zeroFields ...string) *builder.StmtUpdate {
	table := database.T(model)

	fieldValues := FieldValuesFromStructByNonZero(model, zeroFields...)

	if autoIncrementCol := table.AutoIncrement(); autoIncrementCol != nil {
		delete(fieldValues, autoIncrementCol.FieldName)
	}

	return table.Update().Set(table.AssignsByFieldValues(fieldValues)...)
}

func (database *Database) MigrateTo(db *DB, dryRun bool) error {
	logrus.Debugf("=================== migrating database `%s` ====================", database.Name)
	defer logrus.Debugf("=================== migrated database `%s` ====================", database.Name)

	currentDatabase := DBFromInformationSchema(db, database.Name, database.Tables.TableNames()...)

	if !dryRun {
		tasks := NewTasks(db)

		if currentDatabase == nil {
			currentDatabase = NewDatabase(database.Name)
			tasks = tasks.With(func(db *DB) error {
				return db.Do(currentDatabase.Create(true)).Err()
			})
		}

		for name, table := range database.Tables {
			currentTable := currentDatabase.Table(name)
			if currentTable == nil {
				stmt := table.Create(true)
				tasks = tasks.With(func(db *DB) error {
					return db.Do(stmt).Err()
				})
				continue
			}

			stmt := currentTable.Diff(table)
			if stmt != nil {
				tasks = tasks.With(func(db *DB) error {
					return db.Do(stmt).Err()
				})
				continue
			}
		}

		err := tasks.Do()
		if err != nil {
			return err
		}
		return nil
	}

	if currentDatabase == nil {
		currentDatabase = NewDatabase(database.Name)
		fmt.Println(currentDatabase.Create(true).Query)
	}

	for name, table := range database.Tables {
		currentTable := currentDatabase.Table(name)
		if currentTable == nil {
			fmt.Println(table.Create(true).Query)
			continue
		}

		stmt := currentTable.Diff(table)
		if stmt != nil {
			fmt.Println(stmt.Query)
			continue
		}
	}

	return nil
}
