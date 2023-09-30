#include <iostream>

typedef unsigned long long ull;

ull nonzero(ull n) {
    if (n == 0 || n == 1)
        return 1;

    ull x = n / 5;
    ull y = n % 5;

    ull potX = x % 4;
    switch (potX) {
        case 0:
            potX = 1;
            break;
        case 1:
            potX = 2;
            break;
        case 2:
            potX = 4;
            break;
        case 3:
            potX = 8;
            break;
    }

    ull lastY;
    switch (y) {
        case 0:
            lastY = 1;
            break;
        case 1:
            lastY = 1;
            break;
        case 2:
            lastY = 2;
            break;
        case 3:
            lastY = 6;
            break;
        case 4:
            lastY = 4;
            break;
        default:
            lastY = 0;
    }

    return (potX * nonzero(x) * lastY);
}

int main() {
    ull n, result, instancia = 1;

    while (std::cin >> n) {
        std::cout << "Instancia " << instancia << std::endl;
        instancia++;

        result = nonzero(n);
        while (result >= 10)
            result = result % 10;

        std::cout << result << std::endl << std::endl;
    }
    return 0;
}
