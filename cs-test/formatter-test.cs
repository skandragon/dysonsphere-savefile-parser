/*
 * Copyright 2021-2022 Michael Graff
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
