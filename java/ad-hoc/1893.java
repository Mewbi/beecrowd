import java.io.IOException;
import java.util.Scanner;

public class Main {
 
    public static void main(String[] args) throws IOException {
 
        Scanner scan = new Scanner(System.in);
        int m1, m2;

        m1 = scan.nextInt();
        m2 = scan.nextInt();

        if ((m1 >= 0) && (m2 <= 2)){
            System.out.println("nova");
        }else if ((m2 <= 96) && (m2 > m1)){
            System.out.println("crescente");
        }else if ((m2 <= 96) && (m1 > m2)){
            System.out.println("minguante");
        }else {
            System.out.println("cheia");
        }
 
    }
 
}