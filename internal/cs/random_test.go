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

package cs

import (
	"bufio"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNext1000(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/next-1000.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(1000)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.Next()))
	}
	require.NoError(t, sc.Err())
}

func TestNext5432(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/next-5432.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(5432)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.Next()))
	}
	require.NoError(t, sc.Err())
}

func TestNextMax32768(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextmax-32768.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(32768)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.NextWithMax(32768)))
	}
	require.NoError(t, sc.Err())
}

func TestNextMax810485(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextmax-810485.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(810485)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.NextWithMax(810485)))
	}
	require.NoError(t, sc.Err())
}

func TestNextDouble6195(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextdouble-6195.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(6195)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.ParseFloat(line, 64)
		require.NoError(t, err)
		require.InEpsilon(t, expected, r.NextDouble(), 1e-14)
	}
	require.NoError(t, sc.Err())
}

func TestNextDouble19596839(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextdouble-19596839.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(19596839)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.ParseFloat(line, 64)
		require.NoError(t, err)
		require.InEpsilon(t, expected, r.NextDouble(), 1e-14)
	}
	require.NoError(t, sc.Err())
}

func TestNextSmallRange1001(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextsmallrange-1001.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(1001)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.NextRange(-1001, 1001)))
	}
	require.NoError(t, sc.Err())
}

func TestNextSmallRange32768(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextsmallrange-32768.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(32768)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.NextRange(-32768, 32768)))
	}
	require.NoError(t, sc.Err())
}

func TestNextLargeRange1073741824(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextlargerange-1073741824.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(1073741824)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.NextRange(-1073741824, 1073741824)))
	}
	require.NoError(t, sc.Err())
}
func TestNextLargeRange1073741825(t *testing.T) {
	f, err := os.OpenFile("./testdata/random/nextlargerange-1073741825.txt", os.O_RDONLY, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	r := MakePRNGSequence(1073741825)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		expected, err := strconv.Atoi(line)
		require.NoError(t, err)
		require.Equal(t, expected, int(r.NextRange(-1073741825, 1073741825)))
	}
	require.NoError(t, sc.Err())
}

func TestNext12345NegativeSeed(t *testing.T) {
	r1 := MakePRNGSequence(12345)
	r2 := MakePRNGSequence(-12345)

	for i := 0; i < 50000; i++ {
		require.Equal(t, r1.Next(), r2.Next())
	}
}
