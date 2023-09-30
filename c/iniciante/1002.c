#include <stdio.h>
#include <math.h>

int main() {
    double r, pi, a;
    pi = 3.14159;
    scanf("%lf", &r);
    a = pi * pow(r, 2);
    printf("A=%.4lf\n", a);
    return 0;
}
