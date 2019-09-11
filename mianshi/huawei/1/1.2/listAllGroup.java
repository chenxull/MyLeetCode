import java.io.*;
import java.util.Scanner;
import java.math.*;

public class listAllGroup {

	public static void main(String[] args) throws NumberFormatException, IOException {
		Scanner s = new Scanner(System.in);
		int n = s.nextInt();
		int m = s.nextInt();
		// BigInteger n=BigInteger.valueOf(a);
		// BigInteger m=BigInteger.valueOf(b);
		int num = amout(n, m);
		int answer = (num * amout(m * 2 - 1, m - 1)) % 1000000007;
		System.out.println(answer);
	}

	public static int amout(int a, int b) {
		int da = 1;
		int x = 1;
		for (int i = 0; i < b; i++) {
			da = da * a;
			a--;
		}
		for (; b > 0; b--) {
			x = x * b;
		}
		return da / x;
	}

}