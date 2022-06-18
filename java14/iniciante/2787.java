import java.io.IOException;
import java.util.Scanner;

public class Main {
 
    public static void main(String[] args) throws IOException {
 
        Scanner scan = new Scanner(System.in);
        int x, y;

        x = scan.nextInt();
        y = scan.nextInt();

        if((x + y) % 2 == 0){
            System.out.println(1);
        }else{
            System.out.println(0);
        }
 
    }
 
}