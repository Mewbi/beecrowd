import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        double h, p1, p2, a, v;
        int n;

        while (scanner.hasNext()) {
            h = scanner.nextDouble();
            p1 = scanner.nextDouble();
            p2 = scanner.nextDouble();
            n = scanner.nextInt();

            for (int i = 0; i < n; i++) {
                a = scanner.nextDouble();
                v = scanner.nextDouble();
                calc(h, p1, p2, a, v);
            }
        }
    }

    static double bhaskara(double a, double b, double c) {
        double x = Math.sqrt(Math.pow(b, 2) - (4 * a * c));
        double x1 = (-b + x) / (2 * a);
        double x2 = (-b - x) / (2 * a);

        if (x1 > 0) {
            return x1;
        }

        return x2;
    }

    static void calc(double h, double p1, double p2, double a, double v) {
        double pi = 3.14159;
        double g = 9.80665;

        double vy = v * Math.sin(a * pi / 180);
        double vx = v * Math.cos(a * pi / 180);

        double t = bhaskara(-g / 2, vy, h);

        double alc = vx * t;

        String res = "NUCK";
        if (alc >= p1 && alc <= p2) {
            res = "DUCK";
        }

        System.out.printf("%.5f -> %s\n", alc, res);
    }
}
