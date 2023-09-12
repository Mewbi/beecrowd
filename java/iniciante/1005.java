import java.io.IOException;
import java.util.Scanner;
 
public class Main {
 
    public static void main(String[] args) throws IOException {
 
        float a, b;

        Scanner scan = new Scanner(System.in);
        a = scan.nextFloat();
        b = scan.nextFloat();

        System.out.printf("MEDIA = %.5f\n", ((a * 0.35)+(b * 0.75))/1.1);
    }
 
}