import java.io.IOException;
import java.util.Scanner;

public class Main {
 
    public static void main(String[] args) throws IOException {
 
        Scanner scan = new Scanner(System.in);
        int l, c, t1, t2;

        l = scan.nextInt();
        c = scan.nextInt();

        t1 = (l * c) + ((c - 1) * (l - 1));
        t2 = ((l - 1) * 2) + ((c -1) * 2);

        System.out.println(t1);
        System.out.println(t2);
 
    }
 
}