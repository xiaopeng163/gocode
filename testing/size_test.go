package size

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "small"},
}

func TestSize(t *testing.T) {
	for i, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Error("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}

func BenchmarkSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Size(100)
	}
}
