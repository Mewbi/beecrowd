#include <stdio.h>

int main() {
    double n1, n2, n3, n4, m;
    scanf("%lf", &n1);
    scanf("%lf", &n2);
    scanf("%lf", &n3);
    scanf("%lf", &n4);
    m = (n1 * 2 + n2 * 3 + n3 * 4 + n4 * 1) / 10;
    printf("Media: %.1lf\n", m);

    if (m >= 7) {
        printf("Aluno aprovado.\n");
    } else if (m < 5) {
        printf("Aluno reprovado.\n");
    } else {
        printf("Aluno em exame.\n");
        scanf("%lf", &n1);
        printf("Nota do exame: %.1lf\n", n1);
        m = (m + n1) / 2;
        if (m >= 5) {
            printf("Aluno aprovado.\n");
        } else {
            printf("Aluno reprovado.\n");
        }
        printf("Media final: %.1lf\n", m);
    }

    return 0;
}
