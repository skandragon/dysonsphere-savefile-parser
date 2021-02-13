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

func TestNext12345(t *testing.T) {
	r := MakePRNGSequence(12345)
	expected := []int32{
		143337951, 150666398, 1663795458, 1097663221, 1712597933, 1776631026, 356393799, 1580828476,
		558810388, 1086637143, 494509053, 831377771, 463814839, 44691035, 289552956, 1590924033,
		418954878, 1904902962, 1849199486, 770656856, 222698908, 1137270943, 770420532, 1519356451,
		1246560209, 1332617375, 1573024538, 1606065954, 850942673, 526685912, 1473914819, 452144508,
		1111403504, 1042369160, 542895576, 1655234974, 5538230, 1039193352, 961982272, 1044665811,
		1528100810, 969047112, 579718272, 607824875, 1364170491, 633032322, 793567355, 1831117809,
		377238926, 1830086762, 1383740914, 1322492187, 948158774, 1066648348, 64646849, 1153550655,
		1527729513, 144439007, 1998586659, 379980558, 203606488, 897811492, 729885803, 32124476,
	}

	for i, exp := range expected {
		got := r.Next()
		if exp != got {
			t.Errorf("Expected %d, got %d, index %d", exp, got, i)
		}
	}
}
