import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int a, b, q, r;
        a = scanner.nextInt();
        b = scanner.nextInt();

        Pair result = euclideanDivision(a, b);
        q = result.q;
        r = result.r;

        System.out.println(q + " " + r);
    }

    static class Pair {
        int q, r;

        Pair(int q, int r) {
            this.q = q;
            this.r = r;
        }
    }

    static Pair euclideanDivision(int a, int b) {
        int q, r, limit;
        limit = (b < 0) ? -b : b;
        r = 0;
        for (; r < limit; r++) {
            if ((a - r) % b == 0) {
                q = (a - r) / b;
                return new Pair(q, r);
            }
        }
        return new Pair(0, 0);
    }
}
