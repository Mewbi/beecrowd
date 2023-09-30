#include <stdio.h>
#include <math.h>

void sort(float arr[3]) {
    for (int i = 0; i < 3 - 1; i++) {
        for (int j = 0; j < 3 - i - 1; j++) {
            if (arr[j] < arr[j + 1]) {
                float temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
}

int main() {
    float l[3];
    float a, b, c;
    scanf("%f", &l[0]);
    scanf("%f", &l[1]);
    scanf("%f", &l[2]);

    sort(l);
    a = l[0];
    b = l[1];
    c = l[2];

    if (a >= b + c) {
        printf("NAO FORMA TRIANGULO\n");
        return 0;
    }

    if (pow(a, 2) == pow(b, 2) + pow(c, 2)) {
        printf("TRIANGULO RETANGULO\n");
    }

    if (pow(a, 2) > pow(b, 2) + pow(c, 2)) {
        printf("TRIANGULO OBTUSANGULO\n");
    }

    if (pow(a, 2) < pow(b, 2) + pow(c, 2)) {
        printf("TRIANGULO ACUTANGULO\n");
    }

    if (a == b && b == c) {
        printf("TRIANGULO EQUILATERO\n");
    }

    if ((a == b && a != c) || (a == c && a != b) || (b == c && b != a)) {
        printf("TRIANGULO ISOSCELES\n");
    }

    return 0;
}
