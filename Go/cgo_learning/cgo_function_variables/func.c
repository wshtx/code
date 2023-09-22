#include "func.h"
extern void testFunctionVariable(uintptr_t h, char* str);
extern char* callbackFunc(char* str);

//gateway function
char* gatewayFunc(char* str) {
	return callbackFunc(str);
}

void workflow(uintptr_t h, char* str){
	puts("the first step in c side");
	testFunctionVariable(h, str);
	puts("the last step in c side");
}

void workflowForFuncPointer(testfunctionPointer f, char* str){
	puts("the first step in c side");
	puts("the second step in go side: ");
	char* res = f(str);
	puts(res);
	puts("the second step in c side");
}