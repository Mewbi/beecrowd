import java.io.IOException;
import java.util.Scanner;

/**
 * IMPORTANT: 
 *      O nome da classe deve ser "Main" para que a sua solução execute
 *      Class name must be "Main" for your solution to execute
 *      El nombre de la clase debe ser "Main" para que su solución ejecutar
 */
public class Main {
 
    public static void main(String[] args) throws IOException {
 
        Scanner scan = new Scanner(System.in);
        double a, b, c;
        int aresta;

        a = scan.nextInt();
        b = scan.nextInt();
        c = scan.nextInt();

        while (a != 0 && b != 0 && c != 0){
            aresta = (int)Math.cbrt(a * b * c);
            System.out.println(aresta);
            a = scan.nextInt();
            b = scan.nextInt();
            c = scan.nextInt();
        }
 
    }
 
}