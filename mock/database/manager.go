package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type MockDB struct {
	mockDB *sql.DB
	mock sqlmock.Sqlmock
	db *sqlx.DB
}

func (m *MockDB) Init() error {
	var err error
	m.mockDB, m.mock, err = sqlmock.New()
	if err != nil {
		return err
	}

	return nil
}

func (m *MockDB) Open() *sqlx.DB {
	if m.db == nil {
		m.db, _ = sqlx.Open("", "", func(driverName string, dataSourceName string) (db *sql.DB, err error) {
			return m.mockDB, nil
		})
	}

	return m.db
}

func (m *MockDB) GetDB() *sqlx.DB {
	return m.db
}

func (m *MockDB) LoadTestSuite(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	queries := make([]Query, 0)
	err = json.Unmarshal(data, &queries)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, q := range queries {
		if q.ExpectedSQLKeyWord == "TransactionBegin" {
			m.mock.ExpectBegin()
		} else if q.ExpectedSQLKeyWord == "TransactionCommit" {
			m.mock.ExpectCommit()
		} else if q.ExpectedSQLKeyWord == "TransactionRollback" {
			m.mock.ExpectRollback()
		} else {
			if q.Type == "exec" {
				exec := m.mock.ExpectExec(q.ExpectedSQLKeyWord)
				if q.WithArgs != nil {
					exec = exec.WithArgs(q.WithArgs)
				}
				if q.ReturnError != nil {
					exec.WillReturnError(q.ReturnError)
				} else if q.ReturnResult != nil {
					exec.WillReturnResult(sqlmock.NewResult(q.ReturnResult.LastInsertID, q.ReturnResult.RowsEffected))
				} else {
					return fmt.Errorf("error or result are all nil")
				}
			} else if q.Type == "query" {
				query := m.mock.ExpectQuery(q.ExpectedSQLKeyWord)
				if q.WithArgs != nil {
					query = query.WithArgs(q.WithArgs)
				}
				if q.ReturnError != nil {
					query.WillReturnError(q.ReturnError)
				} else if q.ReturnRows != nil {
					rows := sqlmock.NewRows(q.ReturnRows.Columns)
					for _, r := range q.ReturnRows.Rows {
						values := make([]driver.Value, 0)
						for _, v := range r {
							values = append(values, v)
						}
						rows.AddRow(values...)
					}
					query.WillReturnRows(rows)
				} else {
					return fmt.Errorf("error or rows are all nil")
				}
			}
		}
	}

	return nil
}