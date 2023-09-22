package basic_learning

type Person struct {
	Name   string `tag1:"content1" tag2:"content2"`
	Age    int
	School []int
}

func TestStructTag() {
	person := new(Person)
	person.Name = "htx"
	person.Age = 24
	person.School = []int{1, 2, 3}

	//fmt.Println("通过反射获取某个变量的类型信息（该变量的名称，类型以及其字段的名称，类型）")
	//personType := reflect.TypeOf(*person)
	//fmt.Printf("personType.Kind()=%v\n", personType.Kind())
	//fmt.Printf("personType.Size()=%v\n", personType.Size())
	//fmt.Printf("personType.String()=%v\n", personType.String())
	//fmt.Printf("personType.NumField()=%v\n", personType.NumField())
	//fmt.Printf("personType.NumField()=%v\n", personType.NumMethod())
	//for i := 0; i < personType.NumField(); i++ {
	//	field := personType.Field(i)
	//	fmt.Printf("field%v --- field.Type =%v\n", i, field.Type)
	//	fmt.Printf("field%v --- field.Name =%v\n", i, field.Name)
	//	fmt.Printf("field%v --- field.PkgPath =%v\n", i, field.PkgPath)
	//	fmt.Printf("field%v --- field.tag1 =%v\n", i, field.Tag.Get("tag1"))
	//	fmt.Printf("field%v --- field.tag2 =%v\n", i, field.Tag.Get("tag2"))
	//}

	//fmt.Println("通过反射修改变量的字段值")
	//personValue := reflect.ValueOf(person)//接受指针作为参数时，才可以修改字段值
	//personValue.Elem().Field(0).SetString("alj")
	//sliceField := personValue.Elem().Field(2)
	//elem := reflect.New(sliceField.Type().Elem()).Elem()
	//elem.SetInt(123)
	//sliceField.Set(reflect.Append(sliceField, elem))
	//fmt.Printf("%v\n", personValue)

}
