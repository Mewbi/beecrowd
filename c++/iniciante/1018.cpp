#include <iostream>

int main() {
    int n;
    int cedulas[] = {100, 50, 20, 10, 5, 2, 1};
    std::cin >> n;
    std::cout << n << std::endl;
    
    for (int i = 0; i < 7; i++) {
        std::cout << n / cedulas[i] << " nota(s) de R$ " << cedulas[i] << ",00" << std::endl;
        n = n % cedulas[i];
    }
    
    return 0;
}
