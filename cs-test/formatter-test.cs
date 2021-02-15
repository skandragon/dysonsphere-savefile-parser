using System;
public class FormatterTest
{
    public static void Main(string[] args)
    {
        Random random = new Random();

        int num2 = random.Next(15, 26);
        int num3 = random.Next(0, 26);
        char c = (char)(65 + num2);
        char c2 = (char)(65 + num3);
        string s = (int)(c + c2) + " nameGoesHere";
        Console.WriteLine(s);
        Console.WriteLine((int)c);
        Console.WriteLine((int)c2);
    }
}
