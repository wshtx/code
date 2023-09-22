package code_gen

import (
	"golearning/gorm_learning/model"
	"gorm.io/gen"
)

func GenDao() {
	generator := gen.NewGenerator(gen.Config{
		OutPath: "./gorm_learning/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	generator.UseDB(model.DbConn)

	// Generate basic type-safe API for struct `model.Student` following conventions
	generator.ApplyBasic(&model.Student{})
	generator.Execute()
}
