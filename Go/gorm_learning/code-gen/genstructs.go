package code_gen

import (
	"golearning/gorm_learning/model"
	"gorm.io/gen"
)

func GenStructFromTable() {
	generator := gen.NewGenerator(gen.Config{
		OutPath: "D:/Users/Administrator/Desktop/workspace/golearning/gorm_learning/model/",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	generator.UseDB(model.DbConn)

	generator.GenerateModelAs("student", "Student", gen.FieldType("id", "int64"), gen.FieldType("age", "int8"))

	generator.Execute()
}
