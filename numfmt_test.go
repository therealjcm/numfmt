package numfmt

import (
	"testing"
	"math/rand"
)

var intTestData  = map [int]string {
	1: "1",
	-1: "-1",
	0:	"0",
	123: "123",
	-123: "-123",
	1234: "1,234",
	-1234: "-1,234",
	12345: "12,345",
	-12345: "-12,345",
	123456: "123,456",
	-123456: "-123,456",
	123456789: "123,456,789",
	-123456789: "-123,456,789",
}

var floatTestData = map [float64]string {
	1: "1.00",
	.01: "0.01",
	-.01: "-0.01",
	1.01: "1.01",
	-1: "-1.00",
	0:	"0.00",
	123: "123.00",
	-123: "-123.00",
	1234: "1,234.00",
	-1234: "-1,234.00",
	12345: "12,345.00",
	-12345: "-12,345.00",
	123456: "123,456.00",
	-123456: "-123,456.00",
	123456789: "123,456,789.00",
	-123456789: "-123,456,789.00",
	1234.56: "1,234.56",
	-1234.56: "-1,234.56",
}

func TestIntCommas(t *testing.T) {
	for k, v := range intTestData {
		if v != IntCommas(k) {
			t.Errorf("IntCommas(%d), expected(%s), got(%s)\n", k, v, IntCommas(k))
		}
	}
}

func TestFloatCommas(t *testing.T) {
	for k, v := range floatTestData {
		if v != FloatCommas(k, 2) {
			t.Errorf("FloatCommas(%f), expected(%s), got(%s)\n", k, v, FloatCommas(k, 2))
		}
	}
}

func BenchmarkIntCommas(b *testing.B) {
	rand.Seed(1337)
	nt := 1000
	tests := make([]int, nt)
	for i := 0; i < nt; i++ {
		tests[i] = rand.Int()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := tests[i % nt]
		IntCommas(n)
	}
}

func BenchmarkFloatCommas(b *testing.B) {
	rand.Seed(1337)
	nt := 1000
	tests := make([]float64, nt)
	for i := 0; i < nt; i++ {
		tests[i] = rand.Float64()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f := tests[i % nt]
		FloatCommas(f, 2)
	}
}
