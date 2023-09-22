#include <iostream>

extern "C"{
    #include <test.h>
}

void TestCMethod(char* s){
    std::cout << s << std::endl;
}