using System;
using System.IO;

//
// Create a set of random values for testing.  Ideally,
// we will test all the functions we use in the rest of the
// code, so we can ensure that starting at a given seed, the
// go PRNG sequence class will produce the same output.
//
public class RandomTest
{
    const int TESTLEN = 1000;

    public static void Main(string[] args)
    {
        generateNext(TESTLEN, "next-1000.txt", 1000);
        generateNext(TESTLEN, "next-5432.txt", 5432);

        generateNextDouble(TESTLEN, "nextdouble-19596839.txt", 19596839);
        generateNextDouble(TESTLEN, "nextdouble-6195.txt", 6195);

        generateNextMax(TESTLEN, "nextmax-32768.txt", 32768);
        generateNextMax(TESTLEN, "nextmax-810485.txt", 810485);

        generateNextRange(TESTLEN, "nextsmallrange-32768.txt", 32768);
        generateNextRange(TESTLEN, "nextsmallrange-1001.txt", 1001);

        generateNextRange(TESTLEN, "nextlargerange-1073741824.txt", 1073741824);
        generateNextRange(TESTLEN, "nextlargerange-1073741825.txt", 1073741825);
    }

    private static void generateNext(int count, string filename, int seed)
    {
        Random r = new Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.Next().ToString());
            }
        }
    }

    private static void generateNextDouble(int count, string filename, int seed)
    {
        Random r = new Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.NextDouble().ToString());
            }
        }
    }

    private static void generateNextMax(int count, string filename, int seed)
    {
        Random r = new Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.Next(seed).ToString());
            }
        }
    }

    private static void generateNextRange(int count, string filename, int seed)
    {
        Random r = new Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.Next(-seed, seed).ToString());
            }
        }
    }
}
