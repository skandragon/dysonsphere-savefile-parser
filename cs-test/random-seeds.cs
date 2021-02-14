using System;
public class RandomSeedTest
{
    public static void Main(string[] args)
    {
        Random r = new Random(8644885);

        Console.WriteLine("Using seed 8644885 to generate star randomness");

        r.Next();

        r.NextDouble();
        r.NextDouble();
        r.NextDouble();
        r.NextDouble();

        for (int i = 0; i<64;i++) {
            int seed2 = r.Next();
            Console.WriteLine("Star " + i + " seed " + seed2);
        }
    }
}
