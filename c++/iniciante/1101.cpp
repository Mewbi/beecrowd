#include <iostream>

int main() {
    int m, n, maior, menor, sum;
    while (1) {
        std::cin >> m;
        std::cin >> n;
        if (n <= 0 || m <= 0) {
            break;
        }

        maior = m;
        menor = n;
        if (n > m) {
            maior = n;
            menor = m;
        }

        sum = 0;
        for (int i = menor; i <= maior; i++) {
            sum += i;
            std::cout << i << " ";
        }

        std::cout << "Sum=" << sum << std::endl;
    }

    return 0;
}
