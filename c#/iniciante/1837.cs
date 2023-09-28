using System;

class Program
{
    static Tuple<int, int> EuclideanDivision(int a, int b)
    {
        int q, r, limit;
        limit = b < 0 ? -b : b;
        r = 0;
        for (; r < limit; r++)
        {
            if ((a - r) % b == 0)
            {
                q = (a - r) / b;
                return Tuple.Create(q, r);
            }
        }
        return Tuple.Create(0, 0);
    }

    static void Main(string[] args)
    {
        int a, b, q, r;
        string[] input = Console.ReadLine().Split();
        a = int.Parse(input[0]);
        b = int.Parse(input[1]);

        var result = EuclideanDivision(a, b);
        q = result.Item1;
        r = result.Item2;

        Console.WriteLine($"{q} {r}");
    }
}
