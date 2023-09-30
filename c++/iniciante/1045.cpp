#include <iostream>
#include <algorithm>
#include <cmath>

int main() {
    float l[3];
    float a, b, c;
    std::cin >> l[0];
    std::cin >> l[1];
    std::cin >> l[2];

    std::sort(l, l + 3, std::greater<float>());
    a = l[0];
    b = l[1];
    c = l[2];

    if (a >= b + c) {
        std::cout << "NAO FORMA TRIANGULO" << std::endl;
        return 0;
    }

    if (std::pow(a, 2) == std::pow(b, 2) + std::pow(c, 2)) {
        std::cout << "TRIANGULO RETANGULO" << std::endl;
    }

    if (std::pow(a, 2) > std::pow(b, 2) + std::pow(c, 2)) {
        std::cout << "TRIANGULO OBTUSANGULO" << std::endl;
    }

    if (std::pow(a, 2) < std::pow(b, 2) + std::pow(c, 2)) {
        std::cout << "TRIANGULO ACUTANGULO" << std::endl;
    }

    if (a == b && b == c) {
        std::cout << "TRIANGULO EQUILATERO" << std::endl;
    }

    if ((a == b && a != c) || (a == c && a != b) || (b == c && b != a)) {
        std::cout << "TRIANGULO ISOSCELES" << std::endl;
    }

    return 0;
}
