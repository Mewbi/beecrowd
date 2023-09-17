#include <iostream>

int main() {
    unsigned int x, y;

    while (std::cin >> x >> y) {
        std::cout << (x ^ y) << std::endl;
    }

    return 0;
}
