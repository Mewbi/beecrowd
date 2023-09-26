#include <stdio.h>
#include <math.h>

int kamenetsky(double n) {
    if (n < 0) {
        return 0;
    }

    if (n <= 1) {
        return 1;
    }

    double euler = 2.718281828459;
    double pi =    3.141592653589;
    double x = ((n * log10(n / euler)) +
                (log10(2 * pi * n) / 2.0));

    return (int)(floor(x) + 1);
}

int main() {
    double n;
    scanf("%lf", &n);
    printf("%d\n", kamenetsky(n));
    return 0;
}
