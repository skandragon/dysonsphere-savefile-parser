using System;

public class HelloWorld
{
    public static void Main(string[] args)
    {
        Random r = new Random(12345);
        for (int i = 0 ; i < 64 ; i++) {
            Console.WriteLine(r.Next());
        }
    }
}
