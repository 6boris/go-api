package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"fmt"
	"github.com/dicf/go-api/pkg/configure"
	"log"
	"time"
)

var db_mysql *gorm.DB

func Setup() {
	c := configure.Cfg

	var err error
	db_mysql, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQL.UserName,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Database))
	db_mysql = db_mysql.LogMode(true)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db_mysql.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return setting.DatabaseSetting.TablePrefix + defaultTableName
	//}

	//db, err = gorm.Open("mysql", "root:@/blog?charset=utf8&parseTime=True&loc=Asia%2FShanghai")

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	//return c.MySQL.TablePrefix + defaultTableName
	//	return c.MySQL.TablePrefix + defaultTableName
	//}

	db_mysql.SingularTable(true)
	//db_mysql.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db_mysql.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db_mysql.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db_mysql.DB().SetMaxIdleConns(10)
	db_mysql.DB().SetMaxOpenConns(100)
	//db_mysql.Callback().Query().After("gorm:query").Register("my_plugin:after_query", afterQuery)
	//db_mysql.Callback().Query().After("gorm:query").Register("my_plugin:before_query", afterQuery)
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ := createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ := modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
