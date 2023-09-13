using System;

class Program
{
    static void Main()
    {
        string[] nums = Console.ReadLine().Split(' ');
        int h1 = int.Parse(nums[0]);
        int m1 = int.Parse(nums[1]);
        int h2 = int.Parse(nums[2]);
        int m2 = int.Parse(nums[3]);

        int hours = h2 - h1;
        if (hours < 0)
        {
            hours = 24 - Math.Abs(hours);
        }

        int mins = m2 - m1;
        if (mins < 0)
        {
            hours -= 1;
            mins = 60 - Math.Abs(mins);
        }

        if (hours < 0)
        {
            hours = 24 + hours;
        }

        if (hours == 0 && mins == 0)
        {
            hours = 24;
        }

        Console.WriteLine($"O JOGO DUROU {hours} HORA(S) E {mins} MINUTO(S)");
    }
}
