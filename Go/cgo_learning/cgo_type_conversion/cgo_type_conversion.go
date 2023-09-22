package cgo_type_conversion

/*
#include <string.h>
#include <stdint.h>
#include <stdio.h>
struct Person{
	char* name;
	int age;
};
struct Person person = {"htx", 24};

enum Color{
	WHITE,BLUE,BLACK
};
struct Person *personPtr = &person;
int arr[] ={1,2,3,4};

double arrDouble[] ={1.2,2.2,3.2,4.2};

void testPassPointer();

uintptr_t testShareMemory(char* str){
	//str[0] = 'a';
	//puts(str);
	return (uintptr_t)str;
}

uintptr_t testShareMemory2(void* str){
	//str[0] = 'a';
	//puts((char*)str);
	return (uintptr_t)str;
}


*/
import "C"

func TestTypeConversion() {
	//fmt.Printf("visit c enum, %T, %v", C.BLACK, C.BLACK)

	//fmt.Printf("visit a int field of a c struct , Type=%T, value=%v\n", C.person.age, C.person.age)
	//var personS C.struct_Person
	//personS.age = C.int(14)
	//personS.name = C.CString("alj")
	//fmt.Printf("create a c struct in go, Type=%T, name = %v, age = %v\n", personS, personS.name, personS.age)

	//the first method to convert the char* to string of go
	//pnamelen := int(C.strlen(C.person.name))
	//pname := string((*[31]byte)(unsafe.Pointer(C.person.name))[:pnamelen:pnamelen])

	//the second method to convert the char* to string of go, though this method could allocate memory for the new gostring.
	//pname := C.GoString(C.person.name)
	//fmt.Printf("visit a char* field of a c struct, you need to convert the char* of c to string of golang, Type=%T, value=%v\n", pname, pname)

	//the first method to convert the int[] to slice
	//arrGo := (*[1 << 29]int32)(unsafe.Pointer(&C.arr))[0:4:4]
	//the second method to convert the int[] to slice
	//slice := unsafe.Slice(&C.arr[0], 4)
	//fmt.Printf("%T ,%v\n", slice, slice)
	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}
	//fmt.Printf("visit a int[] in c , you need to convert the int[] of c to slice of golang, Type=%T, value=%v\n", arrGo, arrGo)
	//arrDoubleGo := (*[1 << 29]float64)(unsafe.Pointer(&C.arrDouble))[0:4:4]
	//fmt.Printf("visit a int[] in c , you need to convert the int[] of c to slice of golang, Type=%T, value=%v\n", arrDoubleGo, arrDoubleGo)

	//传递go字符串（不拷贝）
	//str := "htxhtx"
	//stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	////fmt.Printf("%T, %v\n", stringHeader.Data, stringHeader.Data)
	//fmt.Printf("%T\n", stringHeader.Data)
	//res := uintptr(C.testShareMemory2(unsafe.Pointer(stringHeader.Data)))
	//fmt.Printf("%T", res)

}
