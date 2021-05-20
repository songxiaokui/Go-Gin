// hello.cpp

#include <iostream>

extern "C" {
    #include "hello.h"
}

void SayHello2(const char* s) {
    std::cout << s;
}