package cgo_type_conversion

/*
#cgo CFLAGS: -I../include
#include "common.h"
void printArr(int* arr, int len){
	for(int i  = 0;i < len; i++){
		printf("%d\n", arr[i]);
	}
	arr[0] = 12;
}

void printDoubleArr(double* arr, int len){
	for(int i  = 0;i < len; i++){
		printf("%f\n", arr[i]);
	}
	arr[0] = 12.222;
}

void printCharArr(char* arr, int len){
	for(int i  = 0;i < len; i++){
		printf("%c\n", arr[i]);
	}
}

struct Person{
	int age;
	char* name;
} personarr[] = {
{24, "htxx"},
{23, "aljj"}
};

struct Person *itemPersons(struct Person *person, int index){
	return &person[index];
}

*/
import "C"

func TestPassingGoPointer() {
	//arrGo := []int32{1, 2, 3, 4}
	//C.printArr((*C.int)(unsafe.Pointer(&arrGo[0])), C.int(len(arrGo)))
	//fmt.Println(arrGo)
	//fmt.Println("----------------")

	//arrfloatGo := []float64{1.22, 2.255, 3.669, 4.77}
	//C.printDoubleArr((*C.double)(unsafe.Pointer(&arrfloatGo[0])), C.int(len(arrfloatGo)))
	//fmt.Println(arrfloatGo)
	//fmt.Println("----------------")

	//strGo := "htxhtx"
	//fmt.Println(len(strGo))
	//strGoBytes := *(*[]byte)(unsafe.Pointer(&strGo))
	//C.printCharArr((*C.char)(unsafe.Pointer(&strGoBytes[0])), C.int(len(strGo)))
	//fmt.Println("----------------")

	//visit the struct array
	//person := C.itemPersons(&C.personarr[0], 1)
	//fmt.Printf("%T, %v\n", person, person)
	//fmt.Println(person.age)
	//fmt.Println(C.GoString(person.name))
}
