#include <stdio.h>

int main() {
    int m, n, maior, menor, sum;
    while (1) {
        scanf("%d", &m);
        scanf("%d", &n);
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
            printf("%d ", i);
        }

        printf("Sum=%d\n", sum);
    }

    return 0;
}
