using System;

class Program
{
    static int Kamenetsky(double n)
    {
        if (n < 0)
        {
            return 0;
        }

        if (n <= 1)
        {
            return 1;
        }

        double x = ((n * Math.Log10(n / Math.E)) +
                    (Math.Log10(2 * Math.PI * n) / 2.0));

        return (int)(Math.Floor(x) + 1);
    }

    static void Main()
    {
        double n;
        if (double.TryParse(Console.ReadLine(), out n))
        {
            Console.WriteLine(Kamenetsky(n));
        }
    }
}
