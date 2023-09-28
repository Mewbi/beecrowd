#include <stdio.h>
#include <math.h>

double bhaskara(double a, double b, double c) {
    double x = sqrt(pow(b, 2) - (4 * a * c));
    double x1 = (-b + x) / (2 * a);
    double x2 = (-b - x) / (2 * a);

    if (x1 > 0) {
        return x1;
    }

    return x2;
}

void calc(double h, double p1, double p2, double a, double v) {
    double pi = 3.14159;
    double g = 9.80665;

    double vy = v * sin(a * pi / 180);
    double vx = v * cos(a * pi / 180);

    double t = bhaskara(-g / 2, vy, h);

    double alc = vx * t;

    const char *res = "NUCK";
    if (alc >= p1 && alc <= p2) {
        res = "DUCK";
    }

    printf("%.5f -> %s\n", alc, res);
}

int main() {
    double h, p1, p2, a, v;
    int n;

    while (scanf("%lf", &h) != EOF) {
        scanf("%lf %lf", &p1, &p2);
        scanf("%d", &n);

        for (int i = 0; i < n; i++) {
            scanf("%lf %lf", &a, &v);
            calc(h, p1, p2, a, v);
        }
    }

    return 0;
}

