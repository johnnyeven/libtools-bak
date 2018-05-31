package gorm

import (
	"fmt"
	"strings"
)

func BeforeCreate(scope *Scope) {
	scope.CallMethodWithErrorCheck("BeforeSave")
	scope.CallMethodWithErrorCheck("BeforeCreate")
}

func UpdateTimeStampWhenCreate(scope *Scope) {
	if !scope.HasError() {
		now := NowFunc()
		scope.SetColumn("CreatedAt", now)
		scope.SetColumn("UpdatedAt", now)
	}
}

func Create(scope *Scope) {
	defer scope.Trace(NowFunc())

	if !scope.HasError() {
		// set create sql
		var sqls, columns []string
		fields := scope.Fields()
		for _, field := range fields {
			if scope.changeableField(field) {
				if field.IsNormal {
					if !field.IsPrimaryKey || (field.IsPrimaryKey && !field.IsBlank) {
						if !field.IsBlank || !field.HasDefaultValue {
							columns = append(columns, scope.Quote(field.DBName))
							sqls = append(sqls, scope.AddToVars(field.Field.Interface()))
						}
					}
				} else if relationship := field.Relationship; relationship != nil && relationship.Kind == "belongs_to" {
					if relationField := fields[relationship.ForeignDBName]; !scope.changeableField(relationField) {
						columns = append(columns, scope.Quote(relationField.DBName))
						sqls = append(sqls, scope.AddToVars(relationField.Field.Interface()))
					}
				}
			}
		}

		returningKey := "*"
		primaryField := scope.PrimaryField()
		if primaryField != nil {
			returningKey = scope.Quote(primaryField.DBName)
		}

		var create_sql string = "INSERT INTO"
		var extraOption string

		if insert_ignore, ok := scope.InstanceGet("gorm:insert_ignore"); ok {
			if insert_ignore.(bool) {
				create_sql = "INSERT IGNORE INTO"
			}
		}
		if str, ok := scope.Get("gorm:insert_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if len(columns) == 0 {
			scope.Raw(fmt.Sprintf("%s %v DEFAULT VALUES%v%v",
				create_sql,
				scope.QuotedTableName(),
				addExtraSpaceIfExist(extraOption),
				addExtraSpaceIfExist(scope.Dialect().ReturningStr(scope.TableName(), returningKey)),
			))
		} else {
			scope.Raw(fmt.Sprintf(
				"%s %v (%v) VALUES (%v)%v%v",
				create_sql,
				scope.QuotedTableName(),
				strings.Join(columns, ","),
				strings.Join(sqls, ","),
				addExtraSpaceIfExist(extraOption),
				addExtraSpaceIfExist(scope.Dialect().ReturningStr(scope.TableName(), returningKey)),
			))
		}

		// execute create sql
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

func AfterCreate(scope *Scope) {
	scope.CallMethodWithErrorCheck("AfterCreate")
	scope.CallMethodWithErrorCheck("AfterSave")
}

func init() {
	DefaultCallback.Create().Register("gorm:before_create", BeforeCreate)
	DefaultCallback.Create().Register("gorm:save_before_associations", SaveBeforeAssociations)
	DefaultCallback.Create().Register("gorm:update_time_stamp_when_create", UpdateTimeStampWhenCreate)
	DefaultCallback.Create().Register("gorm:create", Create)
	DefaultCallback.Create().Register("gorm:save_after_associations", SaveAfterAssociations)
}
