#include <stdio.h>
#include <math.h>

int main() {
    int n, anos;
    double pa, pb, g1, g2;

    scanf("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf("%lf %lf %lf %lf", &pa, &pb, &g1, &g2);

        anos = 0;
        while (1) {
            if (anos > 100 || pa > pb) {
                break;
            }

            anos += 1;
            pa += trunc(pa * (g1 / 100));
            pb += trunc(pb * (g2 / 100));
        }

        if (anos > 100) {
            printf("Mais de 1 seculo.\n");
            continue;
        }

        printf("%d anos.\n", anos);
    }

    return 0;
}
