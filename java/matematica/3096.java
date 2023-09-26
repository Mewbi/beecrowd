import java.util.Scanner;

public class Main {
    public static int kamenetsky(double n) {
        if (n < 0) {
            return 0;
        }

        if (n <= 1) {
            return 1;
        }

        double x = ((n * Math.log10(n/Math.E)) +
                    (Math.log10(2*Math.PI*n) / 2.0));

        return (int)Math.floor(x) + 1;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        double n = scanner.nextDouble();
        System.out.println(kamenetsky(n));
    }
}
