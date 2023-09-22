#include "common.h"
typedef char* (*testfunctionPointer)(char* str);

char* gatewayFunc(char* str);

void workflow(uintptr_t h, char* str);

void workflowForFuncPointer(testfunctionPointer f, char* str);

