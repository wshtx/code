package crud_test

import (
	"database/sql"
	"fmt"
	"golearning/gorm_learning/model"
	"gorm.io/gorm"
)

func TestGorm() {
	//db := model.DbConn

	//after a chained method and finished method,GORM will return an initialized *gorm.DB instance,which is not safe to reuse anymore.
	//And you can use New Session method to reuse the *gorm.DB.GORM defined Session,WithContext,Debug methods as| New Session method.
	//
	//Bad case
	//tx := db.Where("age > ?", 10)
	//tx.Where("name > ?","htx").Find(&model.Student{})//select * from student where age > 10 and name != "htx"
	//tx.Where("id != ?", 1).Find(&model.Student{})//select * from student where age > 10 and name != "htx" and id != 1
	//
	//good example
	//tx := db.Where("age > ?", 10).Debug()
	//tx := db.Where("age > ?", 10).Session(&gorm.Session{})
	//tx := db.Where("age > ?", 10).WithContext(context.Background())
	//tx.Where("name != ?", "htx").Find(&model.Student{}) //select * from student where age > 10 and name != "htx"
	//tx.Where("id != ?", 1).Find(&model.Student{})       //select * from student where age > 10 and id != 1

	//raw SQL
	//var res0 []string
	//db.Raw("select distinct name from student").Scan(&res0)
	//fmt.Println(res0)

	//dryrun mode, just only generate the Sql and the arguments without executing,can be used to prepare or test generated SQL
	//var student model.Student
	//statement := db.Session(&gorm.Session{DryRun: true}).Where("name = ?", "htx").Where("age", "24").Find(&student).Statement
	//fmt.Println(statement.SQL.String())
	//fmt.Println(statement.Vars)

	//TOSQL,just only generate the Sql and the arguments without executing,can be used to prepare or test generated SQL
	//toSQL := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	//	return tx.Where("name = ?", "htx").Where("age", "24").Find(&model.Student{})
	//})
	//fmt.Println(toSQL)

	//sql.Rows
	//rows, _ := db.Model(&model.Student{}).Select("*").Rows()
	//defer rows.Close()
	//var res1 model.Student
	//for rows.Next() {
	//	//scan into the variables
	//	//name := ""
	//	//age := -1
	//	//id := -1
	//	//rows.Scan(&id, &name, &age)
	//	//fmt.Printf("{%v %v %v}\n", id, name, age)
	//
	//	//recommened,scan into struct
	//	db.ScanRows(rows, &res1)
	//	fmt.Println(res1)
	//}

	//TestCreateGorm(db)
	//TestRetriveGorm(db)
	//TestUpdateAndDelete(db)
	//TestHooks(db)
	//TestBatchUpdateAndDelete(db)
}

func TestCreateGorm(db *gorm.DB) {
	//db.Create(&model.Student{Name: "htx", Age: 24})
	//db.Create(&model.Student{Name: "alj", Age: 18})
	//db.Create(&model.Student{Name: "hnb", Age: 6})

}

func TestRetriveGorm(db *gorm.DB) {
	var student model.Student
	//db.First(&student)
	result := map[string]interface{}{}
	db.Model(&student).First(result)
	fmt.Println(result)

	//select by the primary id
	db.Find(&student, 1)
	fmt.Println(result)

	//select all obejcts and specify order
	var students []model.Student
	db.Order("id desc").Find(&students)
	fmt.Println(students)

	//string with placeholder,struct or maps as the condition
	//
	//the condition string
	//"?" as the placeholder in the SQL statement
	db.Where("name = ?", "htx").Where("age", "24").Find(&student)
	fmt.Println(student)
	//named argument
	var res00 model.Student
	db.Where("name != @name", sql.Named("name", "htx")).Find(&res00)
	fmt.Println(res00)
	//struct or maps as the condition
	condition1 := model.Student{Name: "htx", Age: 11}
	condition2 := map[string]interface{}{"age": 24}
	fmt.Println(student)
	db.Where(&condition1).Where(condition2).Find(&student)
	fmt.Println(student)

	//specify the struct search field
	var res1 model.Student
	db.Where(&condition1, "age").Find(&res1)
	fmt.Println(res1)

	//query condition can be inlined into method
	var res2 model.Student
	db.Find(&res2, "name=?", "alj")
	fmt.Println(res2)

	//select specific field
	var res3 model.Student
	db.Select("name").Where("id", 1).Find(&res3)
	fmt.Println(res3)

	//group by & having
	var res4 []model.Student
	db.Select("age").Group("age").Having("age < ?", 20).Find(&res4)
	fmt.Println(res4)

	//distinct
	var res5 []model.Student
	db.Distinct("age").Select("age").Find(&res5)
	fmt.Println(res5)

	//scan the result into structs
	var res6 []string
	db.Model(&model.Student{}).Select("name").Where("name != ?", "htx").Scan(&res6)
	fmt.Println(res6)
}

func TestUpdateAndDelete(db *gorm.DB) {
	//update single column
	//db.Model(&model.Student{}).Where("name = ?", "wshtx").Update("age", 17)
	//update a table by using subquery
	//db.Model(&model.Student{ID: 4}).Update("age", db.Model(&model.Student{}).Select("max(age)").Where("name != ?", "htx"))//error ,just as a example

	//update multiple columns
	//
	//update attributes with struct
	//db.Model(&model.Student{ID: 1}).Updates(model.Student{Name: "htx", Age: 16})

	//update attributes with maps
	//db.Model(&model.Student{ID: 1}).Updates(map[string]interface{}{"name": "htxdashuaibi", "age": 8})

	//update the selected field, Select()/Omit()
	//db.Model(&model.Student{ID: 1}).Select("name").Updates(map[string]interface{}{"name": "htx", "age": 24})
	//db.Model(&model.Student{ID: 1}).Omit("name").Updates(map[string]interface{}{"name": "htx", "age": 24})

	//update a column with a SQL expression,e.g:
	//db.Model(&model.Student{ID: 1}).UpdateColumn("age", gorm.Expr("age + ?", 2))

	//delete with primary key
	//db.Delete(&model.Student{}, 4)
	//db.Where("age = ?", 10).Delete(&model.Student{})

	//soft delete
	//db.Where("id=?", 4).Delete(&model.Student{})
	//var student model.Student
	//db.Unscoped().Where("id=?", 4).Find(&student)
	//fmt.Println(student)
	//db.Unscoped().Delete(&model.Student{ID: 4})

}

// if we haven`t specified a record having a primary key with Modem,Gorm will perform a batch update
func TestBatchUpdateAndDelete(db *gorm.DB) {
	//need to use where condition
	//db.Model(&model.Student{}).Where("age < ?", 18).Updates(map[string]interface{}{"age": 8})

	//haven`t use any condition,you can use raw SQL or enable the "ALLowGlobalUpdate" mode, as following:
	//db.Model(&model.Student{}).Updates(map[string]interface{}{"age": 8}) //error
	//db.Delete(&model.Student{})//error
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&model.Student{}).Updates(map[string]interface{}{"age": 8}) //success
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Student{})//success

	//get the updated records count
	//updateRes := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&model.Student{}).Updates(map[string]interface{}{"age": 18})
	//fmt.Printf("affected rows:%v\n", updateRes.RowsAffected)
}

func TestHooks(db *gorm.DB) {
	// the hooks methods beforeUpdate beforeSave afterUpdate afterSave already are implemented in the Student
	//db.Model(&model.Student{ID: 1}).Updates(map[string]interface{}{"name": "wshtx"})

	//use the UpdateColumn/UpdateColumns like  Update/Updates will skip the hook method
	//db.Model(&model.Student{ID: 1}).UpdateColumns(map[string]interface{}{"name": "htx"})

	//delete hooks, beforehooks,afterhooks
	//db.Where("id = ?", 4).Delete(&model.Student{})

}
