#include <iostream>
#include "instrumentationTimer.h"

int main()
{
    Instrumentor::Get().BeginSession("Profile");
    RunBenchmarks();
    Instrumentor::Get().EndSession();
    std::cin.get();
}
