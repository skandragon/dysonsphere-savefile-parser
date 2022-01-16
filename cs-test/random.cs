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
using System.IO;

using System.Runtime.InteropServices;

namespace System2
{
	public class Random
	{
		public Random() : this(Environment.TickCount)
		{
		}

		public Random(int Seed)
		{
			int num = 161803398 - Math.Abs(Seed);
			this.SeedArray[55] = num;
			int num2 = 1;
			for (int i = 1; i < 55; i++)
			{
				int num3 = 21 * i % 55;
				this.SeedArray[num3] = num2;
				num2 = num - num2;
				if (num2 < 0)
				{
					num2 += int.MaxValue;
				}
				num = this.SeedArray[num3];
			}
			for (int j = 1; j < 5; j++)
			{
				for (int k = 1; k < 56; k++)
				{
					this.SeedArray[k] -= this.SeedArray[1 + (k + 30) % 55];
					if (this.SeedArray[k] < 0)
					{
						this.SeedArray[k] += int.MaxValue;
					}
				}
			}
			this.inext = 0;
			this.inextp = 31;
		}

		protected virtual double Sample()
		{
			if (++this.inext >= 56)
			{
				this.inext = 1;
			}
			if (++this.inextp >= 56)
			{
				this.inextp = 1;
			}
			int num = this.SeedArray[this.inext] - this.SeedArray[this.inextp];
			if (num < 0)
			{
				num += int.MaxValue;
			}
			this.SeedArray[this.inext] = num;
			return (double)num * 4.656612875245797E-10;
		}

		public virtual int Next()
		{
			return (int)(this.Sample() * 2147483647.0);
		}

		public virtual int Next(int maxValue)
		{
			if (maxValue < 0)
			{
				throw new ArgumentOutOfRangeException("Max value is less than min value.");
			}
			return (int)(this.Sample() * (double)maxValue);
		}

		public virtual int Next(int minValue, int maxValue)
		{
			if (minValue > maxValue)
			{
				throw new ArgumentOutOfRangeException("Min value is greater than max value.");
			}
			uint num = (uint)(maxValue - minValue);
			if (num <= 1U)
			{
				return minValue;
			}
			return (int)((ulong)((uint)(this.Sample() * num)) + (ulong)((long)minValue));
		}

		public virtual void NextBytes(byte[] buffer)
		{
			if (buffer == null)
			{
				throw new ArgumentNullException("buffer");
			}
			for (int i = 0; i < buffer.Length; i++)
			{
				buffer[i] = (byte)(this.Sample() * 256.0);
			}
		}

		public virtual double NextDouble()
		{
			return this.Sample();
		}

		private const int MBIG = 2147483647;

		private const int MSEED = 161803398;

		private const int MZ = 0;

		private int inext;

		private int inextp;

		private int[] SeedArray = new int[56];
	}
}

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
        System2.Random r = new System2.Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.Next().ToString());
            }
        }
    }

    private static void generateNextDouble(int count, string filename, int seed)
    {
        System2.Random r = new System2.Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.NextDouble().ToString());
            }
        }
    }

    private static void generateNextMax(int count, string filename, int seed)
    {
        System2.Random r = new System2.Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.Next(seed).ToString());
            }
        }
    }

    private static void generateNextRange(int count, string filename, int seed)
    {
        System2.Random r = new System2.Random(seed);
        using (StreamWriter file = new StreamWriter(filename))
        {
            for (int i = 0 ; i < count ; i++) {
                file.WriteLine(r.Next(-seed, seed).ToString());
            }
        }
    }
}
