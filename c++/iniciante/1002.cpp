#include <iostream>
#include <cmath>

int main() {
    double r, pi, a;
    pi = 3.14159;
    std::cin >> r;
    a = pi * std::pow(r, 2);
    printf("A=%.4lf\n", a);
    return 0;
}
