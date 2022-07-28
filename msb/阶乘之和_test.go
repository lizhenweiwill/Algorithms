package msb_test

import (
	"algorithms/msb"
	"testing"
)

func TestSumV0(t *testing.T) {
	t.Log(msb.SumV0(1))
	t.Log(msb.SumV0(2))
	t.Log(msb.SumV0(3))
	t.Log(msb.SumV0(4))
	t.Log(msb.SumV0(5))
}

// BenchmarkSumV0-10    	38786024	        30.83 ns/op
func BenchmarkSumV0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msb.SumV0(10)
	}
}

func TestSumV1(t *testing.T) {
	t.Log(msb.SumV1(1))
	t.Log(msb.SumV1(2))
	t.Log(msb.SumV1(3))
	t.Log(msb.SumV1(4))
	t.Log(msb.SumV1(5))
}

// BenchmarkSumV1-10    	266143251	         4.359 ns/op
func BenchmarkSumV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msb.SumV1(10)
	}
}
