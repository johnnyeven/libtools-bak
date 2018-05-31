package gorm

import (
	"fmt"
	"strings"
)

func BeforeBatchCreate(scope *Scope) {
	scope.CallMethodWithErrorCheck("BeforeSave")
	scope.CallMethodWithErrorCheck("BeforeBatchCreate")
}

func UpdateTimeStampWhenBatchCreate(scope *Scope) {
	if !scope.HasError() {
		now := NowFunc()
		scope.SetColumn("CreatedAt", now)
		scope.SetColumn("UpdatedAt", now)
	}
}

func BatchCreate(scope *Scope) {
	defer scope.Trace(NowFunc())

	if !scope.HasError() {
		// set BatchCreate sql
		batchFields := scope.BatchFields()
		var batchSqls [][]string
		var batchColumns []string
		var travesalNames []string
		for _, fields := range batchFields {
			var sqls, columns []string
			if len(batchColumns) == 0 {
				for _, field := range fields {
					if scope.changeableField(field) {
						if field.IsNormal {
							if !field.IsPrimaryKey || (field.IsPrimaryKey && !field.IsBlank) {
								if !field.IsBlank || !field.HasDefaultValue {
									columns = append(columns, scope.Quote(field.DBName))
									travesalNames = append(travesalNames, field.DBName)
									sqls = append(sqls, scope.AddToVars(field.Field.Interface()))
								}
							}
						} else if relationship := field.Relationship; relationship != nil && relationship.Kind == "belongs_to" {
							if relationField := fields[relationship.ForeignDBName]; !scope.changeableField(relationField) {
								columns = append(columns, scope.Quote(relationField.DBName))
								travesalNames = append(travesalNames, field.DBName)
								sqls = append(sqls, scope.AddToVars(relationField.Field.Interface()))
							}
						}
					}
				}
				batchColumns = columns
				if len(batchColumns) == 0 {
					// need break,if the columns is empty.
					break
				}
			} else {
				for _, col := range travesalNames {
					field := fields[col]
					if scope.changeableField(field) {
						if field.IsNormal {
							if !field.IsPrimaryKey || (field.IsPrimaryKey && !field.IsBlank) {
								if !field.IsBlank || !field.HasDefaultValue {
									sqls = append(sqls, scope.AddToVars(field.Field.Interface()))
								}
							}
						} else if relationship := field.Relationship; relationship != nil && relationship.Kind == "belongs_to" {
							if relationField := fields[relationship.ForeignDBName]; !scope.changeableField(relationField) {
								sqls = append(sqls, scope.AddToVars(relationField.Field.Interface()))
							}
						}
					}
				}
			}
			batchSqls = append(batchSqls, sqls)
		}

		returningKey := "*"
		primaryField := scope.PrimaryField()
		if primaryField != nil {
			returningKey = scope.Quote(primaryField.DBName)
		}

		var BatchCreate_sql string = "INSERT INTO"
		var extraOption string

		if insert_ignore, ok := scope.InstanceGet("gorm:insert_ignore"); ok {
			if insert_ignore.(bool) {
				BatchCreate_sql = "INSERT IGNORE INTO"
			}
		}
		if str, ok := scope.Get("gorm:insert_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if len(batchColumns) == 0 {
			scope.Raw(fmt.Sprintf("%s %v DEFAULT VALUES%v%v",
				BatchCreate_sql,
				scope.QuotedTableName(),
				addExtraSpaceIfExist(extraOption),
				addExtraSpaceIfExist(scope.Dialect().ReturningStr(scope.TableName(), returningKey)),
			))
		} else {
			rows := []string{}
			for _, sqls := range batchSqls {
				tmpStr := "(" + strings.Join(sqls, ",") + ")"
				rows = append(rows, tmpStr)
			}
			scope.Raw(fmt.Sprintf(
				"%s %v (%v) VALUES %v %v%v",
				BatchCreate_sql,
				scope.QuotedTableName(),
				strings.Join(batchColumns, ","),
				strings.Join(rows, ","),
				addExtraSpaceIfExist(extraOption),
				addExtraSpaceIfExist(scope.Dialect().ReturningStr(scope.TableName(), returningKey)),
			))
		}

		// execute BatchCreate sql
		if scope.Dialect().SupportLastInsertId() {
			if result, err := scope.SqlDB().Exec(scope.Sql, scope.SqlVars...); scope.Err(err) == nil {
				id, err := result.LastInsertId()
				if scope.Err(err) == nil && id != 0 {
					scope.db.RowsAffected, _ = result.RowsAffected()
					if autoIncrementField := scope.AutoIncrementField(); autoIncrementField != nil {
						scope.Err(scope.SetColumn(autoIncrementField, id))
					}
				}
			}
		} else {
			if primaryField == nil {
				if results, err := scope.SqlDB().Exec(scope.Sql, scope.SqlVars...); err != nil {
					scope.db.RowsAffected, _ = results.RowsAffected()
				}
			} else if scope.Err(scope.SqlDB().QueryRow(scope.Sql, scope.SqlVars...).Scan(primaryField.Field.Addr().Interface())) == nil {
				scope.db.RowsAffected = 1
			}
		}
	}
}

func AfterBatchCreate(scope *Scope) {
	scope.CallMethodWithErrorCheck("AfterBatchCreate")
	scope.CallMethodWithErrorCheck("AfterSave")
}

func init() {
	DefaultCallback.BatchCreate().Register("gorm:before_create", BeforeBatchCreate)
	DefaultCallback.BatchCreate().Register("gorm:save_before_associations", SaveBeforeAssociations)
	DefaultCallback.BatchCreate().Register("gorm:update_time_stamp_when_create", UpdateTimeStampWhenCreate)
	DefaultCallback.BatchCreate().Register("gorm:create", BatchCreate)
	DefaultCallback.BatchCreate().Register("gorm:save_after_associations", SaveAfterAssociations)
}
