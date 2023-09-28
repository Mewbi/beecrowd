#include <iostream>

int main() {
    int n;
    while (std::cin >> n) {
        const char* out = "vai ter copa!";
        if (n > 0) {
            out = "vai ter duas!";
        }
        std::cout << out << std::endl;
    }
    return 0;
}

