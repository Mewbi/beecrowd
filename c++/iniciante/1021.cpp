#include <iostream>
#include <cmath>

double notas(double v, double n[], int size) {
    std::cout << "NOTAS:" << std::endl;
    for (int i = 0; i < size; i++) {
        printf("%d nota(s) de R$ %.2f\n", (int)(v / n[i]), n[i]);
        v = std::fmod(v, n[i]);
    }
    return v;
}

void moedas(double v, double m[], int size) {
    std::cout << "MOEDAS:" << std::endl;
    v += 0.0001; // Avoid precision problems
    for (int i = 0; i < size; i++) {
        printf("%d moeda(s) de R$ %.2f\n", (int)(v / m[i]), m[i]);
        v = std::fmod(v, m[i]);
    }
}

int main() {
    double n[] = {100.00, 50.00, 20.00, 10.00, 5.00, 2.00};
    double m[] = {1.00, 0.50, 0.25, 0.10, 0.05, 0.01};
    int nSize = sizeof(n) / sizeof(n[0]);
    int mSize = sizeof(m) / sizeof(m[0]);

    double v;
    std::cin >> v;

    v = notas(v, n, nSize);
    moedas(v, m, mSize);

    return 0;
}
