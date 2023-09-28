#include <stdio.h>

int euclideanDivision(int a, int b, int *r) {
    int q, limit;
    limit = b < 0 ? -b : b;
    for (*r = 0; *r < limit; (*r)++) {
        if ((a - (*r)) % b == 0) {
            q = (a - (*r)) / b;
            return q;
        }
    }
    return 0;
}

int main() {
    int a, b, q, r;
    scanf("%d %d", &a, &b);

    q = euclideanDivision(a, b, &r);

    printf("%d %d\n", q, r);
    return 0;
}
