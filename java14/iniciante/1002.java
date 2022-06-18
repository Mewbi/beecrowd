import java.io.IOException;
import java.util.Scanner; 
 
public class Main {
 
    public static void main(String[] args) throws IOException {
 
        double raio;
        double pi = 3.14159;

        Scanner scan = new Scanner(System.in);
        raio = scan.nextDouble();

        System.out.printf("A=%.4f\n", pi * Math.pow(raio, 2));
 
    }
 
}