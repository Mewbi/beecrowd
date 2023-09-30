#include <iostream>
#include <cmath>

int main() {
    int n, anos;
    double pa, pb, g1, g2;

    std::cin >> n;

    for (int i = 0; i < n; i++) {
        std::cin >> pa >> pb >> g1 >> g2;

        anos = 0;
        while (1) {
            if (anos > 100 || pa > pb) {
                break;
            }

            anos += 1;
            pa += std::trunc(pa * (g1 / 100));
            pb += std::trunc(pb * (g2 / 100));
        }

        if (anos > 100) {
            std::cout << "Mais de 1 seculo." << std::endl;
            continue;
        }

        std::cout << anos << " anos." << std::endl;
    }

    return 0;
}
